package script

import (
	"context"
	"fmt"
	"go-ops/pkg/agent/action"
	"go-ops/pkg/proto"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/charlievieth/fs"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "OPS_AGENT_PLUGIN",
	MagicCookieValue: "luxingwen",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"run": &OpsAgentGRPCPlugin{},
}

// Here is the gRPC server that GRPCClient talks to.
type OpsAgentGRPCPluginServer struct {
	// This is the real implementation
	Impl OpsAgentPlugin
}

func (m *OpsAgentGRPCPluginServer) Run(
	ctx context.Context,
	req *proto.Request) (*proto.Response, error) {

	r, err := m.Impl.Run(req.Body)
	return &proto.Response{Body: r}, err
}

type OpsAgentPlugin interface {
	Run(body []byte) (r []byte, err error)
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type OpsAgentGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl OpsAgentPlugin
}

func (p *OpsAgentGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterOpsAgentPluginServer(s, &OpsAgentGRPCPluginServer{Impl: p.Impl})
	return nil
}

func (p *OpsAgentGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &AgentPlugin{client: proto.NewOpsAgentPluginClient(c)}, nil
}

type AgentPlugin struct {
	DownloadUrl string        `json:"downloadUrl"` // 插件下载地址
	Name        string        `json:"name"`        // 插件名称
	Version     string        `json:"version"`     // 插件版本
	Md5         string        `json:"md5"`         // 插件md5
	Cmd         string        `json:"cmd"`         // 插件运行解释器
	CmdArgs     string        `json:"cmdArgs"`     // 插件解释器运行
	Timeout     time.Duration // 插件运行超时时间
	Args        []string      `json:"args"` // 插件启动参数

	client proto.OpsAgentPluginClient
}

func (agentPlugin *AgentPlugin) DirName() (r string) {
	return fmt.Sprintf("%s-%s-%s", agentPlugin.Name, agentPlugin.Version, agentPlugin.Md5)
}

func (agentPlugin *AgentPlugin) Filename() (r string) {
	dir := agentPlugin.DirName()
	r = filepath.Join(dir, agentPlugin.Name)
	return
}

func (agentPlugin *AgentPlugin) DownloadAndUnTar(ctx context.Context) (err error) {
	dir := agentPlugin.DirName()

	err = fs.MkdirAll(dir, os.FileMode(0750))
	if err != nil {
		return
	}

	filename := filepath.Join(dir, filepath.Base(agentPlugin.DownloadUrl))

	err = action.Download(ctx, filename, agentPlugin.DownloadUrl)
	if err != nil {
		return
	}

	err = action.CheckFileMd5(filename, agentPlugin.Md5)
	if err != nil {
		return
	}

	err = action.Untar(filename)
	return
}

// 插件是否存在
func (agentPlugin *AgentPlugin) IsExist() (r bool) {
	filename := agentPlugin.Filename()
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return
}

// 运行插件
func (agentPlugin *AgentPlugin) Run(ctx context.Context, reqBody []byte) (res []byte, err error) {

	if !agentPlugin.IsExist() {
		err = agentPlugin.DownloadAndUnTar(ctx)
		if err != nil {
			return
		}
		return
	}

	if agentPlugin.Cmd == "" {
		agentPlugin.Cmd = Cmder
	}

	cmdstr, cmdArgs := getCmdArgs(agentPlugin.Cmd)

	args := []string{agentPlugin.Name}
	args = append(args, cmdArgs...)

	args = append(args, agentPlugin.Args...)

	cmd := exec.Command(cmdstr, args...)

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: Handshake,
		Plugins:         PluginMap,
		Cmd:             cmd,
		StartTimeout:    agentPlugin.Timeout,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("run")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	op := raw.(OpsAgentPlugin)

	res, err = op.Run(reqBody)
	return
}
