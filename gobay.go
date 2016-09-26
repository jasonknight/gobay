package gobay

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
)

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
}
type Product struct {
	EbayID      string
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

func NewEbayCallEx(conf []byte) *EbayCall {
	var e EbayCall
	m := make(map[string]string)
	c := make(map[interface{}]interface{})
	LoadConfiguration(conf, &c)

	e.DevID = c["DevID"].(string)
	e.AppID = c["AppID"].(string)
	e.CertID = c["CertID"].(string)
	e.CompatLevel = c["CompatLevel"].(string)
	e.SiteID = c["SiteID"].(string)
	e.EndPoint = c["EndPoint"].(string)
	e.EbayAuthToken = c["EbayAuthToken"].(string)
	e.Country = c["Country"].(string)
	e.Currency = c["Currency"].(string)
	e.PaypalEmailAddress = c["PaypalEmailAddress"].(string)

	m["X-EBAY-API-COMPATIBILITY-LEVEL"] = fmt.Sprintf("%s", e.CompatLevel)
	m["X-EBAY-API-DEV-NAME"] = fmt.Sprintf("%s", e.DevID)
	m["X-EBAY-API-APP-NAME"] = fmt.Sprintf("%s", e.AppID)
	m["X-EBAY-API-CERT-NAME"] = fmt.Sprintf("%s", e.CertID)
	//m["X-EBAY-API-CALL-NAME"] = fmt.Sprintf("%s", e.CallName)
	m["X-EBAY-API-SITEID"] = fmt.Sprintf("%s", e.SiteID)
	e.Headers = m

	return &e
}
func LoadConfiguration(y []byte, e *map[interface{}]interface{}) error {
	return yaml.Unmarshal(y, e)
}

func (o *EbayCall) SetHeader(k string, v string) {
    o.Headers[k] = v
}
func (o *EbayCall) GetHeader(k string) string {
    return o.Headers[k]
}
