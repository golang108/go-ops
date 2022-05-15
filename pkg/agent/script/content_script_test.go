package script

import (
	"context"
	"fmt"
	"go-ops/internal/model"

	"testing"

	"github.com/gogf/gf/v2/util/guid"
)

func testGetContentScriptJob(content, spath string) model.ScriptJob {
	s := model.Script{
		Content: content,
		Path:    spath,
		ExecWay: model.ExecContent,
		Ext:     ".sh",
	}
	s1 := model.ScriptJob{
		Jobid:  guid.S(),
		Script: s,
	}
	return s1
}

func TestContentGenericScript(t *testing.T) {
	//	s := testGetScriptJob("ls", "")
	ctx, _ := context.WithCancel(context.Background())
	//	res := NewJobScriptProvider(ctx, s).Run()

	s := testGetContentScriptJob("echo $HOME", "")
	res := NewJobScriptProvider(ctx, s).Run()
	fmt.Println("res:", res)
}
