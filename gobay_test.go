package gobay

import "testing"

import "fmt"

var runSandboxTests bool


func shouldRunSandbox() bool {
	runSandboxTests = false // turn this to true if you have
	// setup your sandbox account
	if fileExists("../secret.yml") == false {
		fmt.Printf("Not Running sandbox tests because no secret\n")
		return false
	}
	if runSandboxTests == false {
		fmt.Printf("Not Running sandbox tests because of settings\n")
	}
	return runSandboxTests
}

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
	if shouldRunSandbox() == false {
		return
	}
	var results []Result
	
	cnf, err := fileGetContents("../secret.yml")
	
	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}
	
	call, err := NewEbayCallEx(cnf)

	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}

    
	call.SetCallname("GeteBayOfficialTime")
    call.Execute(&results)

    for _,r := range results {
    	if r.Ack != "Success" {
    		t.Errorf("GeteBayOfficialTime call failed %+v\n", r)
    	}
    }

}
