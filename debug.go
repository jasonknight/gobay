package gobay

import "fmt"

const (
	DBG_DEBUG = iota
	DBG_INFO
	DBG_WARN
	DBG_ERROR
	DBG_NONE
)

type DebugFunc func(lvl int, s string)

var globalDebugFunction DebugFunc
var globalDebugLevel int

func SetDebugLevel(l int) {
	globalDebugLevel = l
}

func SetDebugFunction(f DebugFunc) {
	globalDebugFunction = f
}

func (o_PictureDetail *PictureDetail) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sPictureDetail.GalleryDuration: %s\n", txt, o_PictureDetail.GalleryDuration)
	txt = fmt.Sprintf("%sPictureDetail.GalleryType: %s\n", txt, o_PictureDetail.GalleryType)
	txt = fmt.Sprintf("%sPictureDetail.GalleryURL: %s\n", txt, o_PictureDetail.GalleryURL)
	txt = fmt.Sprintf("%sPictureDetail.PhotoDisplay: %s\n", txt, o_PictureDetail.PhotoDisplay)
	txt = fmt.Sprintf("%sPictureDetail.PictureURL: %s\n", txt, o_PictureDetail.PictureURL)
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

func (o_ShippingServiceOption *ShippingServiceOption) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sShippingServiceOption.Service: %s\n", txt, o_ShippingServiceOption.Service)
	txt = fmt.Sprintf("%sShippingServiceOption.Cost: %s\n", txt, o_ShippingServiceOption.Cost)
	txt = fmt.Sprintf("%sShippingServiceOption.Priority: %s\n", txt, o_ShippingServiceOption.Priority)
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

func (o_Product *Item) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sProduct.SKU: %s\n", txt, o_Product.SKU)
	txt = fmt.Sprintf("%sProduct.Title: %s\n", txt, o_Product.Title)
	txt = fmt.Sprintf("%sProduct.StartPrice: %s\n", txt, o_Product.StartPrice)
	txt = fmt.Sprintf("%sProduct.ListingType: %s\n", txt, o_Product.ListingType)
	for _, v := range o_Product.ShipToLocations {
		txt = fmt.Sprintf("%s%s\n", txt, v)
	}
	for _, v := range o_Product.PictureDetails {
		txt = fmt.Sprintf("%s%+v\n", txt, v.Debug())
	}
	for _, v := range o_Product.PaymentMethods {
		txt = fmt.Sprintf("%s%s\n", txt, v)
	}
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}

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
	txt = fmt.Sprintf("%sEbayCall.PayPalEmailAddress: %s\n", txt, o_EbayCall.PayPalEmailAddress)
	for k, v := range o_EbayCall.Headers {
		txt = fmt.Sprintf("%sEbayCall.Headers [%s]: %s\n", txt, k, v)
	}
	for _, v := range o_EbayCall.Products {
		txt = fmt.Sprintf("%s%s\n", txt, v.Debug())
	}
	globalDebugFunction(DBG_DEBUG, txt)
	return txt
}
