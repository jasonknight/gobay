package gobay

import (
	//"fmt"
	"net/http"
    "gopkg.in/yaml.v2"
)
type PaymentMethod string
type ShipToLocation string
type PictureDetail struct {
    GalleryDuration string
    GalleryType string
    GalleryURL string
    PhotoDisplay string
    PictureURL string
}

type ShippingServiceOption struct {
	Service  string
	Cost     string
	Priority string
}
type Product struct {
	SKU         string
	Title       string
	Price       string
    Quantity    string
	ListingType string

	ShipToLocations []ShipToLocation
	PictureDetails  []PictureDetail
	PaymentMethods  []PaymentMethod
}
type EbayCall struct {
	DevID              string
	AppID              string
	CertID             string
	CompatLevel        string
	SiteID             string
	EndPoint           string
	EbayAuthToken      string
	Country            string
	Currency           string
	PaypalEmailAddress string
	Headers            map[string]string
	Products           []Product
	TheRequest         http.Request
}

func NewEbayCall(conf []byte) *EbayCall {
	var e EbayCall
	m := make(map[string]string)
    ApplyConfiguration(conf,&e);
	m["X-EBAY-API-COMPATIBILITY-LEVEL"] = fmt.Sprintf("%s", e.CompatLevel)
	m["X-EBAY-API-DEV-NAME"] = fmt.Sprintf("%s", e.DevID)
	m["X-EBAY-API-APP-NAME"] = fmt.Sprintf("%s", e.AppID)
	m["X-EBAY-API-CERT-NAME"] = fmt.Sprintf("%s", e.CertID)
	m["X-EBAY-API-CALL-NAME"] = fmt.Sprintf("%s", e.CallName)
	m["X-EBAY-API-SITEID"] = fmt.Sprintf("%s", e.SiteID)
    e.Headers = m

	return &e
}
func ApplyConfiguration(y []byte, e *EbayCall) error {
	return yaml.Unmarshal(y, e)
}
