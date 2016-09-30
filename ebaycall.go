package gobay

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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
	MessageIDs         []string
	WarningLevel       string
	PayPalEmailAddress string
	Callname           string
	XMLData            string
	Cache              string
	AddItemsLimit      int
	CallDepthLimit     int
	CallDepth          int
	Headers            map[string]string
	Items              []*Item
	TheClient          *http.Client
	CategoryCallInfo   *GetCategoriesStruct
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

func (o *EbayCall) SetHeader(k string, v string) {
	o.Headers[k] = v
}
func (o *EbayCall) GetHeader(k string) string {
	return o.Headers[k]
}

func (o *EbayCall) Execute(r *[]Result) error {
	o.CallDepth = 0
	cl := o.GetCallname()
	if cl == "GeteBayOfficialTime" {
		err := o.GeteBayOfficialTime(r)
		if err != nil {
			appendFakeResult(fmt.Sprintf("%s", err), r)
			return err
		}
		return o.Send(r)
	}

	if cl == "GetAllCategories" {
		o.SetCallname("GetCategories")
		o.Callname = "GetAllCategories"
		err := o.GetAllCategories(r)
		if err != nil {
			appendFakeResult(fmt.Sprintf("%s", err), r)
			return err
		}
		return o.Send(r)
	}

	if cl == "AddItems" {
		return o.AddItems(r)
	}

	return nil
}
func (o *EbayCall) CollectAddItems() (*AddItemsStruct, error) {
	var s AddItemsStruct
	if o.HasItemsToSend() {
		for i := 0; i < len(o.Items); i++ {

			ci := o.Items[i]
			// we can to skip over items we've sent
			if ci.sent == true {
				continue
			}
			body, err := compileGoString("Item", ItemTemplate(), ci, nil)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s %v", err, ci))
			}

			s.Children = append(s.Children, AddItemsChild{Item: ci, Text: body, MessageID: ci.internal_id})
			if len(s.Children) == o.AddItemsLimit {
				return &s, nil
			}
		}
		return &s, nil
	}
	return nil, errors.New("Got to the end of CollectAddItems")
}
func (o *EbayCall) CollectAddItemsXML(s *AddItemsStruct) (string, error) {
	items_xml, err := compileGoString("AddItems", AddItemsTemplate(), s, nil)
	if err != nil {
		return "", err
	}
	final_xml, err := compileGoString(
		"FinalAddItems",
		WrapCall(
			"AddItems",
			"",
			items_xml,
			"",
		),
		o,
		nil,
	)
	if err != nil {
		return "", err
	}
	return final_xml, nil
}
func (o *EbayCall) AddItems(r *[]Result) error {
	// helps to prevent flooding the server if you do shit wrong
	if o.CallDepth >= o.CallDepthLimit {
		err := errors.New(fmt.Sprintf("CallDepthLimit of %d reached for AddItems!", o.CallDepthLimit))
		appendFakeResult(fmt.Sprintf("%s", err), r)
		return err
	}

	o.CallDepth = o.CallDepth + 1
	var tr []Result

	s, err := o.CollectAddItems()

	if err != nil {
		appendFakeResult(fmt.Sprintf("%s", err), r)
		return err
	}

	if len(s.Children) > len(o.Items) {
		err := errors.New(fmt.Sprintf("Too many Children %d!", len(s.Children)))
		appendFakeResult(fmt.Sprintf("%s", err), r)
		return err
	}

	o.XMLData, err = o.CollectAddItemsXML(s)
	if err != nil {
		appendFakeResult(fmt.Sprintf("%s", err), r)
		return err
	}
	o.MessageID, _ = pseudoUUID()
	o.MessageIDs = append(o.MessageIDs, o.MessageID)
	err = o.Send(&tr)

	for _, cr := range tr {
		*r = append(*r, cr)
	}
	for _, child := range s.Children {
		child.Item.sent = true
	}
	if err != nil { // the err from o.Send
		appendFakeResult(fmt.Sprintf("%s", err), r)
		return err
	}
	if o.HasItemsToSend() {
		return o.AddItems(r)
	}
	return nil
}
func (o *EbayCall) HasItemsToSend() bool {
	for _, i := range o.Items {
		if i.sent == false && i.failed != true {
			return true
		}
	}
	return false
}
func appendFakeResult(msg string, r *[]Result) {
	e := NewFakeResult(fmt.Sprintf("%s", msg))
	*r = append(*r, *e)
}
func (o *EbayCall) Send(r *[]Result) error {
	o.TheClient = new(http.Client)

	globalDebugFunction(DBG_DEBUG, fmt.Sprintf("About to send [[%s]]\n\n", o.XMLData))

	if o.XMLData == "" {
		err := errors.New("XMLData was empty!")
		e := NewFakeResult(fmt.Sprintf("%s", err))
		*r = append(*r, *e)
		return err
	}
	filePutContents(fmt.Sprintf("%s/last-sent.xml", o.Cache), o.XMLData)
	req, err := http.NewRequest("POST", o.EndPoint, bytes.NewBufferString(o.XMLData))
	if err != nil {
		e := NewFakeResult(fmt.Sprintf("%s", err))
		*r = append(*r, *e)
		return err
	}
	//req.URL.Host = "148.251.124.116:9090"
	for k, v := range o.Headers {
		req.Header.Set(k, v)
	}

	resp, err := o.TheClient.Do(req)
	//Post(o.EndPoint, "text/xml; charset=utf-8", )
	if err != nil {
		e := NewFakeResult(fmt.Sprintf("%s", err))
		*r = append(*r, *e)
		return err
	}
	//fmt.Printf("%+v\n", resp)
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	// We should cache the results of certain calls

	globalDebugFunction(DBG_DEBUG, fmt.Sprintf("[[BODY: %s]]", string(b)))
	if err != nil {
		e := NewFakeResult(fmt.Sprintf("%s", err))
		*r = append(*r, *e)
		return err
	}
	if !fileExists(o.Cache) {
		err = os.Mkdir(o.Cache, 0777)
		if err != nil {
			globalDebugFunction(DBG_WARN, "Could not create .cache")
		}
	}
	if fileExists(o.Cache) {
		if o.Callname == "GetAllCategories" {
			// GetAllCategories is a semi-special case, because we really want
			// to cache this result, and rarely if ever update it
			filePutContents(fmt.Sprintf("%s/%s.xml", o.Cache, o.Callname), string(b))
		} else {
			filePutContents(fmt.Sprintf("%s/%s-%s.xml", o.Cache, o.GetCallname(), o.MessageID), string(b))
		}
	}
	res, err := NewResult(b)
	if err != nil {
		res = NewFakeResult(fmt.Sprintf("%s", err))
	}
	*r = append(*r, *res)
	return nil
}

// Ebay Calls

func (o *EbayCall) GeteBayOfficialTime(r *[]Result) error {
	o.MessageID, _ = pseudoUUID()
	body, err := compileGoString("Time", GeteBayOfficialTimeTemplate(), o, nil)
	if err != nil {
		return err
	}
	final_xml, err := compileGoString("FinalTime", WrapCall("GeteBayOfficialTime", "", body, ""), o, nil)
	if err != nil {
		return err
	}
	o.XMLData = final_xml
	return nil
}

func (o *EbayCall) GetAllCategories(r *[]Result) error {
	o.MessageID, _ = pseudoUUID()
	o.CategoryCallInfo.SiteID = o.SiteID
	if o.CategoryCallInfo.LevelLimit == "" {
		o.CategoryCallInfo.LevelLimit = "3"
	}
	if o.CategoryCallInfo.ViewAllNodes != "true" {
		o.CategoryCallInfo.ViewAllNodes = "false"
	}

	body, err := compileGoString("Time", GetAllCategoriesTemplate(), o.CategoryCallInfo, nil)
	if err != nil {
		appendFakeResult(fmt.Sprintf("%s", err), r)
		return err
	}
	final_xml, err := compileGoString("FinalGetAllCategories", WrapCall("GetCategories", "", body, ""), o, nil)
	if err != nil {
		appendFakeResult(fmt.Sprintf("%s", err), r)
		return err
	}
	o.XMLData = final_xml
	return nil
}

// Getters and Setters

type ItemFilter func(o *Item) bool

func NewEbayCall() *EbayCall {
	return &EbayCall{}
}

func (o *EbayCall) Clone() *EbayCall {
	var no EbayCall
	no.DevID = o.DevID
	no.AppID = o.AppID
	no.CertID = o.CertID
	no.CompatLevel = o.CompatLevel
	no.SiteID = o.SiteID
	no.EndPoint = o.EndPoint
	no.EbayAuthToken = o.EbayAuthToken
	no.Country = o.Country
	no.Currency = o.Currency
	no.Language = o.Language
	no.MessageID = o.MessageID
	no.WarningLevel = o.WarningLevel
	no.PayPalEmailAddress = o.PayPalEmailAddress
	no.Callname = o.Callname
	no.XMLData = o.XMLData
	no.Headers = o.Headers
	no.Items = o.Items
	no.CategoryCallInfo = o.CategoryCallInfo
	return &no
}

func (o *EbayCall) SetDevID(v string) {
	o.DevID = v
}

func (o *EbayCall) GetDevID() string {
	return o.DevID
}

func (o *EbayCall) SetAppID(v string) {
	o.AppID = v
}

func (o *EbayCall) GetAppID() string {
	return o.AppID
}

func (o *EbayCall) SetCertID(v string) {
	o.CertID = v
}

func (o *EbayCall) GetCertID() string {
	return o.CertID
}

func (o *EbayCall) SetCompatLevel(v string) {
	o.CompatLevel = v
}

func (o *EbayCall) GetCompatLevel() string {
	return o.CompatLevel
}

func (o *EbayCall) SetSiteID(v string) {
	o.SiteID = v
}

func (o *EbayCall) GetSiteID() string {
	return o.SiteID
}

func (o *EbayCall) SetEndPoint(v string) {
	o.EndPoint = v
}

func (o *EbayCall) GetEndPoint() string {
	return o.EndPoint
}

func (o *EbayCall) SetEbayAuthToken(v string) {
	o.EbayAuthToken = v
}

func (o *EbayCall) GetEbayAuthToken() string {
	return o.EbayAuthToken
}

func (o *EbayCall) SetCountry(v string) {
	o.Country = v
}

func (o *EbayCall) GetCountry() string {
	return o.Country
}

func (o *EbayCall) SetCurrency(v string) {
	o.Currency = v
}

func (o *EbayCall) GetCurrency() string {
	return o.Currency
}

func (o *EbayCall) SetLanguage(v string) {
	o.Language = v
}

func (o *EbayCall) GetLanguage() string {
	return o.Language
}

func (o *EbayCall) SetMessageID(v string) {
	o.MessageID = v
}

func (o *EbayCall) GetMessageID() string {
	return o.MessageID
}

func (o *EbayCall) SetWarningLevel(v string) {
	o.WarningLevel = v
}

func (o *EbayCall) GetWarningLevel() string {
	return o.WarningLevel
}

func (o *EbayCall) SetPayPalEmailAddress(v string) {
	o.PayPalEmailAddress = v
}

func (o *EbayCall) GetPayPalEmailAddress() string {
	return o.PayPalEmailAddress
}

func (o *EbayCall) SetXMLData(v string) {
	o.XMLData = v
}
func (o *EbayCall) SetCallname(v string) {
	o.Callname = v
	o.Headers["X-EBAY-API-CALL-NAME"] = v
}
func (o *EbayCall) GetCallname() string {
	return o.Headers["X-EBAY-API-CALL-NAME"]
}
func (o *EbayCall) GetXMLData() string {
	return o.XMLData
}

func (o *EbayCall) SetHeaders(v map[string]string) {
	o.Headers = v
}

func (o *EbayCall) GetHeaders() map[string]string {
	return o.Headers
}

func (o *EbayCall) FilterItems(f ItemFilter) []*Item {
	tmp := o.Items[:0]
	for _, x := range o.Items {
		if f(x) {
			tmp = append(tmp, x)
		}
	}
	return tmp
}

func (o *EbayCall) AddItem(v *Item) {
	v.internal_id, _ = pseudoUUID()
	o.Items = append(o.Items, v)
}

func (o *EbayCall) RemoveItem(i int) {
	if i > len(o.Items) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "EbayCall", "Items", len(o.Items)))
	}
	o.Items = o.Items[:i+copy(o.Items[i:], o.Items[i+1:])]
}

func (o *EbayCall) GetItem(i int) *Item {
	if i > len(o.Items) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "EbayCall", "Items", len(o.Items)))
	}
	return o.Items[i]
}

func (o *EbayCall) GetItemByInternalID(id string) (*Item, error) {
	for _, i := range o.Items {
		if i.internal_id == id {
			return i, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("No item found with id: %s", id))
}

func (o *EbayCall) SetItems(v []*Item) {
	o.Items = v
}

func (o *EbayCall) GetItems() []*Item {
	return o.Items
}

func (o *EbayCall) SetCategoryCallInfo(v *GetCategoriesStruct) {
	o.CategoryCallInfo = v
}

func (o *EbayCall) GetCategoryCallInfo() *GetCategoriesStruct {
	return o.CategoryCallInfo
}

// Debug Functions

func (o_EbayCall *EbayCall) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sEbayCall.DevID: %s\n", txt, o_EbayCall.DevID)
	txt = fmt.Sprintf("%sEbayCall.AppID: %s\n", txt, o_EbayCall.AppID)
	txt = fmt.Sprintf("%sEbayCall.CertID: %s\n", txt, o_EbayCall.CertID)
	txt = fmt.Sprintf("%sEbayCall.CompatLevel: %s\n", txt, o_EbayCall.CompatLevel)
	txt = fmt.Sprintf("%sEbayCall.SiteID: %s\n", txt, o_EbayCall.SiteID)
	txt = fmt.Sprintf("%sEbayCall.EndPoint: %s\n", txt, o_EbayCall.EndPoint)
	txt = fmt.Sprintf("%sEbayCall.EbayAuthToken: %s\n", txt, o_EbayCall.EbayAuthToken)
	txt = fmt.Sprintf("%sEbayCall.Country: %s\n", txt, o_EbayCall.Country)
	txt = fmt.Sprintf("%sEbayCall.Currency: %s\n", txt, o_EbayCall.Currency)
	txt = fmt.Sprintf("%sEbayCall.Language: %s\n", txt, o_EbayCall.Language)
	txt = fmt.Sprintf("%sEbayCall.MessageID: %s\n", txt, o_EbayCall.MessageID)
	txt = fmt.Sprintf("%sEbayCall.WarningLevel: %s\n", txt, o_EbayCall.WarningLevel)
	txt = fmt.Sprintf("%sEbayCall.PayPalEmailAddress: %s\n", txt, o_EbayCall.PayPalEmailAddress)
	txt = fmt.Sprintf("%sEbayCall.Callname: %s\n", txt, o_EbayCall.Callname)
	txt = fmt.Sprintf("%sEbayCall.XMLData: %s\n", txt, o_EbayCall.XMLData)
	for k, v := range o_EbayCall.Headers {
		txt = fmt.Sprintf("%sEbayCall.Headers [%s]: %s\n", txt, k, v)
	}
	for i, v := range o_EbayCall.Items {
		txt = fmt.Sprintf("%sEbayCall.Items [%d]: %s\n", txt, i, v.Debug())
	}
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}
