package gobay

import (
	"encoding/xml"
	"gopkg.in/yaml.v2"
)

type SetNotificationPreferencesRequest struct {
	ApplicationDeliveryPreferences ApplicationDeliveryPreference `xml:"ApplicationDeliveryPreferences" yaml:"ApplicationDeliveryPreferences"`
	DeliveryURLName                string                        `xml:"DeliveryURLName" yaml:"DeliveryURLName"`
	EventProperties                  []EventProperty               `xml:"EventProperty" yaml:"EventProperty"`
	UserData                       UserData                      `xml:"UserData" yaml:"UserData"`
	UserDeliveryPreferenceArray    UserDeliveryPreference        `xml:"UserDeliveryPreferenceArray" yaml:"UserDeliveryPreferenceArray"`
}

func NewSetNotificationPreferencesRequest() *SetNotificationPreferencesRequest {
	return &SetNotificationPreferencesRequest{}
}

func (o *SetNotificationPreferencesRequest) FromXML(data []byte) error {
	return xml.Unmarshal(data, o)
}

func (o *SetNotificationPreferencesRequest) FromYAML(data []byte) error {
	return yaml.Unmarshal(data, o)
}

type DeliveryURLDetail struct {
	DeliveryURL     string `xml:"DeliveryURL" yaml:"DeliveryURL"`
	DeliveryURLName string `xml:"DeliveryURLName" yaml:"DeliveryURLName"`
	Status          string `xml:"Status" yaml:"Status"`
}

type EventProperty struct {
	EventType string `xml:"EventType" yaml:"EventType"`
	Name      string `xml:"Name" yaml:"Name"`
	Value     string `xml:"Value" yaml:"Value"`
}

type UserData struct {
	ExternalUserData string            `xml:"ExternalUserData" yaml:"ExternalUserData"`
	SMSSubscription  SMSSubscription   `xml:"SMSSubscription" yaml:"SMSSubscription"`
	SummarySchedules  []SummarySchedule `xml:"SummarySchedule" yaml:"SummarySchedule"`
}

type SMSSubscription struct {
	CarrierID         string `xml:"CarrierID" yaml:"CarrierID"`
	ErrorCode         string `xml:"ErrorCode" yaml:"ErrorCode"`
	ItemToUnsubscribe string `xml:"ItemToUnsubscribe" yaml:"ItemToUnsubscribe"`
	SMSPhone          string `xml:"SMSPhone" yaml:"SMSPhone"`
	UserStatus        string `xml:"UserStatus" yaml:"UserStatus"`
}

type SummarySchedule struct {
	EventType     string `xml:"EventType" yaml:"EventType"`
	Frequency     string `xml:"Frequency" yaml:"Frequency"`
	SummaryPeriod string `xml:"SummaryPeriod" yaml:"SummaryPeriod"`
}
