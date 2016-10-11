package gobay

import (
	"encoding/xml"
	"gopkg.in/yaml.v2"
)

type EndItemsRequest struct {
	EndItemRequestContainer []EndItemRequestContainer `xml:"EndItemRequestContainer" yaml:"EndItemRequestContainer"`
}

func NewEndItemsRequest() *EndItemsRequest {
	return &EndItemsRequest{}
}
func (o *EndItemsRequest) FromXML(data []byte) error {
	return xml.Unmarshal(data, o)
}
func (o *EndItemsRequest) FromYAML(data []byte) error {

	return yaml.Unmarshal(data, o)
}

type EndItemRequestContainer struct {
	EndingReason      string `xml:"EndingReason" yaml:"EndingReason"`
	ItemID            string `xml:"ItemID" yaml:"ItemID"`
	MessageID         string `xml:"MessageID" yaml:"MessageID"`
	SellerInventoryID string `xml:"SellerInventoryID" yaml:"SellerInventoryID"`
}
