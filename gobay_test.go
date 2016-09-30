package gobay

import "testing"

import "fmt"


func TestNewEbayCallEx(t *testing.T) {
	cnf, err := fileGetContents("test_data/test.yml")
	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}
	ebay, err := NewEbayCallEx(cnf)
	if err != nil {
		t.Errorf("Failed to create ebay call  %v\n", err)
	}

	if ebay.DevID != "1234567" {
		t.Errorf("TestNewEbayCallEx DevID is not properly set!\n")
	}
}

func TestLoadConfiguration(t *testing.T) {
	// this is an empty test
	cnf, err := fileGetContents("test_data/test.yml")
	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}
	m := make(map[interface{}]interface{})
	err = LoadConfiguration(cnf, &m)

	if err != nil {
		t.Errorf("Failed to apply test.yml %v\n", err)
	}

	if m["DevID"] != "1234567" {
		t.Errorf("TestLoadConfiguration DevID is not properly set!\n")
	}
}

func TestGetTime(t *testing.T) {
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

	ebay.SetCallname("GeteBayOfficialTime")
	err = ebay.Execute(&results)
	if err != nil {
		t.Errorf("Failed to Execute %v\n", err)
	}

	for _, r := range results {
		if r.Ack != "Success" {
			t.Errorf("GeteBayOfficialTime ebay failed %+v\n", r)
		}
	}

}
func TestAddItemsXML(t *testing.T) {
	cnf, err := fileGetContents("test_data/test.yml")
	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}
	ebay, err := NewEbayCallEx(cnf)

	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}
	ebay.SetCallname("AddItems")
}

func TestGetAllCategories(t *testing.T) {
	are_you_sure := "NO"
	if are_you_sure != "YES" {
		return
	}
	fmt.Printf("Going ahead with GetAllCategories test!\n")
	var results []Result

	cnf, err := fileGetContents("../secret.yml")

	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}

	ebay, err := NewEbayCallEx(cnf)

	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}

	ebay.SetCallname("GetAllCategories")

	ebay.CategoryCallInfo = NewGetCategoriesStruct()

	err = ebay.Execute(&results)
	if err != nil {
		t.Errorf("Failed to Execute %v\n", err)
	}

	for _, r := range results {
		if r.Ack != "Success" {
			t.Errorf("GetAllCategories failed %+v\n", r)
		}
	}

}
