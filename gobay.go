package gobay

import (
    "fmt"
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
}

func NewEbayCall(compatLevel, devID, appID, certID, siteID, callName string) EbayCall {
    var e EbayCall
    e.Headers = {
      fmt.Sprintf('X-EBAY-API-COMPATIBILITY-LEVEL:%s',compatLevel),
      fmt.Sprintf('X-EBAY-API-DEV-NAME:%s' , devID,
      fmt.Sprintf('X-EBAY-API-APP-NAME:%s', appID,
      fmt.Sprintf('X-EBAY-API-CERT-NAME:%s', certID,
      fmt.Sprintf('X-EBAY-API-CALL-NAME:%s', callName, 
      fmt.Sprintf('X-EBAY-API-SITEID:%s',siteID,
    }
}