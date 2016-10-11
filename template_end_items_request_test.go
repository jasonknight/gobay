package gobay

import "testing"

func TestEndItemsTemplate(t *testing.T) {
	tmpl := EndItemsTemplate()
	o := NewEndItemsRequest()
	data, err := compileGoString("TestEndItemsTemplate", tmpl, o, nil)
	if err != nil {
		t.Errorf("failed to compile EndItemsTemplate %v", err)
	}
	if data == "" {
		t.Errorf("failed to compite EndItemsTemplate %v", data)
	}
}
