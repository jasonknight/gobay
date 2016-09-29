package gobay

import "testing"

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
