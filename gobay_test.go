package gobay

import "testing"

import "fmt"

func TestNewEbayCallEx(t *testing.T) {
	cnf, err := fileGetContents("test_data/test.yml")
	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}
	call, err := NewEbayCallEx(cnf)

	if call.DevID != "1234567" {
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
	// This test only runs if you've filled out
	// a yaml file with actual sandbox data
	if fileExists("../secret.yml") == false {
		return
	}
	fmt.Printf("Running sandbox test\n")
	cnf, err := fileGetContents("../secret.yml")
	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}
	call, err := NewEbayCallEx(cnf)

    var results []Result
	call.SetCallname("GeteBayOfficialTime")
    call.Execute(&results)
    fmt.Printf("RESULTS: %+v",results)

}
