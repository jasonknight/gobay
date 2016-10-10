package gobay

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v2"
)

type NotificationPreferencesResult struct {
	Timestamp                      string                        `xml:"Timestamp" yaml:"Timestamp"`
	Ack                            string                        `xml:"Ack" yaml:"Ack"`
	Version                        int64                         `xml:"Version" yaml:"Version"`
	Build                          string                        `xml:"Build" yaml:"Build"`
	ApplicationDeliveryPreferences ApplicationDeliveryPreference `xml:"ApplicationDeliveryPreferences" yaml:"ApplicationDeliveryPreferences"`
	UserDeliveryPreferenceArray    UserDeliveryPreference        `xml:"UserDeliveryPreferenceArray" yaml:"UserDeliveryPreferenceArray"`
	Errors                         []ErrorMessage
}

func NewNotificationPreferencesResult() *NotificationPreferencesResult {
	return &NotificationPreferencesResult{}
}

func (o *NotificationPreferencesResult) FromXML(data []byte) error {
	return xml.Unmarshal(data, o)
}

func (o *NotificationPreferencesResult) FromYAML(data []byte) error {
	return yaml.Unmarshal(data, o)
}

func (o *NotificationPreferencesResult) Success() bool {
	if o.Ack == "Success" {
		return true
	}
	return false
}

func (o *NotificationPreferencesResult) Failure() bool {
	if o.Ack == "Failure" {
		return true
	}
	return false
}

func (o *NotificationPreferencesResult) Warning() bool {
	if o.Ack == "Warning" {
		return true
	}
	return false
}

type ApplicationDeliveryPreference struct {
	ApplicationURL          string              `xml:"ApplicationURL" yaml:"ApplicationURL"`
	ApplicationEnable       string              `xml:"ApplicationEnable" yaml:"ApplicationEnable"`
	AlertEmail              string              `xml:"AlertEmail" yaml:"AlertEmail"`
	AlertEnable             string              `xml:"AlertEnable" yaml:"AlertEnable"`
	NotificationPayloadType string              `xml:"NotificationPayloadType" yaml:"NotificationPayloadType"`
	DeviceType              string              `xml:"DeviceType" yaml:"DeviceType"`
	PayloadEncodingType     string              `xml:"PayloadEncodingType" yaml:"PayloadEncodingType"`
	PayloadVersion          int64               `xml:"PayloadVersion" yaml:"PayloadVersion"`
	DeliveryURLDetails      []DeliveryURLDetail `xml:"DeliveryURLDetails" yaml:"DeliveryURLDetails"`
}

type UserDeliveryPreference struct {
	EnabledNotifications []NotificationEnable `xml:"NotificationEnable" yaml:"NotificationEnable"`
}

type NotificationEnable struct {
	EventType   string `xml:"EventType" yaml:"EventType"`
	EventEnable string `xml:"EventEnable" yaml:"EventEnable"`
}

type GenericNotificationPreferenceResults struct {
	Results []Result
}

func (r *GenericNotificationPreferenceResults) AddXML(b []byte) error {
	var nr Result
	err := nr.FromXML(b)
	if err != nil {
		return err
	}
	r.Results = append(r.Results, nr)
	return nil
}
func (r *GenericNotificationPreferenceResults) AddString(b string) {
	nr := NewFakeResult(fmt.Sprintf("%s", b))
	r.Results = append(r.Results, *nr)
}
