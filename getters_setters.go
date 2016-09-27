package gobay

import "fmt"

type ProductFilter func(o Product) bool

func NewPictureDetail() *PictureDetail {
	return &PictureDetail{}
}

func (o *PictureDetail) Clone() *PictureDetail {
	var no PictureDetail
	no.GalleryDuration = o.GalleryDuration
	no.GalleryType = o.GalleryType
	no.GalleryURL = o.GalleryURL
	no.PhotoDisplay = o.PhotoDisplay
	no.PictureURL = o.PictureURL
	return &no
}

func (o *PictureDetail) SetGalleryDuration(v string) {
	o.GalleryDuration = v
}

func (o *PictureDetail) GetGalleryDuration() string {
	return o.GalleryDuration
}

func (o *PictureDetail) SetGalleryType(v string) {
	o.GalleryType = v
}

func (o *PictureDetail) GetGalleryType() string {
	return o.GalleryType
}

func (o *PictureDetail) SetGalleryURL(v string) {
	o.GalleryURL = v
}

func (o *PictureDetail) GetGalleryURL() string {
	return o.GalleryURL
}

func (o *PictureDetail) SetPhotoDisplay(v string) {
	o.PhotoDisplay = v
}

func (o *PictureDetail) GetPhotoDisplay() string {
	return o.PhotoDisplay
}

func (o *PictureDetail) SetPictureURL(v string) {
	o.PictureURL = v
}

func (o *PictureDetail) GetPictureURL() string {
	return o.PictureURL
}

func NewShippingServiceOption() *ShippingServiceOption {
	return &ShippingServiceOption{}
}

func NewEbayCall() *EbayCall {
	return &EbayCall{}
}

func (o *ShippingServiceOption) Clone() *ShippingServiceOption {
	var no ShippingServiceOption
	no.Service = o.Service
	no.Cost = o.Cost
	no.Priority = o.Priority
	return &no
}

func (o *ShippingServiceOption) SetService(v string) {
	o.Service = v
}

func (o *ShippingServiceOption) GetService() string {
	return o.Service
}

func (o *ShippingServiceOption) SetCost(v string) {
	o.Cost = v
}

func (o *ShippingServiceOption) GetCost() string {
	return o.Cost
}

func (o *ShippingServiceOption) SetPriority(v string) {
	o.Priority = v
}

func (o *ShippingServiceOption) GetPriority() string {
	return o.Priority
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
	no.PayPalEmailAddress = o.PayPalEmailAddress
	no.Headers = o.Headers
	no.Products = o.Products
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

func (o *EbayCall) SetPayPalEmailAddress(v string) {
	o.PayPalEmailAddress = v
}

func (o *EbayCall) GetPayPalEmailAddress() string {
	return o.PayPalEmailAddress
}

func (o *EbayCall) SetHeaders(v map[string]string) {
	o.Headers = v
}

func (o *EbayCall) GetHeaders() map[string]string {
	return o.Headers
}

func (o *EbayCall) FilterProducts(f ProductFilter) []Product {
	tmp := o.Products[:0]
	for _, x := range o.Products {
		if f(x) {
			tmp = append(tmp, x)
		}
	}
	return tmp
}

func (o *EbayCall) AddProduct(v Product) {
	o.Products = append(o.Products, v)
}

func (o *EbayCall) RemoveProduct(i int) {
	if i > len(o.Products) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "EbayCall", "Products", len(o.Products)))
	}
	o.Products = o.Products[:i+copy(o.Products[i:], o.Products[i+1:])]
}

func (o *EbayCall) GetProduct(i int) Product {
	if i > len(o.Products) {
		panic(fmt.Sprintf("i:%d is out of bounds for %s.%s(%d)!\n", "EbayCall", "Products", len(o.Products)))
	}
	return o.Products[i]
}

func (o *EbayCall) SetProducts(v []Product) {
	o.Products = v
}

func (o *EbayCall) GetProducts() []Product {
	return o.Products
}
