package gobay

import "encoding/xml"

type Result struct {
	Timestamp string
	Ack       string
	Version   int64
	Build     string
	ItemID    int64
	SKU       string
	StartTime string
	EndTime   string
	Errors    []struct {
		ShortMessage        string
		LongMessage         string
		ErrorCode           int64
		SeverityCode        string
		ErrorClassification string
		ErrorParameters     []struct {
			Value string
		}
	}

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
