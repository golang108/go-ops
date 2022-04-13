package v1

type FileJSONInfo struct {
	Name    string      `json:"name" dc:"文件名称"`
	Type    string      `json:"type" dc:"文件类型"`
	Size    int64       `json:"size" dc:"大小"`
	Path    string      `json:"path" dc:"路径"`
	ModTime int64       `json:"mtime" dc:"修改时间"`
	Extra   interface{} `json:"extra,omitempty"`
}
