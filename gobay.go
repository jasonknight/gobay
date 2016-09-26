package gobay

import (
	"bytes"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	Language           string
	MessageID          string
	WarningLevel       string
	PayPalEmailAddress string
	Callname           string
	XMLData            string
	Headers            map[string]string
	Products           []Product
	TheClient          *http.Client
}

func NewEbayCallEx(conf []byte) (*EbayCall, error) {
	var e EbayCall
	m := make(map[string]string)
	c := make(map[interface{}]interface{})
	err := LoadConfiguration(conf, &c)

	if err != nil {
		return nil, err
	}

	e.DevID = c["DevID"].(string)
	e.AppID = c["AppID"].(string)
	e.CertID = c["CertID"].(string)
	e.CompatLevel = c["CompatLevel"].(string)
	e.SiteID = c["SiteID"].(string)
	e.EndPoint = c["EndPoint"].(string)
	e.EbayAuthToken = c["EbayAuthToken"].(string)
	e.Country = c["Country"].(string)
	e.Currency = c["Currency"].(string)
	e.PayPalEmailAddress = c["PayPalEmailAddress"].(string)
	e.Language = c["Language"].(string)
	e.WarningLevel = c["WarningLevel"].(string)

	m["X-EBAY-API-COMPATIBILITY-LEVEL"] = fmt.Sprintf("%s", e.CompatLevel)
	m["X-EBAY-API-DEV-NAME"] = fmt.Sprintf("%s", e.DevID)
	m["X-EBAY-API-APP-NAME"] = fmt.Sprintf("%s", e.AppID)
	m["X-EBAY-API-CERT-NAME"] = fmt.Sprintf("%s", e.CertID)
	//m["X-EBAY-API-CALL-NAME"] = fmt.Sprintf("%s", e.CallName)
	m["X-EBAY-API-SITEID"] = fmt.Sprintf("%s", e.SiteID)
	e.Headers = m

	return &e, nil
}
func (o *EbayCall) NewProduct() *Product {
	p := NewProduct()
	p.Country = o.Country
	p.Site = o.SiteID
	p.Currency = o.Currency
	p.PayPalEmailAddress = o.PayPalEmailAddress
	return p
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
func (o *EbayCall) SetCallname(v string) {
	o.Headers["X-EBAY-API-CALL-NAME"] = v
}
func (o *EbayCall) GetCallname() string {
	return o.Headers["X-EBAY-API-CALL-NAME"]
}
func (o *EbayCall) Execute(r *[]Result) error {
	cl := o.GetCallname()
	if cl == "GeteBayOfficialTime" {
		return o.GeteBayOfficialTime(r)
	}
	return nil
}
func (o *EbayCall) GeteBayOfficialTime(r *[]Result) error {
	body, err := compileGoString("Time", GeteBayOfficialTimeTemplate(), o, nil)
	if err != nil {
		return err
	}
	final_xml := WrapCall(o.EbayAuthToken, "GeteBayOfficialTime", "", body, "")
	o.XMLData = final_xml
	err = o.Send(r)
	return err
}
func (o *EbayCall) Send(r *[]Result) error {
	o.TheClient = new(http.Client)
	fmt.Printf("About to send [[%s]]\n\n", o.XMLData)

	if o.XMLData == "" {
		err := errors.New("XMLData was empty!")
		e := NewFakeResult(fmt.Sprintf("%s", err))
		*r = append(*r, *e)
		return err
	}
	resp, err := o.TheClient.Post(o.EndPoint, "text/xml; charset=utf-8", bytes.NewBufferString(o.XMLData))
	if err != nil {
		e := NewFakeResult(fmt.Sprintf("%s", err))
		*r = append(*r, *e)
		return err
	}
	fmt.Printf("%+v\n", resp)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e := NewFakeResult(fmt.Sprintf("%s", err))
		*r = append(*r, *e)
		return err
	}
	res := NewResult(b)
	*r = append(*r, *res)
	return nil
}

func WrapCall(authToken string, name string, pre string, text string, post string) string {
	s := `<?xml version="1.0" encoding="utf-8"?>
<%sRequest xmlns="urn:ebay:apis:eBLBaseComponents">
 <RequesterCredentials> 
    <eBayAuthToken>%s</eBayAuthToken> 
  </RequesterCredentials> 
  <WarningLevel>High</WarningLevel> 
%s
%s
%s
</%sRequest>
`
	return fmt.Sprintf(s, name, authToken, pre, text, post, name)
}
