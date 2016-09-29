package gobay

import "testing"
import "encoding/xml"
import "fmt"

func TestItemTemplate(t *testing.T) {
	data := ItemTemplate()
	cnf, err := fileGetContents("test_data/test.yml")
	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}
	call, err := NewEbayCallEx(cnf)
	p := call.NewItem()
	p.Title = "xxx123"
	p.StartPrice = "33.28"
	s, err := compileGoString("ItemTemplate", data, p, nil)
	if err != nil {
		t.Errorf("Failed to compile template %v\n", err)
	}
	data = `<?xml version="1.0" encoding="UTF-8"?>`
	data = fmt.Sprintf("%s%s", data, s)
	np := NewItem()
	xml.Unmarshal([]byte(data), &np)
	if np.Title != p.Title && np.StartPrice == p.StartPrice {
		t.Errorf("Failed reverse template \n%v\n%v", p, np)
	}
}
