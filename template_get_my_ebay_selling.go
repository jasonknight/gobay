package gobay

import "encoding/xml"

type ItemListCustomization struct {
	Include      string
	IncludeNotes string
	ListingType  string
	Sort         string
	Pagination   []struct {
		EntriesPerPage string
		PageNumber     string
	}
	DurationInDays    string
	OrderStatusFilter string
}
type GetMyeBaySellingStruct struct {
	HideVariations        string
	DetailLevel           []string
	ErrorLanguage         string
	MessageID             string
	OutputSelector        []string
	Version               string
	WarningLevel          string
	ActiveList            ItemListCustomization
	BidList               ItemListCustomization
	DeletedFromSoldList   ItemListCustomization
	DeletedFromUnsoldList ItemListCustomization
	ScheduledList         ItemListCustomization
	SellingSummary        ItemListCustomization
	SoldList              ItemListCustomization
	UnsoldList            ItemListCustomization
}

func NewGetMyeBaySellingStruct(data []byte) *GetMyeBaySellingStruct {
	var o GetMyeBaySellingStruct
	xml.Unmarshal(data, &o)
	return &o
}
