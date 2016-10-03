package gobay
import "testing"
func TestTemplate_get_my_ebay_selling(t *testing.T) {
	data := `<?xml version="1.0" encoding="utf-8"?>
<GetMyeBaySellingRequest xmlns="urn:ebay:apis:eBLBaseComponents">
  <!-- Call-specific Input Fields -->
  <ActiveList> 
    <Include> boolean </Include>
    <IncludeNotes> boolean </IncludeNotes>
    <ListingType> ListingTypeCodeType </ListingType>
    <Pagination> PaginationType
      <EntriesPerPage> int </EntriesPerPage>
      <PageNumber> int </PageNumber>
    </Pagination>
    <Sort> ItemSortTypeCodeType </Sort>
  </ActiveList>
  <BidList> 
    <Include> boolean </Include>
    <IncludeNotes> boolean </IncludeNotes>
    <Pagination> PaginationType
      <EntriesPerPage> int </EntriesPerPage>
      <PageNumber> int </PageNumber>
    </Pagination>
    <Sort> ItemSortTypeCodeType </Sort>
  </BidList>
  <DeletedFromSoldList> 
    <DurationInDays> int </DurationInDays>
    <Include> boolean </Include>
    <IncludeNotes> boolean </IncludeNotes>
    <Sort> ItemSortTypeCodeType </Sort>
  </DeletedFromSoldList>
  <DeletedFromUnsoldList> 
    <DurationInDays> int </DurationInDays>
    <Include> boolean </Include>
    <IncludeNotes> boolean </IncludeNotes>
    <Sort> ItemSortTypeCodeType </Sort>
  </DeletedFromUnsoldList>
  <HideVariations> boolean </HideVariations>
  <ScheduledList> 
    <Include> boolean </Include>
    <IncludeNotes> boolean </IncludeNotes>
    <Pagination> PaginationType
      <EntriesPerPage> int </EntriesPerPage>
      <PageNumber> int </PageNumber>
    </Pagination>
    <Sort> ItemSortTypeCodeType </Sort>
  </ScheduledList>
  <SellingSummary> 
    <Include> boolean </Include>
  </SellingSummary>
  <SoldList>
    <DurationInDays> int </DurationInDays>
    <Include> boolean </Include>
    <IncludeNotes> boolean </IncludeNotes>
    <OrderStatusFilter> OrderStatusFilterCodeType </OrderStatusFilter>
    <Pagination> PaginationType
      <EntriesPerPage> int </EntriesPerPage>
      <PageNumber> int </PageNumber>
    </Pagination>
    <Sort> ItemSortTypeCodeType </Sort>
  </SoldList>
  <UnsoldList> 
    <DurationInDays> int </DurationInDays>
    <Include> boolean </Include>
    <IncludeNotes> boolean </IncludeNotes>
    <Pagination> PaginationType
      <EntriesPerPage> int </EntriesPerPage>
      <PageNumber> int </PageNumber>
    </Pagination>
    <Sort> ItemSortTypeCodeType </Sort>
  </UnsoldList>
  <!-- Standard Input Fields -->
  <DetailLevel> DetailLevelCodeType </DetailLevel>
  <DetailLevel> DetailLevelCodeType </DetailLevel>
  <!-- ... more DetailLevel values allowed here ... -->
  <ErrorLanguage> string </ErrorLanguage>
  <MessageID> string </MessageID>
  <OutputSelector> string </OutputSelector>
  <OutputSelector> string </OutputSelector>
  <!-- ... more OutputSelector values allowed here ... -->
  <Version> string </Version>
  <WarningLevel> WarningLevelCodeType </WarningLevel>
</GetMyeBaySellingRequest>`
	o := NewTemplate_get_my_ebay_selling([]byte(data))
	if o.ActiveList.Include != " boolean " {
		t.Errorf("o.ActiveList.Include has not been filled out!! %+v\n",o.ActiveList)
	}
	if o.ActiveList.IncludeNotes != " boolean " {
		t.Errorf("o.ActiveList.IncludeNotes has not been filled out!! %+v\n",o.ActiveList)
	}
	if o.ActiveList.ListingType != " ListingTypeCodeType " {
		t.Errorf("o.ActiveList.ListingType has not been filled out!! %+v\n",o.ActiveList)
	}
	if o.ActiveList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.ActiveList.Sort has not been filled out!! %+v\n",o.ActiveList)
	}
	if o.BidList.Include != " boolean " {
		t.Errorf("o.BidList.Include has not been filled out!! %+v\n",o.BidList)
	}
	if o.BidList.IncludeNotes != " boolean " {
		t.Errorf("o.BidList.IncludeNotes has not been filled out!! %+v\n",o.BidList)
	}
	if o.BidList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.BidList.Sort has not been filled out!! %+v\n",o.BidList)
	}
	if o.DeletedFromSoldList.DurationInDays != " int " {
		t.Errorf("o.DeletedFromSoldList.DurationInDays has not been filled out!! %+v\n",o.DeletedFromSoldList)
	}
	if o.DeletedFromSoldList.Include != " boolean " {
		t.Errorf("o.DeletedFromSoldList.Include has not been filled out!! %+v\n",o.DeletedFromSoldList)
	}
	if o.DeletedFromSoldList.IncludeNotes != " boolean " {
		t.Errorf("o.DeletedFromSoldList.IncludeNotes has not been filled out!! %+v\n",o.DeletedFromSoldList)
	}
	if o.DeletedFromSoldList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.DeletedFromSoldList.Sort has not been filled out!! %+v\n",o.DeletedFromSoldList)
	}
	if o.DeletedFromUnsoldList.DurationInDays != " int " {
		t.Errorf("o.DeletedFromUnsoldList.DurationInDays has not been filled out!! %+v\n",o.DeletedFromUnsoldList)
	}
	if o.DeletedFromUnsoldList.Include != " boolean " {
		t.Errorf("o.DeletedFromUnsoldList.Include has not been filled out!! %+v\n",o.DeletedFromUnsoldList)
	}
	if o.DeletedFromUnsoldList.IncludeNotes != " boolean " {
		t.Errorf("o.DeletedFromUnsoldList.IncludeNotes has not been filled out!! %+v\n",o.DeletedFromUnsoldList)
	}
	if o.DeletedFromUnsoldList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.DeletedFromUnsoldList.Sort has not been filled out!! %+v\n",o.DeletedFromUnsoldList)
	}
	if o.HideVariations != " boolean " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n",o.HideVariations)
	}
	if o.ScheduledList.Include != " boolean " {
		t.Errorf("o.ScheduledList.Include has not been filled out!! %+v\n",o.ScheduledList)
	}
	if o.ScheduledList.IncludeNotes != " boolean " {
		t.Errorf("o.ScheduledList.IncludeNotes has not been filled out!! %+v\n",o.ScheduledList)
	}
	if o.ScheduledList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.ScheduledList.Sort has not been filled out!! %+v\n",o.ScheduledList)
	}
	if o.SellingSummary.Include != " boolean " {
		t.Errorf("o.SellingSummary.Include has not been filled out!! %+v\n",o.SellingSummary)
	}
	if o.SoldList.DurationInDays != " int " {
		t.Errorf("o.SoldList.DurationInDays has not been filled out!! %+v\n",o.SoldList)
	}
	if o.SoldList.Include != " boolean " {
		t.Errorf("o.SoldList.Include has not been filled out!! %+v\n",o.SoldList)
	}
	if o.SoldList.IncludeNotes != " boolean " {
		t.Errorf("o.SoldList.IncludeNotes has not been filled out!! %+v\n",o.SoldList)
	}
	if o.SoldList.OrderStatusFilter != " OrderStatusFilterCodeType " {
		t.Errorf("o.SoldList.OrderStatusFilter has not been filled out!! %+v\n",o.SoldList)
	}
	if o.SoldList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.SoldList.Sort has not been filled out!! %+v\n",o.SoldList)
	}
	if o.UnsoldList.DurationInDays != " int " {
		t.Errorf("o.UnsoldList.DurationInDays has not been filled out!! %+v\n",o.UnsoldList)
	}
	if o.UnsoldList.Include != " boolean " {
		t.Errorf("o.UnsoldList.Include has not been filled out!! %+v\n",o.UnsoldList)
	}
	if o.UnsoldList.IncludeNotes != " boolean " {
		t.Errorf("o.UnsoldList.IncludeNotes has not been filled out!! %+v\n",o.UnsoldList)
	}
	if o.UnsoldList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.UnsoldList.Sort has not been filled out!! %+v\n",o.UnsoldList)
	}
	if o.DetailLevel[0] != " DetailLevelCodeType " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n",o.DetailLevel)
	}
	if o.DetailLevel[1] != " DetailLevelCodeType " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n",o.DetailLevel)
	}
	if o.ErrorLanguage != " string " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n",o.ErrorLanguage)
	}
	if o.MessageID != " string " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n",o.MessageID)
	}
	if o.OutputSelector[0] != " string " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n",o.OutputSelector)
	}
	if o.OutputSelector[1] != " string " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n",o.OutputSelector)
	}
	if o.Version != " string " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n",o.Version)
	}
	if o.WarningLevel != " WarningLevelCodeType " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n",o.WarningLevel)
	}
}
