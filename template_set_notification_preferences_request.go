package gobay
func SetNotificationPreferencesTemplate() string {
	return`
  <ApplicationDeliveryPreferences> 
    <AlertEmail> anyURI </AlertEmail>
    <AlertEnable> EnableCodeType </AlertEnable>
    <ApplicationEnable> EnableCodeType </ApplicationEnable>
    <ApplicationURL> anyURI </ApplicationURL>
    <DeliveryURLDetails>
      <DeliveryURL> anyURI </DeliveryURL>
      <DeliveryURLName> string </DeliveryURLName>
      <Status> EnableCodeType </Status>
    </DeliveryURLDetails>
    <DeliveryURLDetails>
      <DeliveryURL> anyURI </DeliveryURL>
      <DeliveryURLName> string </DeliveryURLName>
      <Status> EnableCodeType </Status>
    </DeliveryURLDetails>
    <!-- ... more DeliveryURLDetails nodes allowed here ... -->
    <DeviceType> DeviceTypeCodeType </DeviceType>
    <PayloadVersion> string </PayloadVersion>
  </ApplicationDeliveryPreferences>
  <DeliveryURLName> string </DeliveryURLName>
  <EventProperty>
    <EventType> NotificationEventTypeCodeType </EventType>
    <Name> NotificationEventPropertyNameCodeType </Name>
    <Value> string </Value>
  </EventProperty>
  <EventProperty>
    <EventType> NotificationEventTypeCodeType </EventType>
    <Name> NotificationEventPropertyNameCodeType </Name>
    <Value> string </Value>
  </EventProperty>
  <!-- ... more EventProperty nodes allowed here ... -->
  <UserData>
    <ExternalUserData> string </ExternalUserData>
    <SMSSubscription>
      <CarrierID> WirelessCarrierIDCodeType </CarrierID>
      <ErrorCode> SMSSubscriptionErrorCodeCodeType </ErrorCode>
      <ItemToUnsubscribe> ItemIDType (string) </ItemToUnsubscribe>
      <SMSPhone> string </SMSPhone>
      <UserStatus> SMSSubscriptionUserStatusCodeType </UserStatus>
    </SMSSubscription>
    <SummarySchedule>
      <EventType> NotificationEventTypeCodeType </EventType>
      <Frequency> SummaryFrequencyCodeType </Frequency>
      <SummaryPeriod> SummaryWindowPeriodCodeType </SummaryPeriod>
    </SummarySchedule>
    <SummarySchedule>
      <EventType> NotificationEventTypeCodeType </EventType>
      <Frequency> SummaryFrequencyCodeType </Frequency>
      <SummaryPeriod> SummaryWindowPeriodCodeType </SummaryPeriod>
    </SummarySchedule>
    <!-- ... more SummarySchedule nodes allowed here ... -->
  </UserData>
  <UserDeliveryPreferenceArray>
    <NotificationEnable>
      <EventEnable> EnableCodeType </EventEnable>
      <EventType> NotificationEventTypeCodeType </EventType>
    </NotificationEnable>
    <NotificationEnable>
      <EventEnable> EnableCodeType </EventEnable>
      <EventType> NotificationEventTypeCodeType </EventType>
    </NotificationEnable>
    <!-- ... more NotificationEnable nodes allowed here ... -->
  </UserDeliveryPreferenceArray>
`
}
