package gobay

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v2"
)

type PaginationStruct struct {
	TotalNumberOfEntries string
	TotalNumberOfPages   string
}

type SummaryStruct struct {
	ActiveAuctionCount       string
	AmountLimitRemaining     string
	AuctionBidCount          string
	AuctionSellingCount      string
	ClassifiedAdCount        string
	ClassifiedAdOfferCount   string
	QuantityLimitRemaining   string
	SoldDurationInDays       string
	TotalAuctionSellingValue string
	TotalLeadCount           string
	TotalListingsWithLeads   string
	TotalSoldCount           string
	TotalSoldValue           string
}
type BuyerStruct struct {
	Email           string
	StaticAlias     string
	UserID          string
	ShippingAddress struct {
		PostalCode string
	}
}
type TransactionStruct struct {
	ConvertedTransactionPrice string
	CreatedDate               string
	IsMultiLegShipping        string
	LineItemID                string
	PaidTime                  string
	PaisaPayID                string
	Platform                  string
	QuantityPurchased         string
	SellerPaidStatus          string
	ShippedTime               string
	ID                        string `xml:"TransactionID"`
	Price                     string
	Buyer                     BuyerStruct `xml:"Buyer>BuyerInfo"`
	OrderLineItemID           string

	FeedbackLeft string `xml:"FeedbackLeft>CommentType"`

	FeedbackReceived string `xml:"FeedbackReceived>CommentType"`

	Item ItemResultStruct `xml:"Item"`

	Status struct {
		PaymentHold string
	}
}
type OrderStruct struct {
	ID           string `xml:"OrderID"`
	Subtotal     string
	Transactions []TransactionStruct `xml:"TransactionArray>Transaction"`
}
type OrderTransactionArray struct {
	Order        OrderStruct         `xml:"OrderTransaction>Order"`
	Transactions []TransactionStruct `xml:"Transaction"`
	Pagination   PaginationStruct    `xml:"PaginationResult"`
}

type ItemArrayStruct struct {
	Items      []ItemResultStruct `xml:"Item"`
	Pagination PaginationStruct   `xml:PaginationResult`
}
type MyeBaySellingResult struct {
	Ack                             string
	Build                           string
	CorrelationID                   string
	HardExpirationWarning           string
	Timestamp                       string
	Version                         string
	Summary                         SummaryStruct
	SellingSummary                  SummaryStruct
	ActiveList                      []ItemArrayStruct       `xml:"ActiveList>ItemArray"`
	DeletedFromSoldList             []OrderTransactionArray `xml:"DeletedFromSoldList>OrderTransactionArray"`
	DeletedFromUnsoldList           []ItemArrayStruct       `xml:"DeletedFromUnsoldList>ItemArray"`
	ScheduledList                   []ItemArrayStruct       `xml:"ScheduledList>ItemArray"`
	SoldList                        []OrderTransactionArray `xml:"SoldList>OrderTransactionArray"`
	UnsoldList                      []ItemArrayStruct       `xml:"UnsoldList>ItemArray"`
	ActiveListPagination            PaginationStruct        `xml:"ActiveList>PaginationResult"`
	DeletedFromSoldListPagination   PaginationStruct        `xml:"DeletedFromSoldList>PaginationResult"`
	DeletedFromUnsoldListPagination PaginationStruct        `xml:"DeletedFromUnsoldList>PaginationResult"`
	ScheduledListPagination         PaginationStruct        `xml:"ScheduledList>PaginationResult"`
	SoldListPagination              PaginationStruct        `xml:"SoldList>PaginationResult"`
	UnsoldListPagination            PaginationStruct        `xml:"UnsoldList>PaginationResult"`
	Errors                          []ErrorMessage
}

func NewMyeBaySellingResult() *MyeBaySellingResult {
	return &MyeBaySellingResult{}
}
func (o *MyeBaySellingResult) FromXML(data []byte) error {
	return xml.Unmarshal(data, o)
}
func (o *MyeBaySellingResult) FromYAML(data []byte) error {

	return yaml.Unmarshal(data, o)
}

type GenericMyeBaySellingResults struct {
	Results []MyeBaySellingResult
}

func (r *GenericMyeBaySellingResults) AddXML(b []byte) error {
	var nr MyeBaySellingResult
	err := nr.FromXML(b)
	if err != nil {
		return err
	}
	r.Results = append(r.Results, nr)
	return nil
}
func (r *GenericMyeBaySellingResults) AddString(b string) {
	nr := NewFakeMyeBaySellingResult(fmt.Sprintf("%s", b))
	r.Results = append(r.Results, *nr)
}
func NewFakeMyeBaySellingResult(msg string) *MyeBaySellingResult {
	var o MyeBaySellingResult
	o.Ack = "InternalFailure"
	o.Errors = append(o.Errors, ErrorMessage{ShortMessage: msg})
	return &o
}
func (o *MyeBaySellingResult) Success() bool {
	if o.Ack == "Success" {
		return true
	}
	return false
}

func (o *MyeBaySellingResult) Failure() bool {
	if o.Ack == "Failure" {
		return true
	}
	return false
}

func (o *MyeBaySellingResult) Warning() bool {
	if o.Ack == "Warning" {
		return true
	}
	return false
}
