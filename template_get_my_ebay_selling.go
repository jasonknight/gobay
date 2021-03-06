package gobay

import "encoding/xml"

type PaginationOpts struct {
	EntriesPerPage int
	PageNumber     int
}
type MyeBaySellingOpts struct {
	Include           bool
	IncludeNotes      bool
	ListingType       string
	OrderStatusFilter string
	Pagination        PaginationOpts

	Sort           string
	DurationInDays int
}
type MyeBaySelling struct {
	ActiveList            MyeBaySellingOpts
	SoldList              MyeBaySellingOpts
	UnsoldList            MyeBaySellingOpts
	BidList               MyeBaySellingOpts
	DeletedFromSoldList   MyeBaySellingOpts
	DeletedFromUnsoldList MyeBaySellingOpts
	ScheduledList         MyeBaySellingOpts
	DetailLevels          []string `xml:"DetailLevel"`
	OutputSelectors       []string `xml:"OutputSelector"`
	SellingSummary        struct {
		Include bool
	}
	HideVariations bool
}

func (o *MyeBaySelling) FromXML(data []byte) error {
	err := xml.Unmarshal(data, o)
	return err
}

func (o *MyeBaySelling) AddActiveList() {
	o.ActiveList = MyeBaySellingOpts{
		Include:      true,
		IncludeNotes: false,
		ListingType:  "Auction",
		Pagination:   PaginationOpts{EntriesPerPage: 25, PageNumber: 1},
		Sort:         "EndTimeDescending",
	}

}
func (o *MyeBaySelling) AddScheduledList() {
	o.ScheduledList = MyeBaySellingOpts{
		Include:      true,
		IncludeNotes: false,
		ListingType:  "Auction",
		Pagination:   PaginationOpts{EntriesPerPage: 25, PageNumber: 1},
		Sort:         "EndTimeDescending",
	}
}
func (o *MyeBaySelling) AddBidList() {
	o.BidList = MyeBaySellingOpts{
		Include:      true,
		IncludeNotes: false,
		Pagination:   PaginationOpts{EntriesPerPage: 25, PageNumber: 1},
		Sort:         "BestOffer",
	}
}
func (o *MyeBaySelling) AddDeletedFromSoldList() {
	o.DeletedFromSoldList = MyeBaySellingOpts{
		DurationInDays: 10,
		Include:        true,
		IncludeNotes:   false,
		Sort:           "ItemIDDescending",
	}
}
func (o *MyeBaySelling) AddDeletedFromUnsoldList() {
	o.DeletedFromUnsoldList = MyeBaySellingOpts{
		DurationInDays: 10,
		Include:        true,
		IncludeNotes:   false,
		Sort:           "ItemIDDescending",
	}
}
func (o *MyeBaySelling) AddSoldList() {
	o.SoldList = MyeBaySellingOpts{
		DurationInDays: 10,
		Include:        true,
		IncludeNotes:   false,
		Pagination:     PaginationOpts{EntriesPerPage: 25, PageNumber: 1},
		Sort:           "ItemIDDescending",
	}
}
func (o *MyeBaySelling) AddUnsoldList() {
	o.UnsoldList = MyeBaySellingOpts{
		DurationInDays: 10,
		Include:        true,
		IncludeNotes:   false,
		Pagination:     PaginationOpts{EntriesPerPage: 25, PageNumber: 1},
		Sort:           "ItemIDDescending",
	}
}

func GetMyeBaySellingTemplate() string {
	return `
  {{ if .ActiveList.Include }}
  {{ with .ActiveList }}
	<ActiveList> 
		<Include>{{ .Include }}</Include>
		<IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
		<ListingType>{{ .ListingType }}</ListingType>
		{{ with .Pagination}}
			<Pagination> 
			  <EntriesPerPage>{{ .EntriesPerPage }}</EntriesPerPage>
			  <PageNumber>{{ .PageNumber }}</PageNumber>
			</Pagination>
		{{ end }}
		<Sort>{{ .Sort }}</Sort>
	</ActiveList>
  {{ end }}
  {{ end }}
  {{ if .BidList.Include }}
  {{ with .BidList }}
	<BidList> 
		<Include>{{ .Include }}</Include>
		<IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
		<ListingType>{{ .ListingType }}</ListingType>
		{{ with .Pagination}}
			<Pagination> 
			  <EntriesPerPage>{{ .EntriesPerPage }}</EntriesPerPage>
			  <PageNumber>{{ .PageNumber }}</PageNumber>
			</Pagination>
		{{ end }}
		<Sort>{{ .Sort }}</Sort>
	</BidList>
  {{ end }}
  {{ end }}
  {{ if .DeletedFromSoldList.Include }}
  {{ with .DeletedFromSoldList }}
	  <DeletedFromSoldList> 
	    <DurationInDays>{{ .DurationInDays }}</DurationInDays>
	    <Include>{{ .Include }}</Include>
	    <IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
	    <Sort>{{ .Sort }}</Sort>
	  </DeletedFromSoldList>
  {{ end }}
  {{ end }}
  {{ if .DeletedFromUnsoldList.Include }}
  {{ with .DeletedFromUnsoldList }}
	  <DeletedFromUnsoldList> 
	    <DurationInDays>{{ .DurationInDays }}</DurationInDays>
	    <Include>{{ .Include }}</Include>
	    <IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
	    <Sort>{{ .Sort }}</Sort>
	  </DeletedFromUnsoldList>
  {{ end }}
  {{ end }}
  {{ if .ScheduledList.Include }}
  {{ with .ScheduledList }}
  <ScheduledList> 
    <Include>{{ .Include }}</Include>
	<IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
	<ListingType>{{ .ListingType }}</ListingType>
	{{ with .Pagination}}
		<Pagination> 
		  <EntriesPerPage>{{ .EntriesPerPage }}</EntriesPerPage>
		  <PageNumber>{{ .PageNumber }}</PageNumber>
		</Pagination>
	{{ end }}
	<Sort>{{ .Sort }}</Sort>
  </ScheduledList>
  {{ end }}
  {{ end }}
  {{ if .SoldList.Include }}
  {{ with .SoldList }}
  <SoldList>
    <DurationInDays>{{ .DurationInDays }}</DurationInDays>
    <Include>{{ .Include }}</Include>
	<IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
	<ListingType>{{ .ListingType }}</ListingType>
	{{ with .Pagination}}
		<Pagination> 
		  <EntriesPerPage>{{ .EntriesPerPage }}</EntriesPerPage>
		  <PageNumber>{{ .PageNumber }}</PageNumber>
		</Pagination>
	{{ end }}
	<Sort>{{ .Sort }}</Sort>
  </SoldList>
  {{ end }}
  {{ end }}
  {{ if .UnsoldList.Include }}
  {{ with .UnsoldList }}
  <UnsoldList> 
    <DurationInDays>{{ .DurationInDays }}</DurationInDays>
    <Include>{{ .Include }}</Include>
	<IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
	<ListingType>{{ .ListingType }}</ListingType>
	{{ with .Pagination}}
		<Pagination> 
		  <EntriesPerPage>{{ .EntriesPerPage }}</EntriesPerPage>
		  <PageNumber>{{ .PageNumber }}</PageNumber>
		</Pagination>
	{{ end }}
	<Sort>{{ .Sort }}</Sort>
  </UnsoldList>
  {{ end }}
  {{ end }}
  <HideVariations>{{ .HideVariations }}</HideVariations>
  
  {{ if .SellingSummary.Include }}
  {{ with .SellingSummary }}
	  <SellingSummary> 
	    <Include>{{ .Include }}</Include>
	  </SellingSummary>
  {{ end }}
  {{ end }}
  {{ range $x,$y :=  .DetailLevels }}
  <DetailLevel>{{ $y }}</DetailLevel>
  {{ end }}
  {{ range $x,$y :=  .OutputSelectors }}
  <OutputSelector>{{ $y }}</OutputSelector>
  {{ end }}`
}
