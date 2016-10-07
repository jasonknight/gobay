package gobay
import (
	"encoding/xml"
	"gopkg.in/yaml.v2"
)
type NotificationPreferencesRequest struct{
	PreferenceLevel string `xml:"PreferenceLevel" yaml:"PreferenceLevel"`
}

func NewNotificationPreferencesRequest() *NotificationPreferencesRequest {
	return &NotificationPreferencesRequest{}
}

func (o *NotificationPreferencesRequest) FromXML(data []byte) error {
	 return xml.Unmarshal(data,o)
}

func (o *NotificationPreferencesRequest) FromYAML(data []byte) error {
	 return yaml.Unmarshal(data,o)
}

