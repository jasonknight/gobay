package gobay

import "testing"

func TestGetMyeBaySellingResult(t *testing.T) {
	data, err := fileGetContents("test_data/GetMyeBaySellingResponse.xml")

	if err != nil {
		t.Errorf("%s", err)
		return
	}
	o := GetMyeBaySellingResultStructFromXML([]byte(data))
	if o.SellingSummary.ActiveAuctionCount != " int " {
		t.Errorf("o.SellingSummary.ActiveAuctionCount has not been filled out!! %+v\n", o.SellingSummary)
	}
	if o.SellingSummary.AuctionBidCount != " int " {
		t.Errorf("o.SellingSummary.AuctionBidCount has not been filled out!! %+v\n", o.SellingSummary)
	}
	if o.SellingSummary.AuctionSellingCount != " int " {
		t.Errorf("o.SellingSummary.AuctionSellingCount has not been filled out!! %+v\n", o.SellingSummary)
	}
	if o.SellingSummary.SoldDurationInDays != " int " {
		t.Errorf("o.SellingSummary.SoldDurationInDays has not been filled out!! %+v\n", o.SellingSummary)
	}
	if o.SellingSummary.TotalAuctionSellingValue != " AmountType (double) " {
		t.Errorf("o.SellingSummary.TotalAuctionSellingValue has not been filled out!! %+v\n", o.SellingSummary)
	}
	if o.SellingSummary.TotalSoldCount != " int " {
		t.Errorf("o.SellingSummary.TotalSoldCount has not been filled out!! %+v\n", o.SellingSummary)
	}
	if o.SellingSummary.TotalSoldValue != " AmountType (double) " {
		t.Errorf("o.SellingSummary.TotalSoldValue has not been filled out!! %+v\n", o.SellingSummary)
	}
	if o.Summary.ActiveAuctionCount != " int " {
		t.Errorf("o.Summary.ActiveAuctionCount has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.AmountLimitRemaining != " AmountType (double) " {
		t.Errorf("o.Summary.AmountLimitRemaining has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.AuctionBidCount != " int " {
		t.Errorf("o.Summary.AuctionBidCount has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.AuctionSellingCount != " int " {
		t.Errorf("o.Summary.AuctionSellingCount has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.ClassifiedAdCount != " int " {
		t.Errorf("o.Summary.ClassifiedAdCount has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.ClassifiedAdOfferCount != " int " {
		t.Errorf("o.Summary.ClassifiedAdOfferCount has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.QuantityLimitRemaining != " long " {
		t.Errorf("o.Summary.QuantityLimitRemaining has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.SoldDurationInDays != " int " {
		t.Errorf("o.Summary.SoldDurationInDays has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.TotalAuctionSellingValue != " AmountType (double) " {
		t.Errorf("o.Summary.TotalAuctionSellingValue has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.TotalLeadCount != " int " {
		t.Errorf("o.Summary.TotalLeadCount has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.TotalListingsWithLeads != " int " {
		t.Errorf("o.Summary.TotalListingsWithLeads has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.TotalSoldCount != " int " {
		t.Errorf("o.Summary.TotalSoldCount has not been filled out!! %+v\n", o.Summary)
	}
	if o.Summary.TotalSoldValue != " AmountType (double) " {
		t.Errorf("o.Summary.TotalSoldValue has not been filled out!! %+v\n", o.Summary)
	}
	if o.Ack != " AckCodeType " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n", o.Ack)
	}
	if o.Build != " string " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n", o.Build)
	}
	if o.CorrelationID != " string " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n", o.CorrelationID)
	}
	if o.Errors[0].ErrorClassification != " ErrorClassificationCodeType " {
		t.Errorf("o.Errors[0].ErrorClassification has not been filled out!! %+v\n", o.Errors[0])
	}
	if o.Errors[0].ErrorCode != 700 {
		t.Errorf("o.Errors[0].ErrorCode has not been filled out!! %+v\n", o.Errors[0])
	}
	if o.Errors[0].LongMessage != " string " {
		t.Errorf("o.Errors[0].LongMessage has not been filled out!! %+v\n", o.Errors[0])
	}
	if o.Errors[0].SeverityCode != " SeverityCodeType " {
		t.Errorf("o.Errors[0].SeverityCode has not been filled out!! %+v\n", o.Errors[0])
	}
	if o.Errors[0].ShortMessage != " string " {
		t.Errorf("o.Errors[0].ShortMessage has not been filled out!! %+v\n", o.Errors[0])
	}
	if o.HardExpirationWarning != " string " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n", o.HardExpirationWarning)
	}
	if o.Timestamp != " dateTime " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n", o.Timestamp)
	}
	if o.Version != " string " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n", o.Version)
	}

	if len(o.ActiveList) == 0 {
		t.Errorf("ActiveList length is 0")
		return
	}

	if len(o.ActiveList[0].Items) == 0 {
		t.Errorf("Items length is 0")
		return
	}

	if o.ActiveList[0].Items[0].BestOfferDetails.BestOfferCount != "int" {
		t.Errorf("BestOfferCount: %+v -- '%s' \n", o.ActiveList[0].Items[0].BestOfferDetails, o.ActiveList[0].Items[0].BestOfferDetails.BestOfferCount)
		return
	}

	if o.ActiveList[0].Items[0].BuyItNowPrice != " AmountType (double) " {
		t.Errorf("%+v\n", o.ActiveList[0].Items[0].BuyItNowPrice)
		return
	}

	if o.ActiveList[0].Items[0].ListingDetails.ConvertedBuyItNowPrice != " AmountType (double) " {
		t.Errorf("%+v\n", o.ActiveList[0].Items[0].ListingDetails)
		return
	}

	if o.ActiveList[0].Items[0].PictureDetails.GalleryURL != " anyURI " {
		t.Errorf("%+v\n", o.ActiveList[0].Items[0].PictureDetails)
		return
	}

	if len(o.ActiveList[0].Items[0].Variations) == 0 {
		t.Errorf("Items.Variations length is 0")
		return
	}

	if o.ActiveList[0].Items[0].Variations[0].SellingStatus.BidCount != " int " {
		t.Errorf("%+v\n", o.ActiveList[0].Items[0].Variations[0].SellingStatus)
		return
	}

}


