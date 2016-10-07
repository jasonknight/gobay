package gobay
import "testing"
func TestSetNotificationPreferencesRequest ( t *testing.T ) {
	data := `<?xml version="1.0" encoding="utf-8"?>
<SetNotificationPreferencesRequest xmlns="urn:ebay:apis:eBLBaseComponents">
  <!-- Call-specific Input Fields -->
  <ApplicationDeliveryPreferences> 
    <AlertEmail>anyURI</AlertEmail>
    <AlertEnable>EnableCodeType</AlertEnable>
    <ApplicationEnable>EnableCodeType</ApplicationEnable>
    <ApplicationURL>anyURI</ApplicationURL>
    <DeliveryURLDetails>
      <DeliveryURL>anyURI</DeliveryURL>
      <DeliveryURLName>string</DeliveryURLName>
      <Status>EnableCodeType</Status>
    </DeliveryURLDetails>
    <DeliveryURLDetails>
      <DeliveryURL>anyURI</DeliveryURL>
      <DeliveryURLName>string</DeliveryURLName>
      <Status>EnableCodeType</Status>
    </DeliveryURLDetails>
    <!-- ... more DeliveryURLDetails nodes allowed here ... -->
    <DeviceType>DeviceTypeCodeType</DeviceType>
    <PayloadVersion>1234</PayloadVersion>
  </ApplicationDeliveryPreferences>
  <DeliveryURLName>string</DeliveryURLName>
  <EventProperty>
    <EventType>NotificationEventTypeCodeType</EventType>
    <Name>NotificationEventPropertyNameCodeType</Name>
    <Value>string</Value>
  </EventProperty>
  <EventProperty>
    <EventType>NotificationEventTypeCodeType</EventType>
    <Name>NotificationEventPropertyNameCodeType</Name>
    <Value>string</Value>
  </EventProperty>
  <!-- ... more EventProperty nodes allowed here ... -->
  <UserData>
    <ExternalUserData>string</ExternalUserData>
    <SMSSubscription>
      <CarrierID>WirelessCarrierIDCodeType</CarrierID>
      <ErrorCode>SMSSubscriptionErrorCodeCodeType</ErrorCode>
      <ItemToUnsubscribe>ItemIDType (string)</ItemToUnsubscribe>
      <SMSPhone>string</SMSPhone>
      <UserStatus>SMSSubscriptionUserStatusCodeType</UserStatus>
    </SMSSubscription>
    <SummarySchedule>
      <EventType>NotificationEventTypeCodeType</EventType>
      <Frequency>SummaryFrequencyCodeType</Frequency>
      <SummaryPeriod>SummaryWindowPeriodCodeType</SummaryPeriod>
    </SummarySchedule>
    <SummarySchedule>
      <EventType>NotificationEventTypeCodeType</EventType>
      <Frequency>SummaryFrequencyCodeType</Frequency>
      <SummaryPeriod>SummaryWindowPeriodCodeType</SummaryPeriod>
    </SummarySchedule>
    <!-- ... more SummarySchedule nodes allowed here ... -->
  </UserData>
  <UserDeliveryPreferenceArray>
    <NotificationEnable>
      <EventEnable>EnableCodeType</EventEnable>
      <EventType>NotificationEventTypeCodeType</EventType>
    </NotificationEnable>
    <NotificationEnable>
      <EventEnable>EnableCodeType</EventEnable>
      <EventType>NotificationEventTypeCodeType</EventType>
    </NotificationEnable>
    <!-- ... more NotificationEnable nodes allowed here ... -->
  </UserDeliveryPreferenceArray>
</SetNotificationPreferencesRequest>`
	var o SetNotificationPreferencesRequest
	err := o.FromXML([]byte(data))
	if err != nil {
		t.Errorf("failed loading xml %v",err)
		return
	}
	if o.ApplicationDeliveryPreferences.AlertEmail != "anyURI" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.AlertEmail != anyURI %v",o)
		return
	}

	if o.ApplicationDeliveryPreferences.AlertEnable != "EnableCodeType" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.AlertEnable != EnableCodeType")
		return
	}

	if o.ApplicationDeliveryPreferences.ApplicationEnable != "EnableCodeType" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.ApplicationEnable != EnableCodeType")
		return
	}

	if o.ApplicationDeliveryPreferences.ApplicationURL != "anyURI" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.ApplicationURL != anyURI")
		return
	}

	if o.ApplicationDeliveryPreferences.DeliveryURLDetails[0].DeliveryURL != "anyURI" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.DeliveryURLDetails[0].DeliveryURL != anyURI")
		return
	}

	if o.ApplicationDeliveryPreferences.DeliveryURLDetails[0].DeliveryURLName != "string" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.DeliveryURLDetails[0].DeliveryURLName != string")
		return
	}

	if o.ApplicationDeliveryPreferences.DeliveryURLDetails[0].Status != "EnableCodeType" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.DeliveryURLDetails[0].Status != EnableCodeType")
		return
	}


	if o.ApplicationDeliveryPreferences.DeviceType != "DeviceTypeCodeType" {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.DeviceType != DeviceTypeCodeType")
		return
	}

	if o.ApplicationDeliveryPreferences.PayloadVersion != 1234 {
		t.Errorf("failed because o.ApplicationDeliveryPreferences.PayloadVersion != string")
		return
	}


	if o.DeliveryURLName != "string" {
		t.Errorf("failed because o.DeliveryURLName != string")
		return
	}

	if o.EventProperty[0].EventType != "NotificationEventTypeCodeType" {
		t.Errorf("failed because o.EventProperty[0].EventType != NotificationEventTypeCodeType")
		return
	}

	if o.EventProperty[0].Name != "NotificationEventPropertyNameCodeType" {
		t.Errorf("failed because o.EventProperty[0].Name != NotificationEventPropertyNameCodeType")
		return
	}

	if o.EventProperty[0].Value != "string" {
		t.Errorf("failed because o.EventProperty[0].Value != string")
		return
	}


	if o.UserData.ExternalUserData != "string" {
		t.Errorf("failed because o.UserData.ExternalUserData != string")
		return
	}

	if o.UserData.SMSSubscription.CarrierID != "WirelessCarrierIDCodeType" {
		t.Errorf("failed because o.UserData.SMSSubscription.CarrierID != WirelessCarrierIDCodeType")
		return
	}

	if o.UserData.SMSSubscription.ErrorCode != "SMSSubscriptionErrorCodeCodeType" {
		t.Errorf("failed because o.UserData.SMSSubscription.ErrorCode != SMSSubscriptionErrorCodeCodeType")
		return
	}

	if o.UserData.SMSSubscription.ItemToUnsubscribe != "ItemIDType (string)" {
		t.Errorf("failed because o.UserData.SMSSubscription.ItemToUnsubscribe != ItemIDType (string) %+v",o.UserData.SMSSubscription)
		return
	}

	if o.UserData.SMSSubscription.SMSPhone != "string" {
		t.Errorf("failed because o.UserData.SMSSubscription.SMSPhone != string")
		return
	}

	if o.UserData.SMSSubscription.UserStatus != "SMSSubscriptionUserStatusCodeType" {
		t.Errorf("failed because o.UserData.SMSSubscription.UserStatus != SMSSubscriptionUserStatusCodeType")
		return
	}


	if o.UserData.SummarySchedule[0].EventType != "NotificationEventTypeCodeType" {
		t.Errorf("failed because o.UserData.SummarySchedule[0].EventType != NotificationEventTypeCodeType")
		return
	}

	if o.UserData.SummarySchedule[0].Frequency != "SummaryFrequencyCodeType" {
		t.Errorf("failed because o.UserData.SummarySchedule[0].Frequency != SummaryFrequencyCodeType")
		return
	}

	if o.UserData.SummarySchedule[0].SummaryPeriod != "SummaryWindowPeriodCodeType" {
		t.Errorf("failed because o.UserData.SummarySchedule[0].SummaryPeriod != SummaryWindowPeriodCodeType")
		return
	}



	if o.UserDeliveryPreferenceArray.NotificationEnable[0].EventEnable != "EnableCodeType" {
		t.Errorf("failed because o.UserDeliveryPreferenceArray.NotificationEnable[0].EventEnable != EnableCodeType")
		return
	}

	if o.UserDeliveryPreferenceArray.NotificationEnable[0].EventType != "NotificationEventTypeCodeType" {
		t.Errorf("failed because o.UserDeliveryPreferenceArray.NotificationEnable[0].EventType != NotificationEventTypeCodeType")
		return
	}



}
