package gobay

import "testing"

func TestResult(t *testing.T) {
	data := `<?xml version="1.0" encoding="UTF-8"?>
<AddFixedPriceItemResponse xmlns="urn:ebay:apis:eBLBaseComponents"><Timestamp>2016-09-26T07:38:18.534Z</Timestamp><Ack>Warning</Ack><Errors><ShortMessage>Return Policy input not applicable.</ShortMessage><LongMessage>The Return Policy field Refund in the input has been ignored.</LongMessage><ErrorCode>21916711</ErrorCode><SeverityCode>Warning</SeverityCode><ErrorParameters ParamID="0"><Value>Refund</Value></ErrorParameters><ErrorClassification>RequestError</ErrorClassification></Errors><Errors><ShortMessage>Postage cost information incomplete.</ShortMessage><LongMessage>You did not provide a value for additional postage cost.  The same value you provided for delivery cost has been copied to the additional postage cost field.</LongMessage><ErrorCode>219026</ErrorCode><SeverityCode>Warning</SeverityCode><ErrorClassification>RequestError</ErrorClassification></Errors><Errors><ShortMessage>Seller profiles will soon be mandatory when creating a new listing for sellers opted in to business policies.</ShortMessage><LongMessage>Seller Profiles will soon be mandatory while creating a new listing. Support for legacy Postage, Payment and Returns fields will be removed for sellers opted in to business policies.</LongMessage><ErrorCode>21919456</ErrorCode><SeverityCode>Warning</SeverityCode><ErrorClassification>RequestError</ErrorClassification></Errors><Version>987</Version><Build>E987_UNI_API5_18120651_R1</Build><ItemID>252557106431</ItemID><SKU>M147675</SKU><StartTime>2016-09-26T07:38:16.159Z</StartTime><EndTime>2016-10-26T07:38:16.159Z</EndTime>
<Fees>
	<Fee>
		<Name>AuctionLengthFee</Name>
		<Fee currencyID="GBP">0.0</Fee>
	</Fee>
	<Fee>
		<Name>BoldFee</Name>
		<Fee currencyID="GBP">0.0</Fee>
	</Fee>
		<Fee><Name>BuyItNowFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>CategoryFeaturedFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>FeaturedFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>GalleryPlusFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>FeaturedGalleryFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>FixedPriceDurationFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>GalleryFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>GiftIconFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>HighLightFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>InsertionFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>InternationalInsertionFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>ListingDesignerFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>ListingFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>PhotoDisplayFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>PhotoFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>ReserveFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>SchedulingFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>SubtitleFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>BorderFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>ProPackBundleFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>BasicUpgradePackBundleFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>ValuePackBundleFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>PrivateListingFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>ProPackPlusBundleFee</Name><Fee currencyID="GBP">0.0</Fee></Fee><Fee><Name>MotorsGermanySearchFee</Name><Fee currencyID="GBP">0.0</Fee></Fee></Fees></AddFixedPriceItemResponse>`
	o, err := NewResult([]byte(data))
	if err != nil {
		t.Errorf("NewResult returned error %+v!!\n", err)
	}
	if o.Timestamp != "2016-09-26T07:38:18.534Z" {
		t.Errorf("Result has not been filled out %+v!!\n", o.Timestamp)
	}
	if o.Ack != "Warning" {
		t.Errorf("Result has not been filled out %+v!!\n", o.Ack)
	}
	if o.Errors[0].ShortMessage != "Return Policy input not applicable." {
		t.Errorf("o.Errors[0].ShortMessage has not been filled out!! %+v\n", o.Errors[0])
	}
	if o.Errors[0].LongMessage != "The Return Policy field Refund in the input has been ignored." {
		t.Errorf("o.Errors[0].LongMessage has not been filled out!! %+v\n", o.Errors[0])
	}
	if o.Errors[0].ErrorCode != 21916711 {
		t.Errorf("o.Errors[0].ErrorCode has not been filled out!! %+v\n", o.Errors[0])
	}
	if o.Errors[0].SeverityCode != "Warning" {
		t.Errorf("o.Errors[0].SeverityCode has not been filled out!! %+v\n", o.Errors[0])
	}
	if o.Errors[0].ErrorClassification != "RequestError" {
		t.Errorf("o.Errors[0].ErrorClassification has not been filled out!! %+v\n", o.Errors[0])
	}
	if o.Errors[1].ShortMessage != "Postage cost information incomplete." {
		t.Errorf("o.Errors[1].ShortMessage has not been filled out!! %+v\n", o.Errors[1])
	}
	if o.Errors[1].LongMessage != "You did not provide a value for additional postage cost.  The same value you provided for delivery cost has been copied to the additional postage cost field." {
		t.Errorf("o.Errors[1].LongMessage has not been filled out!! %+v\n", o.Errors[1])
	}
	if o.Errors[1].ErrorCode != 219026 {
		t.Errorf("o.Errors[1].ErrorCode has not been filled out!! %+v\n", o.Errors[1])
	}
	if o.Errors[1].SeverityCode != "Warning" {
		t.Errorf("o.Errors[1].SeverityCode has not been filled out!! %+v\n", o.Errors[1])
	}
	if o.Errors[1].ErrorClassification != "RequestError" {
		t.Errorf("o.Errors[1].ErrorClassification has not been filled out!! %+v\n", o.Errors[1])
	}
	if o.Errors[2].ShortMessage != "Seller profiles will soon be mandatory when creating a new listing for sellers opted in to business policies." {
		t.Errorf("o.Errors[2].ShortMessage has not been filled out!! %+v\n", o.Errors[2])
	}
	if o.Errors[2].LongMessage != "Seller Profiles will soon be mandatory while creating a new listing. Support for legacy Postage, Payment and Returns fields will be removed for sellers opted in to business policies." {
		t.Errorf("o.Errors[2].LongMessage has not been filled out!! %+v\n", o.Errors[2])
	}
	if o.Errors[2].ErrorCode != 21919456 {
		t.Errorf("o.Errors[2].ErrorCode has not been filled out!! %+v\n", o.Errors[2])
	}
	if o.Errors[2].SeverityCode != "Warning" {
		t.Errorf("o.Errors[2].SeverityCode has not been filled out!! %+v\n", o.Errors[2])
	}
	if o.Errors[2].ErrorClassification != "RequestError" {
		t.Errorf("o.Errors[2].ErrorClassification has not been filled out!! %+v\n", o.Errors[2])
	}
	if o.Version != "987" {
		t.Errorf("Result has not been filled out %+v!!\n", o.Version)
	}
	if o.Build != "E987_UNI_API5_18120651_R1" {
		t.Errorf("Result has not been filled out %+v!!\n", o.Build)
	}
	if o.ItemID != "252557106431" {
		t.Errorf("Result has not been filled out %+v!!\n", o.ItemID)
	}
	if o.SKU != "M147675" {
		t.Errorf("Result has not been filled out %+v!!\n", o.SKU)
	}
	if o.StartTime != "2016-09-26T07:38:16.159Z" {
		t.Errorf("Result has not been filled out %+v!!\n", o.StartTime)
	}
	if o.EndTime != "2016-10-26T07:38:16.159Z" {
		t.Errorf("Result has not been filled out %+v!!\n", o.EndTime)
	}
	if o.Fees[0].Amount != "0.0" {
		t.Errorf("Result.Fees[0] has not been filled out %+v!!\n", o.Fees)
	}
}

/* We need to test it with more results, to make sure they're all parsable. */
func TestResultTwo(t *testing.T) {
	data := `<?xml version="1.0" encoding="UTF-8"?>
<EndItemResponse xmlns="urn:ebay:apis:eBLBaseComponents"><Timestamp>2016-09-26T06:44:59.344Z</Timestamp><Ack>Failure</Ack><Errors><ShortMessage>The auction has been closed.</ShortMessage><LongMessage>The auction has already been closed.</LongMessage><ErrorCode>1047</ErrorCode><SeverityCode>Error</SeverityCode><ErrorClassification>RequestError</ErrorClassification></Errors><Version>987</Version><Build>E987_UNI_API5_18120651_R1</Build></EndItemResponse>`
	o, err := NewResult([]byte(data))

	if err != nil {
		t.Errorf("NewResult returned error %+v!!\n", err)
	}
	if o.Ack != "Failure" {
		t.Errorf("Result has not been filled out %+v!!\n", o.Ack)
	}
	if o.Errors[0].ShortMessage != "The auction has been closed." {
		t.Errorf("o.Errors[0].ShortMessage has not been filled out!! %+v\n", o.Errors[0])
	}
}
func TestAddItemsResponseParsing(t *testing.T) {
	data, err := fileGetContents("test_data/AddItemsResponse.xml")
	if err != nil {
		t.Errorf("fileGetContents returned error %+v!!\n", err)
	}
	o, err := NewResult(data)
	if err != nil {
		t.Errorf("NewResult returned error %+v!!\n", err)
	}
	if len(o.Items) < 2 {
		t.Errorf("There should be two items!!\n")
	}
	if o.Items[0].ItemID != "FIRSTITEMID" {
		t.Errorf("ItemID wrong!!\n")
	}
	if o.Items[1].ItemID != "SECONDITEMID" {
		t.Errorf("ItemID wrong!!\n")
	}

}

func TestGetAllCategoriesResponseParsing(t *testing.T) {
	data, err := fileGetContents("test_data/GetAllCategoriesResponse.xml")
	if err != nil {
		t.Errorf("fileGetContents returned error %+v!!\n", err)
		return
	}
	o, err := NewResult(data)
	if err != nil {
		t.Errorf("NewResult returned error %+v!!\n", err)
		return
	}
	if len(o.Categories) < 1 {
		t.Errorf("There should be two items!!\n")
		return
	}
	if o.Categories[0].CategoryID != "66840" {
		t.Errorf("The first category in this result should be: %s not %s!!\n", "66840", o.Categories[0].CategoryID)
		return
	}
}

func TestNotificationResultParsing(t *testing.T) {
	data, err := fileGetContents("test_data/GetFeedbackResponse.xml")
	if err != nil {
		t.Errorf("fileGetContents returned error %+v!!\n", err)
		return
	}
	o, err := NewNotificationResult(data)
	if err != nil {
		t.Errorf("NewNotificationResult returned error %+v!!\n", err)
		return
	}
	t.Errorf("%+v", o)
	return
}
