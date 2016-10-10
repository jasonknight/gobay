package gobay

func SetNotificationPreferencesTemplate() string {
	return `{{ with .ApplicationDeliveryPreferences }}<ApplicationDeliveryPreferences> 
    <AlertEmail>{{ .AlertEmail }}</AlertEmail>
    <AlertEnable>{{ .AlertEnable }}</AlertEnable>
    <ApplicationEnable>{{ .ApplicationEnable }}</ApplicationEnable>
    <ApplicationURL>{{ .ApplicationURL }}</ApplicationURL>
    {{ range .DeliveryURLDetails }}
    <DeliveryURLDetails>
      <DeliveryURL>{{ .DeliveryURL }}</DeliveryURL>
      <DeliveryURLName>{{ .DeliveryURLName }}</DeliveryURLName>
      <Status>{{ .Status }}</Status>
    </DeliveryURLDetails>
    {{ end }}
    <DeviceType>{{ .DeviceType }}</DeviceType>
    <PayloadVersion>{{ .PayloadVersion }}</PayloadVersion>
  </ApplicationDeliveryPreferences>{{end}}
  <DeliveryURLName>{{ .DeliveryURLName }}</DeliveryURLName>
  {{ range .EventProperties }}
  <EventProperty>
    <EventType>{{ .EventType }}</EventType>
    <Name>{{ .Name }}</Name>
    <Value>{{ .Value }}</Value>
  </EventProperty>
  {{end}}
  {{ with .UserData }}
  <UserData>
    <ExternalUserData>{{ .ExternalUserData }}</ExternalUserData>
    {{ if ne .SMSSubscription.CarrierID "" }}
    {{ with .SMSSubscription }}
    <SMSSubscription>
      <CarrierID>{{ .CarrierID }}</CarrierID>
      <ErrorCode>{{ .ErrorCode }}</ErrorCode>
      <ItemToUnsubscribe>{{ .ItemToUnsubscribe }}</ItemToUnsubscribe>
      <SMSPhone>{{ .SMSPhone }}</SMSPhone>
      <UserStatus>{{ .UserStatus }}</UserStatus>
    </SMSSubscription>
    {{ end }}
    {{ end }}
    {{ range .SummarySchedules }}
    <SummarySchedule>
      <EventType>{{ .EventType }}</EventType>
      <Frequency>{{ .Frequency }}</Frequency>
      <SummaryPeriod>{{ .SummaryPeriod }}</SummaryPeriod>
    </SummarySchedule>
  {{ end }}
  </UserData>
  {{ end }}
  {{ with .UserDeliveryPreferenceArray }}
  <UserDeliveryPreferenceArray>
    {{ range .EnabledNotifications }}
      <NotificationEnable>
        <EventEnable>{{ .EventEnable }}</EventEnable>
        <EventType>{{ .EventType }}</EventType>
      </NotificationEnable>
    {{end}}
  </UserDeliveryPreferenceArray>
  {{ end }}
`
}
