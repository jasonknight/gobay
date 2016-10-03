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
type Template_get_my_ebay_selling struct {
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

func NewTemplate_get_my_ebay_selling(data []byte) *Template_get_my_ebay_selling {
	var o Template_get_my_ebay_selling
	xml.Unmarshal(data, &o)
	return &o
}
