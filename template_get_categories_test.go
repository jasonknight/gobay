package gobay

import "testing"
import "encoding/xml"
import "fmt"

func TestGetAllCategoriesTemplate(t *testing.T) {
	data := GetAllCategoriesTemplate()
	var o GetCategoriesStruct
	var o2 GetCategoriesStruct

	o.SiteID = "UK"
	o.LevelLimit = "5"
	o.ViewAllNodes = "true"
	o.OutputSelectors = append(o.OutputSelectors, "TestSelector")

	s, err := compileGoString("GetAllCategoriesTemplate", data, o, nil)
	if err != nil {
		t.Errorf("Failed to compile template %v\n", err)
	}
	data = `<?xml version="1.0" encoding="UTF-8"?>`
	data = fmt.Sprintf("%s<Root>%s</Root>", data, s)
	xml.Unmarshal([]byte(data), &o2)
	if o2.SiteID != "5" && o2.ViewAllNodes != "true" && o2.OutputSelectors[0] != "TestSelector" {
		t.Errorf("Failed reverse template \n%v\n%v", o, o2)
	}

}
func TestGetChildCategoriesTemplate(t *testing.T) {

}
