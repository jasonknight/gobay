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
	var results GenericResults

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

	for _, r := range results.Results {
		if r.Ack != "Success" {
			t.Errorf("GeteBayOfficialTime ebay failed %+v\n", r)
		}
	}

}

func TestGeteBayDetails(t *testing.T) {
	if shouldRunSandbox() == false {
		return
	}
	var results GenericResults

	cnf, err := fileGetContents("../secret.yml")

	if err != nil {
		t.Errorf("Failed to load secret.yml %v\n", err)
	}

	ebay, err := NewEbayCallEx(cnf)

	if err != nil {
		t.Errorf("Failed to initialize ebay call from secret.yml %v\n", err)
	}

	ebay.SetCallname("GeteBayDetails")
	ebay.EbayDetailsCallInfo = NewGetEbayDetailsStruct()
	err = ebay.Execute(&results)
	if err != nil {
		t.Errorf("Failed to Execute %v\n", err)
	}

	for _, r := range results.Results {
		if r.Ack != "Success" {
			t.Errorf("GeteBayDetails failed %+v\n", r)
		}
	}

}

func TestGetAllCategories(t *testing.T) {
	are_you_sure := "NO"
	if are_you_sure != "YES" {
		return
	}
	fmt.Printf("Going ahead with GetAllCategories test!\n")
	var results GenericResults

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

	for _, r := range results.Results {
		if r.Ack != "Success" {
			t.Errorf("GetAllCategories failed %+v\n", r)
		}
	}

}

func TestGetMyeBaySelling(t *testing.T) {
	if shouldRunSandbox() == false {
		return
	}
	fmt.Printf("Going ahead with TestGetMyeBaySelling!\n")
	var results GenericResults

	cnf, err := fileGetContents("../secret.yml")

	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}

	ebay, err := NewEbayCallEx(cnf)

	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}

	ebay.SetCallname("GetMyeBaySelling")

	ebay.MyeBaySellingCallInfo = NewMyeBaySelling()
	ebay.MyeBaySellingCallInfo.AddActiveList()

	err = ebay.Execute(&results)
	if err != nil {
		t.Errorf("Failed to Execute %v\n", err)
	}

	for _, r := range results.Results {
		if r.Ack != "Success" {
			t.Errorf("GetMyeBaySelling failed %+v\n", r)
		}
	}
}

func TestGetNotificationPreferences(t *testing.T) {
	if shouldRunSandbox() == false {
		return
	}
	fmt.Printf("Going ahead with TestGetNotificationPreferences!\n")
	var results GenericNotificationPreferenceResults

	cnf, err := fileGetContents("../secret.yml")

	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}

	ebay, err := NewEbayCallEx(cnf)

	if err != nil {
		t.Errorf("Failed to load test.yml %v\n", err)
	}

	ebay.SetCallname("GetNotificationPreferences")
	ebay.NotificationPreferencesCallInfo = NewNotificationPreferencesRequest()
	ebay.NotificationPreferencesCallInfo.PreferenceLevel = "Application"

	err = ebay.Execute(&results)
	if err != nil {
		t.Errorf("Failed to Execute %v\n", err)
		return
	}

	for _, r := range results.Results {
		if r.Ack != "Success" {
			t.Errorf("GetMyeBaySelling failed %+v\n", r)
		}
	}
}

func TestSetNotificationPreferences(t *testing.T) {
	if shouldRunSandbox() == false {
		return
	}

	data, err := fileGetContents("../setnots.xml")
	if err != nil {
		t.Errorf("failed to load secret file %s", err)
		return
	}

	fmt.Printf("Going ahead with TestGetNotificationPreferences!\n")
	var results GenericResults

	cnf, err := fileGetContents("../secret.yml")

	if err != nil {
		t.Errorf("Failed to load secret.yml %v\n", err)
	}

	ebay, err := NewEbayCallEx(cnf)

	if err != nil {
		t.Errorf("Failed to create new ebay call%v\n", err)
	}

	ebay.SetCallname("SetNotificationPreferences")
	info := NewSetNotificationPreferencesRequest()
	info.FromXML([]byte(data))
	ebay.SetNotificationPreferencesCallInfo = info

	err = ebay.Execute(&results)
	if err != nil {
		t.Errorf("Failed to Execute %v\n", err)
		return
	}

	for _, r := range results.Results {
		if r.Ack != "Success" {
			t.Errorf("SetNotificationPreferences failed %+v\n", r)
		}
	}
}

func TestSiteIDToCode(t *testing.T) {
	sid := "3"
	scode := SiteIDToCode(sid)

	if scode != "UK" {
		t.Errorf("SiteIDToCode failed")
		return
	}
}
