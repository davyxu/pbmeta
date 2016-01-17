package main

import (
	"github.com/davyxu/pbmeta"
	"testing"
)

func TestMeta(t *testing.T) {

	// 请先运行ExportPluginMeta导出test.pb
	fds, err := pbmeta.LoadFileDescriptorSet("test.pb")

	if err != nil {
		t.Error(err)
		return
	}

	// 描述池
	pool := pbmeta.NewDescriptorPool(fds)

	// 取消息
	m := pool.MessageByFullName("test.TestStruct")

	if m == nil {
		t.Fail()
		return
	}

	// 检查Comment读取正确
	if "FieldA comment" != m.FieldByName("FieldA").CommentMeta.TrailingComment() {
		t.Fail()
		return
	}

}
