package gobay

//import "fmt"

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

