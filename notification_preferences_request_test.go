package gobay
import "testing"
func TestNotificationPreferencesRequest ( t *testing.T ) {
	data := `<?xml version="1.0" encoding="utf-8"?>
<GetNotificationPreferencesRequest xmlns="urn:ebay:apis:eBLBaseComponents">
<PreferenceLevel>User</PreferenceLevel>
</GetNotificationPreferencesRequest>`
	var o NotificationPreferencesRequest
	err := o.FromXML([]byte(data))
	if err != nil {
		t.Errorf("failed loading xml %v",err)
		return
	}
	if o.PreferenceLevel != "User" {
		t.Errorf("failed because o.PreferenceLevel != User")
		return
	}

}
