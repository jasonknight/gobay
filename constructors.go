package gobay

import "encoding/xml"

import "fmt"
import "errors"

func NewNotificationResultEx(data []byte) (*NotificationResult, error) {
	var o NotificationResult
	err := xml.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}
	return &o, nil
}
func NewResultEx(r interface{}, data []byte) (interface{}, error) {
	switch r.(type) {
	case *[]Result:
		var o Result
		err := xml.Unmarshal(data, &o)
		if err != nil {
			return nil, err
		}
		return o, nil
		break
	case *[]MyeBaySellingResult:
		var o MyeBaySellingResult
		err := xml.Unmarshal(data, &o)
		if err != nil {
			return nil, err
		}
		return o, nil
		break
	case *[]NotificationPreferencesResult:
		var o NotificationPreferencesResult
		err := o.FromXML(data)
		if err != nil {
			return nil, err
		}
		return o, nil
		break
	default:
		panic("AddToResult type is not supported, this is bad, DO NOT DO THIS")
	}

	return nil, errors.New("could not detect result type")
}
func NewFakeResult(msg string) *Result {
	var o Result
	o.Ack = "InternalFailure"
	o.Errors = append(o.Errors, ErrorMessage{ShortMessage: msg})
	return &o
}
func NewEbayCall() *EbayCall {
	return &EbayCall{}
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
	e.Cache = c["Cache"].(string)

	e.AddItemsLimit = 5
	e.CallDepthLimit = 3

	m["X-EBAY-API-COMPATIBILITY-LEVEL"] = fmt.Sprintf("%s", e.CompatLevel)
	m["X-EBAY-API-DEV-NAME"] = fmt.Sprintf("%s", e.DevID)
	m["X-EBAY-API-APP-NAME"] = fmt.Sprintf("%s", e.AppID)
	m["X-EBAY-API-CERT-NAME"] = fmt.Sprintf("%s", e.CertID)
	//m["X-EBAY-API-CALL-NAME"] = fmt.Sprintf("%s", e.Callname)
	m["X-EBAY-API-SITEID"] = fmt.Sprintf("%s", e.SiteID)
	e.Headers = m

	if e.Cache == "" {
		return nil, errors.New("You MUST specify a cache location in the yml file!")
	}

	return &e, nil
}

func (o *EbayCall) NewItem() *Item {
	p := NewItem()
	p.Country = o.Country
	p.Site = SiteIDToCode(o.SiteID)
	p.Currency = o.Currency
	p.PayPalEmailAddress = o.PayPalEmailAddress
	return p
}
func NewPictureDetail() *PictureDetail {
	return &PictureDetail{}
}
func NewShippingServiceOption() *ShippingServiceOption {
	return &ShippingServiceOption{}
}
func NewItem() *Item {
	return &Item{}
}
func NewGetCategoriesStruct() *GetCategoriesStruct {
	var s GetCategoriesStruct
	s.LevelLimit = "3"
	s.ViewAllNodes = "false"
	s.DetailLevels = append(s.DetailLevels, "ReturnAll")
	outputs := [...]string{
		"CategoryID",
		"CategoryLevel",
		"CategoryName",
		"CategoryParentID",
	}
	for _, v := range outputs {
		s.OutputSelectors = append(s.OutputSelectors, v)
	}
	return &s
}
func NewMyeBaySelling() *MyeBaySelling {
	o := MyeBaySelling{
		// ActiveList: nil,
		// SoldList: nil,
		// UnsoldList: nil,
		// BidList: nil,
		// DeletedFromSoldList: nil,
		// DeletedFromUnsoldList:nil,
		// ScheduledList: nil,
		SellingSummary: struct {
			Include bool
		}{Include: false},
		HideVariations: false,
	}
	return &o
}
func NewGetEbayDetailsStruct() *EbayDetails {
	var s EbayDetails
	var dnames = []string{
		"CountryDetails",
		"CurrencyDetails",
		"ProductDetails",
		"ReturnPolicyDetails",
		"SiteDetails",
	}

	for _, dn := range dnames {
		s.DetailNames = append(s.DetailNames, dn)
	}
	return &s
}
