package gobay

import (
	// "bytes"
	// "errors"
	"fmt"
	// "io/ioutil"
	// "net/http"
	"gopkg.in/yaml.v2"
)

type ShipToLocationFilter func(o ShipToLocation) bool
type PictureDetailFilter func(o PictureDetail) bool
type PaymentMethodFilter func(o PaymentMethod) bool
type PaymentMethod string
type ShipToLocation string

type PictureDetail struct {
	GalleryDuration string
	GalleryType     string
	GalleryURL      string
	PhotoDisplay    string
	PictureURL      string
}

type ShippingServiceOption struct {
	Service  string
	Cost     string
	Priority string
	Location ShipToLocation
}
type Item struct {
	EbayID             int64
	SKU                string
	Title              string
	Description        string
	StartPrice         float32
	Quantity           string
	ListingType        string
	Country            string
	Currency           string
	Location           string
	PayPalEmailAddress string
	PrimaryCategory    string
	Site               string
	StoreCategoryID    int64
	DispatchTimeMax    string
	ListingDuration    string

	ShipToLocations                     []ShipToLocation
	PictureDetails                      []PictureDetail
	PaymentMethods                      []PaymentMethod
	ShippingServiceOptions              []ShippingServiceOption
	InternationalShippingServiceOptions []ShippingServiceOption
}

func NewItem() *Item {
	return &Item{}
}

func (o *Item) Clone() *Item {
	var no Item
	no.EbayID = o.EbayID
	no.SKU = o.SKU
	no.Title = o.Title
	no.Description = o.Description
	no.StartPrice = o.StartPrice
	no.Quantity = o.Quantity
	no.ListingType = o.ListingType
	no.Country = o.Country
	no.Currency = o.Currency
	no.Location = o.Location
	no.PayPalEmailAddress = o.PayPalEmailAddress
	no.PrimaryCategory = o.PrimaryCategory
	no.Site = o.Site
	no.StoreCategoryID = o.StoreCategoryID
	no.DispatchTimeMax = o.DispatchTimeMax
	no.ListingDuration = o.ListingDuration

	no.ShipToLocations = o.ShipToLocations
	no.PictureDetails = o.PictureDetails
	no.PaymentMethods = o.PaymentMethods
	no.ShippingServiceOptions = o.ShippingServiceOptions
	no.InternationalShippingServiceOptions = o.InternationalShippingServiceOptions
	return &no
}

func (o *Item) SetSKU(v string) {
	o.SKU = v
}

func (o *Item) GetSKU() string {
	return o.SKU
}

func (o *Item) SetTitle(v string) {
	o.Title = v
}

func (o *Item) GetTitle() string {
	return o.Title
}

func (o *Item) SetPrice(v float32) {
	o.StartPrice = v
}

func (o *Item) GetPrice() float32 {
	return o.StartPrice
}

func (o *Item) SetQuantity(v string) {
	o.Quantity = v
}

func (o *Item) GetQuantity() string {
	return o.Quantity
}

func (o *Item) SetListingType(v string) {
	o.ListingType = v
}

func (o *Item) GetListingType() string {
	return o.ListingType
}

func (o *Item) FilterShipToLocations(f ShipToLocationFilter) []ShipToLocation {
	tmp := o.ShipToLocations[:0]
	for _, x := range o.ShipToLocations {
		if f(x) {
			tmp = append(tmp, x)
		}
	}
	return tmp
}

func (o *Item) AddShipToLocation(v ShipToLocation) {
	o.ShipToLocations = append(o.ShipToLocations, v)
}

func (o *Item) RemoveShipToLocation(i int) {
	if i > len(o.ShipToLocations) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Item", "ShipToLocations", len(o.ShipToLocations)))
	}
	o.ShipToLocations = o.ShipToLocations[:i+copy(o.ShipToLocations[i:], o.ShipToLocations[i+1:])]
}

func (o *Item) GetShipToLocation(i int) ShipToLocation {
	if i > len(o.ShipToLocations) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Item", "ShipToLocations", len(o.ShipToLocations)))
	}
	return o.ShipToLocations[i]
}

func (o *Item) SetShipToLocations(v []ShipToLocation) {
	o.ShipToLocations = v
}

func (o *Item) GetShipToLocations() []ShipToLocation {
	return o.ShipToLocations
}

func (o *Item) FilterPictureDetails(f PictureDetailFilter) []PictureDetail {
	tmp := o.PictureDetails[:0]
	for _, x := range o.PictureDetails {
		if f(x) {
			tmp = append(tmp, x)
		}
	}
	return tmp
}

func (o *Item) AddPictureDetail(v PictureDetail) {
	o.PictureDetails = append(o.PictureDetails, v)
}

func (o *Item) RemovePictureDetail(i int) {
	if i > len(o.PictureDetails) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Item", "PictureDetails", len(o.PictureDetails)))
	}
	o.PictureDetails = o.PictureDetails[:i+copy(o.PictureDetails[i:], o.PictureDetails[i+1:])]
}

func (o *Item) GetPictureDetail(i int) PictureDetail {
	if i > len(o.PictureDetails) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Item", "PictureDetails", len(o.PictureDetails)))
	}
	return o.PictureDetails[i]
}

func (o *Item) SetPictureDetails(v []PictureDetail) {
	o.PictureDetails = v
}

func (o *Item) GetPictureDetails() []PictureDetail {
	return o.PictureDetails
}

func (o *Item) FilterPaymentMethods(f PaymentMethodFilter) []PaymentMethod {
	tmp := o.PaymentMethods[:0]
	for _, x := range o.PaymentMethods {
		if f(x) {
			tmp = append(tmp, x)
		}
	}
	return tmp
}

func (o *Item) AddPaymentMethod(v PaymentMethod) {
	o.PaymentMethods = append(o.PaymentMethods, v)
}

func (o *Item) RemovePaymentMethod(i int) {
	if i > len(o.PaymentMethods) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Item", "PaymentMethods", len(o.PaymentMethods)))
	}
	o.PaymentMethods = o.PaymentMethods[:i+copy(o.PaymentMethods[i:], o.PaymentMethods[i+1:])]
}

func (o *Item) GetPaymentMethod(i int) PaymentMethod {
	if i > len(o.PaymentMethods) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Item", "PaymentMethods", len(o.PaymentMethods)))
	}
	return o.PaymentMethods[i]
}

func (o *Item) SetPaymentMethods(v []PaymentMethod) {
	o.PaymentMethods = v
}

func (o *Item) GetPaymentMethods() []PaymentMethod {
	return o.PaymentMethods
}
func (o *Item) FromYAML(data []byte) error {
	return yaml.Unmarshal(data, o)
}

// Debug Functions

func (o_PictureDetail *PictureDetail) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sPictureDetail.GalleryDuration: %s\n",txt, o_PictureDetail.GalleryDuration)
	txt = fmt.Sprintf("%sPictureDetail.GalleryType: %s\n",txt, o_PictureDetail.GalleryType)
	txt = fmt.Sprintf("%sPictureDetail.GalleryURL: %s\n",txt, o_PictureDetail.GalleryURL)
	txt = fmt.Sprintf("%sPictureDetail.PhotoDisplay: %s\n",txt, o_PictureDetail.PhotoDisplay)
	txt = fmt.Sprintf("%sPictureDetail.PictureURL: %s\n",txt, o_PictureDetail.PictureURL)
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

func (o_ShippingServiceOption *ShippingServiceOption) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sShippingServiceOption.Service: %s\n",txt, o_ShippingServiceOption.Service)
	txt = fmt.Sprintf("%sShippingServiceOption.Cost: %s\n",txt, o_ShippingServiceOption.Cost)
	txt = fmt.Sprintf("%sShippingServiceOption.Priority: %s\n",txt, o_ShippingServiceOption.Priority)
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

func (o_Item *Item) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sItem.SKU: %s\n",txt, o_Item.SKU)
	txt = fmt.Sprintf("%sItem.Title: %s\n",txt, o_Item.Title)
	txt = fmt.Sprintf("%sItem.Description: %s\n",txt, o_Item.Description)
	txt = fmt.Sprintf("%sItem.Quantity: %s\n",txt, o_Item.Quantity)
	txt = fmt.Sprintf("%sItem.ListingType: %s\n",txt, o_Item.ListingType)
	txt = fmt.Sprintf("%sItem.Country: %s\n",txt, o_Item.Country)
	txt = fmt.Sprintf("%sItem.Currency: %s\n",txt, o_Item.Currency)
	txt = fmt.Sprintf("%sItem.Location: %s\n",txt, o_Item.Location)
	txt = fmt.Sprintf("%sItem.PayPalEmailAddress: %s\n",txt, o_Item.PayPalEmailAddress)
	txt = fmt.Sprintf("%sItem.PrimaryCategory: %s\n",txt, o_Item.PrimaryCategory)
	txt = fmt.Sprintf("%sItem.Site: %s\n",txt, o_Item.Site)
	txt = fmt.Sprintf("%sItem.DispatchTimeMax: %s\n",txt, o_Item.DispatchTimeMax)
	txt = fmt.Sprintf("%sItem.ListingDuration: %s\n",txt, o_Item.ListingDuration)
	for _,v := range o_Item.ShipToLocations {
		txt = fmt.Sprintf("%s%s\n",txt, v)
	}
	for _,v := range o_Item.PictureDetails {
		txt = fmt.Sprintf("%s%s\n",txt, v.Debug())
	}
	for _,v := range o_Item.PaymentMethods {
		txt = fmt.Sprintf("%s%s\n",txt, v)
	}
	for _,v := range o_Item.ShippingServiceOptions {
		txt = fmt.Sprintf("%s%s\n",txt, v.Debug())
	}
	for _,v := range o_Item.InternationalShippingServiceOptions {
		txt = fmt.Sprintf("%s%s\n",txt, v.Debug())
	}
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

