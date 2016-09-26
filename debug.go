package gobay

import "fmt"

func (o_PictureDetail *PictureDetail) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sPictureDetail.GalleryDuration: %s\n", txt, o_PictureDetail.GalleryDuration)
	txt = fmt.Sprintf("%sPictureDetail.GalleryType: %s\n", txt, o_PictureDetail.GalleryType)
	txt = fmt.Sprintf("%sPictureDetail.GalleryURL: %s\n", txt, o_PictureDetail.GalleryURL)
	txt = fmt.Sprintf("%sPictureDetail.PhotoDisplay: %s\n", txt, o_PictureDetail.PhotoDisplay)
	txt = fmt.Sprintf("%sPictureDetail.PictureURL: %s\n", txt, o_PictureDetail.PictureURL)
	return txt
}

func (o_ShippingServiceOption *ShippingServiceOption) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sShippingServiceOption.Service: %s\n", txt, o_ShippingServiceOption.Service)
	txt = fmt.Sprintf("%sShippingServiceOption.Cost: %s\n", txt, o_ShippingServiceOption.Cost)
	txt = fmt.Sprintf("%sShippingServiceOption.Priority: %s\n", txt, o_ShippingServiceOption.Priority)
	return txt
}

func (o_Product *Product) Debug() string {
	var txt string
	txt = fmt.Sprintf("%sProduct.SKU: %s\n", txt, o_Product.SKU)
	txt = fmt.Sprintf("%sProduct.Title: %s\n", txt, o_Product.Title)
	txt = fmt.Sprintf("%sProduct.Price: %s\n", txt, o_Product.Price)
	txt = fmt.Sprintf("%sProduct.ListingType: %s\n", txt, o_Product.ListingType)
	for _, v := range o_Product.ShipToLocations {
		txt = fmt.Sprintf("%s%s\n", txt, v)
	}
	for _, v := range o_Product.PictureDetails {
		txt = fmt.Sprintf("%s%s\n", txt, v.Debug())
	}
	for _, v := range o_Product.PaymentMethods {
		txt = fmt.Sprintf("%s%s\n", txt, v)
	}
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
	txt = fmt.Sprintf("%sEbayCall.PaypalEmailAddress: %s\n", txt, o_EbayCall.PaypalEmailAddress)
	for k, v := range o_EbayCall.Headers {
		txt = fmt.Sprintf("%sEbayCall.Headers [%s]: %s\n", txt, k, v)
	}
	for _, v := range o_EbayCall.Products {
		txt = fmt.Sprintf("%s%s\n", txt, v.Debug())
	}
	return txt
}
