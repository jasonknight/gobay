package gobay

import "testing"

func TestNotificationPreferencesTemplate(t *testing.T) {
	tmpl := NotificationPreferencesTemplate()
	o := NewNotificationPreferencesRequest()
	data, err := compileGoString("TestNotificationPreferencesTemplate", tmpl, o, nil)
	if err != nil {
		t.Errorf("failed to compile NotificationPreferencesTemplate %v", err)
	}
	if data == "" {
		t.Errorf("failed to compite NotificationPreferencesTemplate %v", data)
	}
}
