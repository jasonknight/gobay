package gobay

import "testing"

func TestNotificationPreferencesResult(t *testing.T) {
	data := `<?xml version="1.0" encoding="UTF-8"?>
<GetNotificationPreferencesResponse xmlns="urn:ebay:apis:eBLBaseComponents">
  <Timestamp>2010-12-30T00:26:04.640Z</Timestamp>
  <Ack>Success</Ack>
  <Version>699</Version>
  <Build>E699_CORE_BUNDLED_12457306_R1</Build>
  <ApplicationDeliveryPreferences>
    <ApplicationURL>mailto://magicalbookseller@yahoo.com</ApplicationURL>
    <ApplicationEnable>Enable</ApplicationEnable>
    <AlertEmail>mailto://magicalbookseller@yahoo.com</AlertEmail>
    <AlertEnable>Enable</AlertEnable>
    <NotificationPayloadType>eBLSchemaSOAP</NotificationPayloadType>
    <DeviceType>Platform</DeviceType>
    <PayloadEncodingType>JSON</PayloadEncodingType>
    <PayloadVersion>557</PayloadVersion>
  </ApplicationDeliveryPreferences>
  <UserDeliveryPreferenceArray>
    <NotificationEnable>
      <EventType>EndOfAuction</EventType>
      <EventEnable>Enable</EventEnable>
    </NotificationEnable>
    <NotificationEnable>
      <EventType>Feedback</EventType>
      <EventEnable>Enable</EventEnable>
    </NotificationEnable>
    <NotificationEnable>
      <EventType>ItemListed</EventType>
      <EventEnable>Enable</EventEnable>
    </NotificationEnable>
    <NotificationEnable>
      <EventType>MyMessageseBayMessageHeader</EventType>
      <EventEnable>Enable</EventEnable>
    </NotificationEnable>
    <NotificationEnable>
      <EventType>BidReceived</EventType>
      <EventEnable>Enable</EventEnable>
    </NotificationEnable>
  </UserDeliveryPreferenceArray>
</GetNotificationPreferencesResponse>`
	var o NotificationPreferencesResult
	err := o.FromXML([]byte(data))
	if err != nil {
		t.Errorf("failed loading xml %v", err)
		return
	}
	if o.Timestamp != "2010-12-30T00:26:04.640Z" {
		t.Errorf("failed because o.Timestamp != 2010-12-30T00:26:04.640Z")
		return
	}

	if o.Ack != "Success" {
		t.Errorf("failed because o.Ack != Success")
		return
	}

	if o.Version != 699 {
		t.Errorf("failed because o.Version != 699")
		return
	}

	if o.Build != "E699_CORE_BUNDLED_12457306_R1" {
		t.Errorf("failed because o.Build != E699_CORE_BUNDLED_12457306_R1")
		return
	}

	if o.ApplicationDeliveryPreferences.ApplicationURL != "mailto://magicalbookseller@yahoo.com" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.ApplicationURL != mailto://magicalbookseller@yahoo.com")
		return
	}

	if o.ApplicationDeliveryPreferences.ApplicationEnable != "Enable" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.ApplicationEnable != Enable")
		return
	}

	if o.ApplicationDeliveryPreferences.AlertEmail != "mailto://magicalbookseller@yahoo.com" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.AlertEmail != mailto://magicalbookseller@yahoo.com")
		return
	}

	if o.ApplicationDeliveryPreferences.AlertEnable != "Enable" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.AlertEnable != Enable")
		return
	}

	if o.ApplicationDeliveryPreferences.NotificationPayloadType != "eBLSchemaSOAP" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.NotificationPayloadType != eBLSchemaSOAP")
		return
	}

	if o.ApplicationDeliveryPreferences.DeviceType != "Platform" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.DeviceType != Platform")
		return
	}

	if o.ApplicationDeliveryPreferences.PayloadEncodingType != "JSON" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.PayloadEncodingType != JSON")
		return
	}

	if o.ApplicationDeliveryPreferences.PayloadVersion != 557 {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.PayloadVersion != 557")
		return
	}

	if o.UserDeliveryPreferenceArray.NotificationEnable[0].EventType != "EndOfAuction" {
		t.Errorf("failed because o.UserDeliveryPreferenceArray.NotificationEnable[0].EventType != EndOfAuction")
		return
	}

	if o.UserDeliveryPreferenceArray.NotificationEnable[0].EventEnable != "Enable" {
		t.Errorf("failed because o.UserDeliveryPreferenceArray.NotificationEnable[0].EventEnable != Enable")
		return
	}

}
