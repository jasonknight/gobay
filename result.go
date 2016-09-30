package gobay

import "encoding/xml"
import "fmt"

//import "bytes"
//import "io"

//import "reflect"

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
type FeedbackComment struct {
	CommentingUser      string
	CommentingUserScore string
	CommentText         string
	CommentTime         string
	CommentType         string
	ItemID              string
	Role                string
	FeedbackID          string
	TransactionID       string
	ItemTitle           string
	ItemPrice           string
}
type NotificationFeedback struct {
	Comments []FeedbackComment `xml:"FeedbackDetail"`
}

type NotificationResult struct {
	Body struct {
		InnerXML         string                 `xml:",innerxml"`
		Feedback         []NotificationFeedback `xml:"GetFeedbackResponse>FeedbackDetailArray"`
		PaginationResult struct {
			TotalNumberOfPages   string
			TotalNumberOfEntries string
		} `xml:"GetFeedbackResponse>PaginationResult"`
		EntriesPerPage        string `xml:"GetFeedbackResponse>EntriesPerPage"`
		PageNumber            string `xml:"GetFeedbackResponse>PageNumber"`
		NotificationEventName string `xml:"GetFeedbackResponse>NotificationEventName"`
		RecipientUserID       string `xml:"GetFeedbackResponse>RecipientUserID"`
	}

	Header struct {
		InnerXML string `xml:",innerxml"`
	}
	// Header string `xml:"Header,innerxml"`
}

func NewNotificationResult(data []byte) (*NotificationResult, error) {
	var o NotificationResult
	err := xml.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}
	return &o, nil
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

func (r *Result) Success() bool {
	s := []string{"Success", "Warning"}
	return InStringSlice(s, r.Ack)
}
func (r *Result) Warning() bool {
	if r.Ack == "Warning" {
		return true
	}
	return false
}
func (r *Result) Failure() bool {
	s := []string{"Failure", "PartialFailure"}
	return InStringSlice(s, r.Ack)
}

// Debug Functions

func (o_ErrorParameter *ErrorParameter) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sErrorParameter.Value: %s\n", txt, o_ErrorParameter.Value)
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

func (o_Category *Category) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sCategory.AutoPayEnabled: %s\n", txt, o_Category.AutoPayEnabled)
	txt = fmt.Sprintf("%sCategory.B2BVATEnabled: %s\n", txt, o_Category.B2BVATEnabled)
	txt = fmt.Sprintf("%sCategory.BestOfferEnabled: %s\n", txt, o_Category.BestOfferEnabled)
	txt = fmt.Sprintf("%sCategory.CategoryID: %s\n", txt, o_Category.CategoryID)
	txt = fmt.Sprintf("%sCategory.CategoryLevel: %s\n", txt, o_Category.CategoryLevel)
	txt = fmt.Sprintf("%sCategory.CategoryName: %s\n", txt, o_Category.CategoryName)
	txt = fmt.Sprintf("%sCategory.Expired: %s\n", txt, o_Category.Expired)
	txt = fmt.Sprintf("%sCategory.LeafCategory: %s\n", txt, o_Category.LeafCategory)
	txt = fmt.Sprintf("%sCategory.LSD: %s\n", txt, o_Category.LSD)
	txt = fmt.Sprintf("%sCategory.ORPA: %s\n", txt, o_Category.ORPA)
	txt = fmt.Sprintf("%sCategory.ORRA: %s\n", txt, o_Category.ORRA)
	txt = fmt.Sprintf("%sCategory.Virtual: %s\n", txt, o_Category.Virtual)
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

func (o_ErrorMessage *ErrorMessage) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sErrorMessage.ShortMessage: %s\n", txt, o_ErrorMessage.ShortMessage)
	txt = fmt.Sprintf("%sErrorMessage.LongMessage: %s\n", txt, o_ErrorMessage.LongMessage)
	txt = fmt.Sprintf("%sErrorMessage.SeverityCode: %s\n", txt, o_ErrorMessage.SeverityCode)
	txt = fmt.Sprintf("%sErrorMessage.ErrorClassification: %s\n", txt, o_ErrorMessage.ErrorClassification)
	txt = fmt.Sprintf("%sErrorMessage.UserDisplayHint: %s\n", txt, o_ErrorMessage.UserDisplayHint)
	for _, v := range o_ErrorMessage.ErrorParameters {
		txt = fmt.Sprintf("%s%s\n", txt, v.Debug())
	}
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

func (o_DuplicateInvocationDetail *DuplicateInvocationDetail) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sDuplicateInvocationDetail.DuplicateInvocationID: %s\n", txt, o_DuplicateInvocationDetail.DuplicateInvocationID)
	txt = fmt.Sprintf("%sDuplicateInvocationDetail.InvocationTrackingID: %s\n", txt, o_DuplicateInvocationDetail.InvocationTrackingID)
	txt = fmt.Sprintf("%sDuplicateInvocationDetail.Status: %s\n", txt, o_DuplicateInvocationDetail.Status)
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

func (o_Fee *Fee) Debug() string {
	var txt string
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

func (o_Metadata *Metadata) Debug() string {
	var txt string
	for _, v := range o_Metadata.Name {
		txt = fmt.Sprintf("%s%s\n", txt, v)
	}
	for _, v := range o_Metadata.Value {
		txt = fmt.Sprintf("%s%s\n", txt, v)
	}
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

func (o_ListingRecommendation *ListingRecommendation) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sListingRecommendation.Code: %s\n", txt, o_ListingRecommendation.Code)
	txt = fmt.Sprintf("%sListingRecommendation.FieldName: %s\n", txt, o_ListingRecommendation.FieldName)
	txt = fmt.Sprintf("%sListingRecommendation.Group: %s\n", txt, o_ListingRecommendation.Group)
	txt = fmt.Sprintf("%sListingRecommendation.Message: %s\n", txt, o_ListingRecommendation.Message)
	txt = fmt.Sprintf("%sListingRecommendation.Type: %s\n", txt, o_ListingRecommendation.Type)
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

func (o_Result *Result) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sResult.Timestamp: %s\n", txt, o_Result.Timestamp)
	txt = fmt.Sprintf("%sResult.Ack: %s\n", txt, o_Result.Ack)
	txt = fmt.Sprintf("%sResult.Version: %s\n", txt, o_Result.Version)
	txt = fmt.Sprintf("%sResult.Build: %s\n", txt, o_Result.Build)
	txt = fmt.Sprintf("%sResult.ItemID: %s\n", txt, o_Result.ItemID)
	txt = fmt.Sprintf("%sResult.SKU: %s\n", txt, o_Result.SKU)
	txt = fmt.Sprintf("%sResult.StartTime: %s\n", txt, o_Result.StartTime)
	txt = fmt.Sprintf("%sResult.EndTime: %s\n", txt, o_Result.EndTime)
	txt = fmt.Sprintf("%sResult.Category2ID: %s\n", txt, o_Result.Category2ID)
	txt = fmt.Sprintf("%sResult.CategoryID: %s\n", txt, o_Result.CategoryID)
	txt = fmt.Sprintf("%sResult.CorrelationID: %s\n", txt, o_Result.CorrelationID)
	txt = fmt.Sprintf("%sResult.HardExpirationWarning: %s\n", txt, o_Result.HardExpirationWarning)
	txt = fmt.Sprintf("%sResult.Message: %s\n", txt, o_Result.Message)
	txt = fmt.Sprintf("%sResult.CategoryCount: %s\n", txt, o_Result.CategoryCount)
	txt = fmt.Sprintf("%sResult.CategoryVersion: %s\n", txt, o_Result.CategoryVersion)
	txt = fmt.Sprintf("%sResult.MinimumReservePrice: %s\n", txt, o_Result.MinimumReservePrice)
	txt = fmt.Sprintf("%sResult.ReduceReserveAllowed: %s\n", txt, o_Result.ReduceReserveAllowed)
	txt = fmt.Sprintf("%sResult.ReservePriceAllowed: %s\n", txt, o_Result.ReservePriceAllowed)
	txt = fmt.Sprintf("%sResult.UpdateTime: %s\n", txt, o_Result.UpdateTime)
	for _, v := range o_Result.Errors {
		txt = fmt.Sprintf("%s%s\n", txt, v.Debug())
	}
	for _, v := range o_Result.DuplicateInvocationDetails {
		txt = fmt.Sprintf("%s%s\n", txt, v.Debug())
	}
	for _, v := range o_Result.ListingRecommendations {
		txt = fmt.Sprintf("%s%s\n", txt, v.Debug())
	}
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}
