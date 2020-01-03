package writer

import (
	"testing"
)

func TestRoutesWriter(t *testing.T) {
	parser := NewRoutesParser("../api.json")
	writer := NewRoutesWriter("../test_Route.php", "//route开始", "//route结束")
	err := Write(parser, writer)
	if err != nil {
		t.Fatal("写入失败", err)
	}
}
func TestControllerWriter(t *testing.T) {
	parser := NewRoutesParser("../api.json")
	writer := NewControllerWriter("../app/Http/Controllers/", "//controller_start", "//controller_end")
	if err := Writes(parser, writer); err != nil {
		t.Fatal(err)
	}
}
func TestRequestRuleWriter(t *testing.T) {
	parser := NewRoutesParser("../api.json")
	writer := NewRequestRuleWriter("../app/Http/Requests/", "//rule_start", "//rule_end")
	Writes(parser, writer)
}
func TestRequestAttributeWriter(t *testing.T) {
	parser := NewRoutesParser("../api.json")
	writer := NewRequestAttributeWriter("../app/Http/Requests/", "//attribute_start", "//attribute_end")
	Writes(parser, writer)
}
