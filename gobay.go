package gobay

import (
    "fmt"
    "net/http"
)

type ShipToLocation string
type PictureDetail string
type PaymentMethod string
type ShippingServiceOption struct {
    Service string
    Cost string
    Priority string
}
type Product struct {
    SKU string
    Title string
    Price string
    []ShipToLocations ShipToLocation
    []PictureDetails PictureDetail
    []PaymentMethods PaymentMethod
}
type EbayCall struct {
    DevID string
    AppID string
    CertID string
    CompatLevel string
    SiteID string
    EndPoint string
    EbayAuthToken string
    Country string
    Currency string
    PaypalEmailAddress string
    Headers map[string]string
    TheProduct Product
    TheRequest http.Request
}

func NewEbayCall(endPoint, compatLevel, devID, appID, certID, siteID, callName string, p Product) *EbayCall {
    var e EbayCall
    m := make(map[string]string)
    m["X-EBAY-API-COMPATIBILITY-LEVEL"] = fmt.Sprintf("%s",compatLevel)
    m["X-EBAY-API-DEV-NAME"] = fmt.Sprintf("%s",devID)
    m["X-EBAY-API-APP-NAME"] = fmt.Sprintf("%s", appID)
    m["X-EBAY-API-CERT-NAME"] = fmt.Sprintf("%s", certID)
    m["X-EBAY-API-CALL-NAME"] = fmt.Sprintf("%s", callName)
    m["X-EBAY-API-SITEID"] = fmt.Sprintf("%s",siteID)
    e.Headers = m
    e.DevID = devID
    e.AppID = appID
    e.CertID = certID
    e.SiteID = siteID
    e.TheProduct = p
    // need to decide body from io.reader...so we need to compile the template
    e.TheRequest = NewRequest("POST",endPoint)
    return &e
}