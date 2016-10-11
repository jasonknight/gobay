package gobay

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v2"
)

type EndItemsResult struct {
	Timestamp              string
	Ack                    string
	Version                string
	Build                  string
	Errors                 []ErrorMessage
	EndItemResultContainer []EndItemResultContainer `xml:"EndItemResponseContainer" yaml:"EndItemResponseContainer"`
}

func NewEndItemsResult() *EndItemsResult {
	return &EndItemsResult{}
}
func (o *EndItemsResult) FromXML(data []byte) error {
	return xml.Unmarshal(data, o)
}
func (o *EndItemsResult) FromYAML(data []byte) error {

	return yaml.Unmarshal(data, o)
}

type GenericEndItemsResults struct {
	Results []EndItemsResult
}

func (r *GenericEndItemsResults) AddXML(b []byte) error {
	var nr EndItemsResult
	err := nr.FromXML(b)
	if err != nil {
		return err
	}
	r.Results = append(r.Results, nr)
	return nil
}
func (r *GenericEndItemsResults) AddString(b string) {
	nr := NewFakeEndItemsResult(fmt.Sprintf("%s", b))
	r.Results = append(r.Results, *nr)
}
func NewFakeEndItemsResult(msg string) *EndItemsResult {
	var o EndItemsResult
	o.Ack = "InternalFailure"
	o.Errors = append(o.Errors, ErrorMessage{ShortMessage: msg})
	return &o
}

type EndItemResultContainer struct {
	Errors        []ErrorMessage
	EndTime       string `xml:"EndTime" yaml:"EndTime"`
	CorrelationID int64  `xml:"CorrelationID" yaml:"CorrelationID"`
}
