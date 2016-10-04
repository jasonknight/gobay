package gobay

import "encoding/xml"

type MyeBaySellingOpts struct {
	Include           bool
	IncludeNotes      bool
	ListingType       string
	OrderStatusFilter string
	Pagination        struct {
		EntriesPerPage int
		PageNumber     int
	}

	Sort           string
	DurationInDays int
}
type MyeBaySellingStruct struct {
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

func GetMyeBaySellingStructFromXML(data []byte) *MyeBaySellingStruct {
	var o MyeBaySellingStruct
	xml.Unmarshal(data, &o)
	return &o
}

func GetMyeBaySellingTemplate() string {
	return `
  {{ if .ActiveList }}
	<ActiveList> 
		<Include>{{ .Include }}</Include>
		<IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
		<ListingType> {{ .ListingType }}</ListingType>
		{{ if .Pagination }}
			<Pagination> 
			  <EntriesPerPage>{{ .EntriesPerPage }}</EntriesPerPage>
			  <PageNumber>{{ .PageNumber }}</PageNumber>
			</Pagination>
		{{ end }}
		<Sort>{{ .Sort }}</Sort>
	</ActiveList>
  {{ end }}
  {{ if .BidList }}
	<BidList> 
		<Include>{{ .Include }}</Include>
		<IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
		<ListingType> {{ .ListingType }}</ListingType>
		{{ if .Pagination }}
			<Pagination> 
			  <EntriesPerPage>{{ .EntriesPerPage }}</EntriesPerPage>
			  <PageNumber>{{ .PageNumber }}</PageNumber>
			</Pagination>
		{{ end }}
		<Sort>{{ .Sort }}</Sort>
	</BidList>
  {{ end }}
  {{ if .DeletedFromSoldList }}
	  <DeletedFromSoldList> 
	    <DurationInDays> int </DurationInDays>
	    <Include> boolean </Include>
	    <IncludeNotes> boolean </IncludeNotes>
	    <Sort> ItemSortTypeCodeType </Sort>
	  </DeletedFromSoldList>
  {{ end }}
  {{ if .DeletedFromUnsoldList }}
	  <DeletedFromUnsoldList> 
	    <DurationInDays> int </DurationInDays>
	    <Include> boolean </Include>
	    <IncludeNotes> boolean </IncludeNotes>
	    <Sort> ItemSortTypeCodeType </Sort>
	  </DeletedFromUnsoldList>
  {{ end }}
  {{ if .ScheduledList }}
  <ScheduledList> 
    <Include>{{ .Include }}</Include>
	<IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
	<ListingType> {{ .ListingType }}</ListingType>
	{{ if .Pagination }}
		<Pagination> 
		  <EntriesPerPage>{{ .EntriesPerPage }}</EntriesPerPage>
		  <PageNumber>{{ .PageNumber }}</PageNumber>
		</Pagination>
	{{ end }}
	<Sort>{{ .Sort }}</Sort>
  </ScheduledList>
  {{ end }}
  {{ if .SoldList }}
  <SoldList>
    <DurationInDays>{{ .DurationInDays }}</DurationInDays>
    <Include>{{ .Include }}</Include>
	<IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
	<ListingType> {{ .ListingType }}</ListingType>
	{{ if .Pagination }}
		<Pagination> 
		  <EntriesPerPage>{{ .EntriesPerPage }}</EntriesPerPage>
		  <PageNumber>{{ .PageNumber }}</PageNumber>
		</Pagination>
	{{ end }}
	<Sort>{{ .Sort }}</Sort>
  </SoldList>
  {{ end }}
  {{ if .UnsoldList }}
  <UnsoldList> 
    <DurationInDays>{{ .DurationInDays }}</DurationInDays>
    <Include>{{ .Include }}</Include>
	<IncludeNotes>{{ .IncludeNotes }}</IncludeNotes>
	<ListingType> {{ .ListingType }}</ListingType>
	{{ if .Pagination }}
		<Pagination> 
		  <EntriesPerPage>{{ .EntriesPerPage }}</EntriesPerPage>
		  <PageNumber>{{ .PageNumber }}</PageNumber>
		</Pagination>
	{{ end }}
	<Sort>{{ .Sort }}</Sort>
  </UnsoldList>
  {{ end }}
  <HideVariations>{{ .HideVariations }}</HideVariations>
  
  {{ if .SellingSummary }}
	  <SellingSummary> 
	    <Include>{{ .Include }}</Include>
	  </SellingSummary>
  {{ end }}
  {{ range $x,$y :=  .DetailLevels }}
  <DetailLevel>{{ $y }}</DetailLevel>
  {{ end }}
  {{ range $x,$y :=  .OutputSelectors }}
  <OutputSelector>{{ $y }}</OutputSelector>
  {{ end }}`
}
