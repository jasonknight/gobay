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
type Product struct {
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

func NewProduct() *Product {
	return &Product{}
}

func (o *Product) Clone() *Product {
	var no Product
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

func (o *Product) SetSKU(v string) {
	o.SKU = v
}

func (o *Product) GetSKU() string {
	return o.SKU
}

func (o *Product) SetTitle(v string) {
	o.Title = v
}

func (o *Product) GetTitle() string {
	return o.Title
}

func (o *Product) SetPrice(v float32) {
	o.StartPrice = v
}

func (o *Product) GetPrice() float32 {
	return o.StartPrice
}

func (o *Product) SetQuantity(v string) {
	o.Quantity = v
}

func (o *Product) GetQuantity() string {
	return o.Quantity
}

func (o *Product) SetListingType(v string) {
	o.ListingType = v
}

func (o *Product) GetListingType() string {
	return o.ListingType
}

func (o *Product) FilterShipToLocations(f ShipToLocationFilter) []ShipToLocation {
	tmp := o.ShipToLocations[:0]
	for _, x := range o.ShipToLocations {
		if f(x) {
			tmp = append(tmp, x)
		}
	}
	return tmp
}

func (o *Product) AddShipToLocation(v ShipToLocation) {
	o.ShipToLocations = append(o.ShipToLocations, v)
}

func (o *Product) RemoveShipToLocation(i int) {
	if i > len(o.ShipToLocations) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Product", "ShipToLocations", len(o.ShipToLocations)))
	}
	o.ShipToLocations = o.ShipToLocations[:i+copy(o.ShipToLocations[i:], o.ShipToLocations[i+1:])]
}

func (o *Product) GetShipToLocation(i int) ShipToLocation {
	if i > len(o.ShipToLocations) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Product", "ShipToLocations", len(o.ShipToLocations)))
	}
	return o.ShipToLocations[i]
}

func (o *Product) SetShipToLocations(v []ShipToLocation) {
	o.ShipToLocations = v
}

func (o *Product) GetShipToLocations() []ShipToLocation {
	return o.ShipToLocations
}

func (o *Product) FilterPictureDetails(f PictureDetailFilter) []PictureDetail {
	tmp := o.PictureDetails[:0]
	for _, x := range o.PictureDetails {
		if f(x) {
			tmp = append(tmp, x)
		}
	}
	return tmp
}

func (o *Product) AddPictureDetail(v PictureDetail) {
	o.PictureDetails = append(o.PictureDetails, v)
}

func (o *Product) RemovePictureDetail(i int) {
	if i > len(o.PictureDetails) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Product", "PictureDetails", len(o.PictureDetails)))
	}
	o.PictureDetails = o.PictureDetails[:i+copy(o.PictureDetails[i:], o.PictureDetails[i+1:])]
}

func (o *Product) GetPictureDetail(i int) PictureDetail {
	if i > len(o.PictureDetails) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Product", "PictureDetails", len(o.PictureDetails)))
	}
	return o.PictureDetails[i]
}

func (o *Product) SetPictureDetails(v []PictureDetail) {
	o.PictureDetails = v
}

func (o *Product) GetPictureDetails() []PictureDetail {
	return o.PictureDetails
}

func (o *Product) FilterPaymentMethods(f PaymentMethodFilter) []PaymentMethod {
	tmp := o.PaymentMethods[:0]
	for _, x := range o.PaymentMethods {
		if f(x) {
			tmp = append(tmp, x)
		}
	}
	return tmp
}

func (o *Product) AddPaymentMethod(v PaymentMethod) {
	o.PaymentMethods = append(o.PaymentMethods, v)
}

func (o *Product) RemovePaymentMethod(i int) {
	if i > len(o.PaymentMethods) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Product", "PaymentMethods", len(o.PaymentMethods)))
	}
	o.PaymentMethods = o.PaymentMethods[:i+copy(o.PaymentMethods[i:], o.PaymentMethods[i+1:])]
}

func (o *Product) GetPaymentMethod(i int) PaymentMethod {
	if i > len(o.PaymentMethods) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "Product", "PaymentMethods", len(o.PaymentMethods)))
	}
	return o.PaymentMethods[i]
}

func (o *Product) SetPaymentMethods(v []PaymentMethod) {
	o.PaymentMethods = v
}

func (o *Product) GetPaymentMethods() []PaymentMethod {
	return o.PaymentMethods
}
func (o *Product) FromYAML(data []byte) error {
	return yaml.Unmarshal(data, o)
}
