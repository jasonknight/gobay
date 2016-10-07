package gobay
import "testing"
func TestSetNotificationPreferencesTemplate(t *testing.T) {
	tmpl := SetNotificationPreferencesTemplate()
	o := NewSetNotificationPreferencesRequest()
	data,err := compileGoString("TestSetNotificationPreferencesTemplate",tmpl,o,nil)
	if err != nil {
		t.Errorf("failed to compile SetNotificationPreferencesTemplate %v",err)
	}
	if data == "" {
		t.Errorf("failed to compite SetNotificationPreferencesTemplate %v",data)
	}
}
