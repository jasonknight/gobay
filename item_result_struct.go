package gobay

type VariationStruct struct {
	PrivateNotes   string
	Quantity       string
	SKU            string
	StartPrice     string
	VariationTitle string
	WatchCount     string
	SellingStatus  SellingStatusStruct

	VariationSpecifics []struct {
		NameValueList []struct {
			Name   string
			Values []string `xml:"Value"`
		}
	}
}
type ListingDetailsStruct struct {
	ConvertedBuyItNowPrice string
	ConvertedReservePrice  string
	ConvertedStartPrice    string
	StartTime              string
	EndTime                string
}
type SellingStatusStruct struct {
	ConvertedCurrentPrice string
	CurrentPrice          string
	QuantitySold          string
	ReserveMet            string
	BidCount              string
	BidderCount           string
	HighBidder            struct {
		FeedbackRatingStar string
		FeedbackScore      string
		UserID             string
	}
	PromotionalSaleDetails struct {
		EndTime       string
		OriginalPrice string
		StartTime     string
	}
}
type ItemResultStruct struct {
	BestOfferDetails struct {
		BestOfferCount string
	}
	BuyItNowPrice             string
	ClassifiedAdPayPerLeadFee string
	eBayNotes                 string
	HideFromSearch            string
	ID                        string
	ListingDetails            ListingDetailsStruct
	ListingDuration           string
	ListingType               string
	PrivateNotes              string
	Quantity                  string
	QuantityAvailable         string
	QuestionCount             string
	ReasonHideFromSearch      string
	ReservePrice              string
	SKU                       string
	StartPrice                string
	TimeLeft                  string
	Title                     string
	WatchCount                string
	PictureDetails            PictureDetail
	SellingStatus             SellingStatusStruct
	ShippingDetails           ShippingDetailsStruct
	Variations                []VariationStruct `xml:"Variations>Variation"`
}

type ShippingDetailsStruct struct {
	GlobalShipping         string
	ShippingType           string
	ShippingServiceOptions []struct {
		LocalPickup         string
		LogisticPlanType    string
		ShippingServiceCost string
		ShippingSurcharge   string
	}
}
