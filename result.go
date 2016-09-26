package gobay

import "encoding/xml"

type ErrorParameter struct {
	Value string
}
type ErrorMessage struct {
	ShortMessage        string
	LongMessage         string
	ErrorCode           int64
	SeverityCode        string
	ErrorClassification string
	ErrorParameters     []ErrorParameter
}
type Result struct {
	Timestamp string
	Ack       string
	Version   int64
	Build     string
	ItemID    int64
	SKU       string
	StartTime string
	EndTime   string
	Errors    []ErrorMessage

	Fees []struct {
		Name   string
		Amount float64 `xml:Fee`
	}
}

func NewResult(data []byte) *Result {
	var o Result
	xml.Unmarshal(data, &o)
	return &o
}
func NewFakeResult(msg string) *Result {
	var o Result
	o.Ack = "InternalFailure"
	o.Errors = append(o.Errors, ErrorMessage{ShortMessage: msg})
	return &o
}
