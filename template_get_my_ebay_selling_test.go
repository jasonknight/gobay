package gobay

import "testing"
import "fmt"

func TestGetMyeBaySellingTemplate(t *testing.T) {
	data := `<?xml version="1.0" encoding="utf-8"?>
<GetMyeBaySellingRequest xmlns="urn:ebay:apis:eBLBaseComponents">
  <!-- Call-specific Input Fields -->
  <ActiveList> 
    <Include>true</Include>
    <IncludeNotes>true</IncludeNotes>
    <ListingType> ListingTypeCodeType </ListingType>
    <Pagination> PaginationType
      <EntriesPerPage>666</EntriesPerPage>
      <PageNumber>666</PageNumber>
    </Pagination>
    <Sort> ItemSortTypeCodeType </Sort>
  </ActiveList>
  <BidList> 
    <Include>true</Include>
    <IncludeNotes>true</IncludeNotes>
    <Pagination> PaginationType
      <EntriesPerPage>666</EntriesPerPage>
      <PageNumber>666</PageNumber>
    </Pagination>
    <Sort> ItemSortTypeCodeType </Sort>
  </BidList>
  <DeletedFromSoldList> 
    <DurationInDays>666</DurationInDays>
    <Include>true</Include>
    <IncludeNotes>true</IncludeNotes>
    <Sort> ItemSortTypeCodeType </Sort>
  </DeletedFromSoldList>
  <DeletedFromUnsoldList> 
    <DurationInDays>666</DurationInDays>
    <Include>true</Include>
    <IncludeNotes>true</IncludeNotes>
    <Sort> ItemSortTypeCodeType </Sort>
  </DeletedFromUnsoldList>
  <HideVariations>true</HideVariations>
  <ScheduledList> 
    <Include>true</Include>
    <IncludeNotes>true</IncludeNotes>
    <Pagination> PaginationType
      <EntriesPerPage>666</EntriesPerPage>
      <PageNumber>666</PageNumber>
    </Pagination>
    <Sort> ItemSortTypeCodeType </Sort>
  </ScheduledList>
  <SellingSummary> 
    <Include>true</Include>
  </SellingSummary>
  <SoldList>
    <DurationInDays>666</DurationInDays>
    <Include>true</Include>
    <IncludeNotes>true</IncludeNotes>
    <OrderStatusFilter> OrderStatusFilterCodeType </OrderStatusFilter>
    <Pagination> PaginationType
      <EntriesPerPage>666</EntriesPerPage>
      <PageNumber>666</PageNumber>
    </Pagination>
    <Sort> ItemSortTypeCodeType </Sort>
  </SoldList>
  <UnsoldList> 
    <DurationInDays>666</DurationInDays>
    <Include>true</Include>
    <IncludeNotes>true</IncludeNotes>
    <Pagination> PaginationType
      <EntriesPerPage>666</EntriesPerPage>
      <PageNumber>666</PageNumber>
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
	o := NewMyeBaySelling()
	err := o.FromXML([]byte(data))
	if err != nil {
		t.Errorf("failed with %v", err)
		return
	}
	if o.ActiveList.Include != true {
		t.Errorf("o.ActiveList.Include has not been filled out!! %+v\n", o.ActiveList)
		return
	}
	if o.ActiveList.IncludeNotes != true {
		t.Errorf("o.ActiveList.IncludeNotes has not been filled out!! %+v\n", o.ActiveList)
		return
	}
	if o.ActiveList.ListingType != " ListingTypeCodeType " {
		t.Errorf("o.ActiveList.ListingType has not been filled out!! %+v\n", o.ActiveList)
		return
	}
	if o.ActiveList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.ActiveList.Sort has not been filled out!! %+v\n", o.ActiveList)
		return
	}
	if o.BidList.Include != true {
		t.Errorf("o.BidList.Include has not been filled out!! %+v\n", o.BidList)
		return
	}
	if o.BidList.IncludeNotes != true {
		t.Errorf("o.BidList.IncludeNotes has not been filled out!! %+v\n", o.BidList)
		return
	}
	if o.BidList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.BidList.Sort has not been filled out!! %+v\n", o.BidList)
		return
	}
	if o.DeletedFromSoldList.DurationInDays != 666 {
		t.Errorf("o.DeletedFromSoldList.DurationInDays has not been filled out!! %+v\n", o.DeletedFromSoldList)
		return
	}
	if o.DeletedFromSoldList.Include != true {
		t.Errorf("o.DeletedFromSoldList.Include has not been filled out!! %+v\n", o.DeletedFromSoldList)
		return
	}
	if o.DeletedFromSoldList.IncludeNotes != true {
		t.Errorf("o.DeletedFromSoldList.IncludeNotes has not been filled out!! %+v\n", o.DeletedFromSoldList)
		return
	}
	if o.DeletedFromSoldList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.DeletedFromSoldList.Sort has not been filled out!! %+v\n", o.DeletedFromSoldList)
		return
	}
	if o.DeletedFromUnsoldList.DurationInDays != 666 {
		t.Errorf("o.DeletedFromUnsoldList.DurationInDays has not been filled out!! %+v\n", o.DeletedFromUnsoldList)
		return
	}
	if o.DeletedFromUnsoldList.Include != true {
		t.Errorf("o.DeletedFromUnsoldList.Include has not been filled out!! %+v\n", o.DeletedFromUnsoldList)
		return
	}
	if o.DeletedFromUnsoldList.IncludeNotes != true {
		t.Errorf("o.DeletedFromUnsoldList.IncludeNotes has not been filled out!! %+v\n", o.DeletedFromUnsoldList)
		return
	}
	if o.DeletedFromUnsoldList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.DeletedFromUnsoldList.Sort has not been filled out!! %+v\n", o.DeletedFromUnsoldList)
		return
	}
	if o.HideVariations != true {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n", o.HideVariations)
		return
	}
	if o.ScheduledList.Include != true {
		t.Errorf("o.ScheduledList.Include has not been filled out!! %+v\n", o.ScheduledList)
		return
	}
	if o.ScheduledList.IncludeNotes != true {
		t.Errorf("o.ScheduledList.IncludeNotes has not been filled out!! %+v\n", o.ScheduledList)
		return
	}
	if o.ScheduledList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.ScheduledList.Sort has not been filled out!! %+v\n", o.ScheduledList)
		return
	}
	if o.SellingSummary.Include != true {
		t.Errorf("o.SellingSummary.Include has not been filled out!! %+v\n", o.SellingSummary)
		return
	}
	if o.SoldList.DurationInDays != 666 {
		t.Errorf("o.SoldList.DurationInDays has not been filled out!! %+v\n", o.SoldList)
		return
	}
	if o.SoldList.Include != true {
		t.Errorf("o.SoldList.Include has not been filled out!! %+v\n", o.SoldList)
		return
	}
	if o.SoldList.IncludeNotes != true {
		t.Errorf("o.SoldList.IncludeNotes has not been filled out!! %+v\n", o.SoldList)
		return
	}
	if o.SoldList.OrderStatusFilter != " OrderStatusFilterCodeType " {
		t.Errorf("o.SoldList.OrderStatusFilter has not been filled out!! %+v\n", o.SoldList)
		return
	}
	if o.SoldList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.SoldList.Sort has not been filled out!! %+v\n", o.SoldList)
		return
	}
	if o.UnsoldList.DurationInDays != 666 {
		t.Errorf("o.UnsoldList.DurationInDays has not been filled out!! %+v\n", o.UnsoldList)
		return
	}
	if o.UnsoldList.Include != true {
		t.Errorf("o.UnsoldList.Include has not been filled out!! %+v\n", o.UnsoldList)
		return
	}
	if o.UnsoldList.IncludeNotes != true {
		t.Errorf("o.UnsoldList.IncludeNotes has not been filled out!! %+v\n", o.UnsoldList)
		return
	}
	if o.UnsoldList.Sort != " ItemSortTypeCodeType " {
		t.Errorf("o.UnsoldList.Sort has not been filled out!! %+v\n", o.UnsoldList)
		return
	}
	if o.DetailLevels[0] != " DetailLevelCodeType " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n", o.DetailLevels)
		return
	}
	if o.DetailLevels[1] != " DetailLevelCodeType " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n", o.DetailLevels)
		return
	}
	if o.OutputSelectors[0] != " string " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n", o.OutputSelectors)
		return
	}
	if o.OutputSelectors[1] != " string " {
		t.Errorf("Template_get_my_ebay_selling has not been filled out %+v!!\n", o.OutputSelectors)
		return
	}
}

func TestLists(t *testing.T) {
	o := NewMyeBaySelling()

	o.AddActiveList()

	if o.ActiveList.Include != true {
		t.Errorf("Failed, including ActiveList is required %+v", o.ActiveList)
		return
	}

	txt, err := compileGoString("Test", GetMyeBaySellingTemplate(), o, nil)
	txt = fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><root>%s</root>", txt)

	to := NewMyeBaySelling()
	err = to.FromXML([]byte(txt))
	if err != nil {
		t.Errorf("failed %v", err)
	}

	if to.ActiveList.Include != o.ActiveList.Include {
		t.Errorf("Failed to reload output for AddActiveList")
		return
	}
	// Generated Code here
	o.AddSoldList()

	if o.SoldList.Include != true {
		t.Errorf("Failed, including SoldList is required %+v", o.SoldList)
		return
	}

	txt, err = compileGoString("Test", GetMyeBaySellingTemplate(), o, nil)
	if err != nil {
		t.Errorf("Failed: %v", err)
	}
	txt = fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><root>%s</root>", txt)

	to = NewMyeBaySelling()
	to.FromXML([]byte(txt))

	if to.SoldList.Include != o.SoldList.Include {
		t.Errorf("Failed to reload output for AddSoldList")
		return
	}

	o.AddUnsoldList()

	if o.UnsoldList.Include != true {
		t.Errorf("Failed, including UnsoldList is required %+v", o.UnsoldList)
		return
	}

	txt, err = compileGoString("Test", GetMyeBaySellingTemplate(), o, nil)
	if err != nil {
		t.Errorf("Failed: %v", err)
	}
	txt = fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><root>%s</root>", txt)

	to = NewMyeBaySelling()
	err = to.FromXML([]byte(txt))
	if err != nil {
		t.Errorf("failed %v", err)
	}

	if to.UnsoldList.Include != o.UnsoldList.Include {
		t.Errorf("Failed to reload output for AddUnsoldList")
		return
	}

	o.AddBidList()

	if o.BidList.Include != true {
		t.Errorf("Failed, including BidList is required %+v", o.BidList)
		return
	}

	txt, err = compileGoString("Test", GetMyeBaySellingTemplate(), o, nil)
	if err != nil {
		t.Errorf("Failed: %v", err)
	}
	txt = fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><root>%s</root>", txt)

	to = NewMyeBaySelling()
	err = to.FromXML([]byte(txt))
	if err != nil {
		t.Errorf("failed %v", err)
	}

	if to.BidList.Include != o.BidList.Include {
		t.Errorf("Failed to reload output for AddBidList")
		return
	}

	o.AddDeletedFromSoldList()

	if o.DeletedFromSoldList.Include != true {
		t.Errorf("Failed, including DeletedFromSoldList is required %+v", o.DeletedFromSoldList)
		return
	}

	txt, err = compileGoString("Test", GetMyeBaySellingTemplate(), o, nil)
	if err != nil {
		t.Errorf("Failed: %v", err)
	}
	txt = fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><root>%s</root>", txt)

	to = NewMyeBaySelling()
	err = to.FromXML([]byte(txt))
	if err != nil {
		t.Errorf("failed %v", err)
	}

	if to.DeletedFromSoldList.Include != o.DeletedFromSoldList.Include {
		t.Errorf("Failed to reload output for AddDeletedFromSoldList")
		return
	}

	o.AddDeletedFromUnsoldList()

	if o.DeletedFromUnsoldList.Include != true {
		t.Errorf("Failed, including DeletedFromUnsoldList is required %+v", o.DeletedFromUnsoldList)
		return
	}

	txt, err = compileGoString("Test", GetMyeBaySellingTemplate(), o, nil)
	if err != nil {
		t.Errorf("Failed: %v", err)
	}
	txt = fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><root>%s</root>", txt)

	to = NewMyeBaySelling()
	err = to.FromXML([]byte(txt))
	if err != nil {
		t.Errorf("failed %v", err)
	}

	if to.DeletedFromUnsoldList.Include != o.DeletedFromUnsoldList.Include {
		t.Errorf("Failed to reload output for AddDeletedFromUnsoldList")
		return
	}

	o.AddScheduledList()

	if o.ScheduledList.Include != true {
		t.Errorf("Failed, including ScheduledList is required %+v", o.ScheduledList)
		return
	}

	txt, err = compileGoString("Test", GetMyeBaySellingTemplate(), o, nil)
	if err != nil {
		t.Errorf("Failed: %v", err)
	}
	txt = fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><root>%s</root>", txt)

	to = NewMyeBaySelling()
	err = to.FromXML([]byte(txt))
	if err != nil {
		t.Errorf("failed %v", err)
	}

	if to.ScheduledList.Include != o.ScheduledList.Include {
		t.Errorf("Failed to reload output for AddScheduledList")
		return
	}

}
