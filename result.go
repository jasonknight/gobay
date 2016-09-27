package gobay

import "encoding/xml"

type ErrorParameter struct {
	Value   string
	ParamID string `xml:"ParamID,attr"`
}
type Category struct {
	AutoPayEnabled    string
	B2BVATEnabled     string
	BestOfferEnabled  string
	CategoryID        string
	CategoryLevel     string
	CategoryName      string
	CategoryParentIDs []string `xml:"CategoryParentID"`

	Expired      string
	LeafCategory string
	LSD          string
	ORPA         string
	ORRA         string
	Virtual      string
}
type ErrorMessage struct {
	ShortMessage        string
	LongMessage         string
	ErrorCode           int64
	SeverityCode        string
	ErrorClassification string
	UserDisplayHint     string
	ErrorParameters     []ErrorParameter
}
type DuplicateInvocationDetail struct {
	DuplicateInvocationID string
	InvocationTrackingID  string
	Status                string
}
type Fee struct {
	Name       string `xml:"Name"`
	Amount     string `xml:"Fee"`
	CurrencyID string `xml:"currencyID,attr"`
}
type Metadata struct {
	Name  []string
	Value []string
}
type ListingRecommendation struct {
	Code      string
	FieldName string
	Group     string
	Message   string
	Metas     []Metadata `xml:"Metadata"`
	Type      string
	Values    []string `xml:"Value"`
}
type Result struct {
	Timestamp             string
	Ack                   string
	Version               string
	Build                 string
	ItemID                string
	SKU                   string
	StartTime             string
	EndTime               string
	Category2ID           string
	CategoryID            string
	CorrelationID         string
	HardExpirationWarning string
	Message               string
	CategoryCount         string
	CategoryVersion       string
	MinimumReservePrice   string
	ReduceReserveAllowed  string
	ReservePriceAllowed   string
	UpdateTime            string

	Items      []Result   `xml:"AddItemResponseContainer"`
	Categories []Category `xml:"CategoryArray>Category"`

	DiscountReasons            []string `xml:"DiscountReason"`
	Errors                     []ErrorMessage
	DuplicateInvocationDetails []DuplicateInvocationDetail
	Fees                       []Fee `xml:"Fees>Fee"`
	ListingRecommendations     []ListingRecommendation
}

func NewResult(data []byte) (*Result, error) {
	var o Result
	err := xml.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}
	return &o, nil
}
func NewFakeResult(msg string) *Result {
	var o Result
	o.Ack = "InternalFailure"
	o.Errors = append(o.Errors, ErrorMessage{ShortMessage: msg})
	return &o
}
