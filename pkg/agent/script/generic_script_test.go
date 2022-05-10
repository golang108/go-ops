package script

import (
	"context"
	"fmt"
	"go-ops/internal/model"

	"testing"

	"github.com/gogf/gf/v2/util/guid"
)

func testGetScriptJob(content, spath string) model.ScriptJob {
	s := model.Script{
		Content: content,
		Path:    spath,
	}
	s1 := model.ScriptJob{
		Jobid:  guid.S(),
		Script: s,
	}
	return s1
}

func TestGenericScript(t *testing.T) {
	s := testGetScriptJob("ls", "")
	ctx, _ := context.WithCancel(context.Background())
	res := NewJobScriptProvider(ctx, s).Run()

	fmt.Println("res0:", res)

	s = testGetScriptJob("pwd", "")
	res = NewJobScriptProvider(ctx, s).Run()
	fmt.Println("res1:", res)

	s = testGetScriptJob("pwd", "/Users/luxingwen/go/src/go-ops/pkg/agent/script")
	res = NewJobScriptProvider(ctx, s).Run()
	fmt.Println("res2:", res)

	s = testGetScriptJob("ls", "/Users/luxingwen/go/src/go-ops/pkg/agent/script")
	res = NewJobScriptProvider(ctx, s).Run()
	fmt.Println("res3:", res)
}
