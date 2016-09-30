package gobay

import "testing"
import "fmt"

func TestLoadingFixture(t *testing.T) {
	var i Item
	cnf, err := fileGetContents("test_data/product_1.yml")
	if err != nil {
		t.Errorf("Failed to load product_1.yml %v\n", err)
		return
	}
	err = i.FromYAML(cnf)
	if err != nil {
		t.Errorf("Failed FromYAML %v\n", err)
		return
	}

	if i.Title != "My Unique Title" {
		t.Errorf("Title was not set: %+v\n", i)
		return
	}
}

func TestCollectAddItems(t *testing.T) {
	cnf, err := fileGetContents("test_data/test.yml")
	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}

	ebay, err := NewEbayCallEx(cnf)

	if err != nil {
		t.Errorf("Failed to initialize ebay call from test.yml %v\n", err)
	}
	oi := ebay.NewItem()

	pcnf, err := fileGetContents("test_data/product_1.yml")
	if err != nil {
		t.Errorf("Failed to load product_1.yml %v\n", err)
		return
	}
	err = oi.FromYAML(pcnf)

	if err != nil {
		t.Errorf("Failed to initialize item from product_1.yml %v\n", err)
		return
	}

	ebay.AddItem(oi)

	for i := 1; i < 11; i++ {
		ni := oi.Clone()
		ni.Title = fmt.Sprintf("%s-%d", ni.Title, i)
		ni.SKU = fmt.Sprintf("%s-%d", ni.SKU, i)
		ebay.AddItem(ni)
	}

	s, err := ebay.CollectAddItems()

	if err != nil {
		t.Errorf("ebay.CollectAddItems failed! %v\n", err)
		return
	}

	if len(s.Children) != ebay.AddItemsLimit {
		t.Errorf("ebay.CollectAddItems failed! %v\n", err)
		return
	}

	for _, c := range s.Children {
		if c.MessageID != c.Item.GetInternalID() {
			t.Errorf("MessageID should == item's internal_id")
		}
		c.Item.sent = true
	}
	os := *s // we'll use this momentarily

	s.Children = nil

	if len(s.Children) > 0 {
		t.Errorf("Children should be empty! %v\n", err)
		return
	}

	s, err = ebay.CollectAddItems()

	if len(s.Children) < 1 {
		t.Errorf("Children should NOT be empty! %v\n", err)
		return
	}

	if len(s.Children) != ebay.AddItemsLimit {
		t.Errorf("ebay.CollectAddItems #2 failed! %v\n", err)
		return
	}

	if s.Children[0].Item.Title == os.Children[0].Item.Title {
		t.Errorf("ebay.CollectAddItems failed, first item shouldn't have the same Title! %v\n", err)
		return
	}

	ebay.XMLData, err = ebay.CollectAddItemsXML(s)

	if err != nil {
		t.Errorf("CollectAddItemsXML has error %s", err)
		return
	}

	if ebay.XMLData == "" {
		t.Errorf("XMLData is empty!")
	}

}
func TestCollectAddItemsLength(t *testing.T) {
	// if shouldRunSandbox() == false {
	// 	return
	// }
	//var results []Result

	cnf, err := fileGetContents("test_data/test.yml")

	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}

	ebay, err := NewEbayCallEx(cnf)

	if err != nil {
		t.Errorf("Failed to initialize ebay call from test.yml %v\n", err)
	}

	i := ebay.NewItem()

	pcnf, err := fileGetContents("test_data/product_1.yml")
	if err != nil {
		t.Errorf("Failed to load product_1.yml %v\n", err)
		return
	}
	err = i.FromYAML(pcnf)

	if err != nil {
		t.Errorf("Failed FromYAML %v\n", err)
		return
	}
	ebay.SetCallname("AddItems")
	// Get an EbayCall

	ebay.AddItem(i)

	check, err := ebay.CollectAddItems()

	if err != nil {
		t.Errorf(fmt.Sprintf("wtf%s", err))
		return
	}

	if len(check.Children) > len(ebay.Items) {
		t.Errorf(fmt.Sprintf("Children: %+v", check.Children))
		return
	}
}

func TestAddItem(t *testing.T) {
	if shouldRunSandbox() == false {
		return
	}
	var results []Result

	cnf, err := fileGetContents("../secret.yml")

	if err != nil {
		t.Errorf("Failed to load secret.yml %v\n", err)
	}

	ebay, err := NewEbayCallEx(cnf)

	if err != nil {
		t.Errorf("Failed to initialize ebay call from secret.yml %v\n", err)
	}

	i := ebay.NewItem()

	pcnf, err := fileGetContents("test_data/product_1.yml")
	if err != nil {
		t.Errorf("Failed to load product_1.yml %v\n", err)
		return
	}
	err = i.FromYAML(pcnf)
	i.SKU, _ = pseudoUUID()
	i.Title, _ = pseudoUUID()

	i.SKU = i.SKU[0:5]
	i.Title = fmt.Sprintf("This is a simple title %s", i.Title[0:5])

	if err != nil {
		t.Errorf("Failed FromYAML %v\n", err)
		return
	}
	ebay.SetCallname("AddItems")
	// Get an EbayCall

	ebay.AddItem(i)

	// check := ebay.CollectAddItems()

	// if len(check.Children) > len(ebay.Items) {
	// 	t.Errorf("What the fuck?")
	// }

	ebay.Execute(&results)

	for _, r := range results {
		if r.Failure() {
			for _, e := range r.Errors {
				t.Errorf("%s\n", e.LongMessage)
				for _, ep := range e.ErrorParameters {
					t.Errorf("\t%s\n", ep.Value)
				}
			}
			t.Errorf("%+v", r.Errors)
			return
		}
	}
}
