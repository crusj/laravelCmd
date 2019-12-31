package writer

import (
	"testing"
)

func TestRoutesWriter(t *testing.T) {
	parser := NewRoutesParser("../api.json")
	writer := NewRoutesWriter("../test_Route.php", "//route开始", "//route结束")
	err := Write(parser,writer)
	if err != nil {
		t.Fatal("写入失败",err)
	}
}
