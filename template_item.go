package gobay

func ItemTemplate() string {
	return `
  {{ $Currency := .Currency }}
  <Item> 
    <CategoryMappingAllowed>true</CategoryMappingAllowed> 
    <Country>{{ .Country }}</Country> 
    <Currency>{{ .Currency }}</Currency> 
    <Description><![CDATA[{{ .Description }}]]></Description> 
    <ListingDuration>{{ .ListingDuration }}</ListingDuration> 
    <ListingType>{{ .ListingType }}</ListingType> 
    <Location>{{ .Location }}</Location> 
    {{ range $key,$value := .PaymentMethods }}
      <PaymentMethods>{{ $value }}</PaymentMethods> 
    {{ end }}
    <PayPalEmailAddress>{{ .PayPalEmailAddress }}</PayPalEmailAddress> 
    <PrimaryCategory> 
      <CategoryID>{{ .PrimaryCategory }}</CategoryID> 
    </PrimaryCategory> 
    <Quantity>{{ .Quantity }}</Quantity> 
    <Site>{{ .Site }}</Site> 
    <StartPrice currencyID="{{ .Currency }}">{{ .StartPrice }}</StartPrice> 
    {{ if .StoreCategoryID }}
    <Storefront> 
      <StoreCategoryID>{{ .StoreCategoryID }}</StoreCategoryID> 
    </Storefront> 
    {{ end }}
    <Title><![CDATA[{{ .Title }}]]></Title>
    {{ range .PictureDetails }}
    <PictureDetails> 
            {{ if .GalleryType }}
            <GalleryType>{{.GalleryType}}</GalleryType>
            {{ end }}
            {{ if .GalleryURL }}
            <GalleryURL>{{ .GalleryURL }}</GalleryURL>
            {{ end }}
            {{ if .PhotoDisplay }}
            <GPhotoDisplay>{{ .PhotoDisplay }}</PhotoDisplay>
            {{ end }}
            <PictureURL>{{ .PictureURL }}</PictureURL>
    </PictureDetails>
    {{ end }}
      <SKU>{{ .SKU }}</SKU>
    {{ range $key, $value := .ShipToLocations }} 
        <ShipToLocations>{{ $value }}</ShipToLocations> 
    {{ end }}
    <ShippingDetails>
    {{ range $key, $value := .ShippingServiceOptions }}
      <ShippingServiceOptions>
        <ShippingService>{{ $value.Service }}</ShippingService>
        <ShippingServiceCost currencyID="{{ $Currency }}">{{ $value.Cost }}</ShippingServiceCost>
        <ShippingServicePriority>{{ $value.Priority }}</ShippingServicePriority>
        {{ if $value.Locations }}
          {{ range $x,$y := $value.Locations }}
            <ShipToLocation>{{ $y }}</ShipToLocation>
          {{ end }}
        {{ end }}
      </ShippingServiceOptions>
    {{ end }}
    {{ range $key, $value := .InternationalShippingServiceOptions }}
        <InternationalShippingServiceOption>
          <ShippingService>{{ $value.Service }}</ShippingService>
          <ShippingServiceCost currencyID="{{ $Currency }}">{{ $value.Cost }}</ShippingServiceCost>
          <ShippingServicePriority>{{ $value.Priority }}</ShippingServicePriority>
          {{ if $value.Locations }}
            {{ range $x,$y := $value.Locations }}
              <ShipToLocation>{{ $y }}</ShipToLocation>
            {{ end }}
          {{ end }}
        </InternationalShippingServiceOption>
    {{ end }}
    </ShippingDetails>
    <DispatchTimeMax>{{ .DispatchTimeMax }}</DispatchTimeMax>
    {{ if .ReturnPolicy }} 

    	{{ if .ReturnPolicy.Description }}
    		<Description>{{ .ReturnPolicy.Description }}</Description>
    	{{ end }}
        	{{ if .ReturnPolicy.EAN }}
    		<EAN>{{ .ReturnPolicy.EAN }}</EAN>
    	{{ end }}
        	{{ if .ReturnPolicy.ExtendedHolidayReturns }}
    		<ExtendedHolidayReturns>{{ .ReturnPolicy.ExtendedHolidayReturns }}</ExtendedHolidayReturns>
    	{{ end }}
        	{{ if .ReturnPolicy.Refund }}
    		<Refund>{{ .ReturnPolicy.Refund }}</Refund>
    	{{ end }}
        	{{ if .ReturnPolicy.RefundOption }}
    		<RefundOption>{{ .ReturnPolicy.RefundOption }}</RefundOption>
    	{{ end }}
        	{{ if .ReturnPolicy.RestockingFeeValue }}
    		<RestockingFeeValue>{{ .ReturnPolicy.RestockingFeeValue }}</RestockingFeeValue>
    	{{ end }}
        	{{ if .ReturnPolicy.RestockingFeeValueOption }}
    		<RestockingFeeValueOption>{{ .ReturnPolicy.RestockingFeeValueOption }}</RestockingFeeValueOption>
    	{{ end }}
        	{{ if .ReturnPolicy.ReturnsAccepted }}
    		<ReturnsAccepted>{{ .ReturnPolicy.ReturnsAccepted }}</ReturnsAccepted>
    	{{ end }}
        	{{ if .ReturnPolicy.ReturnsAcceptedOption }}
    		<ReturnsAcceptedOption>{{ .ReturnPolicy.ReturnsAcceptedOption }}</ReturnsAcceptedOption>
    	{{ end }}
        	{{ if .ReturnPolicy.ReturnsWithin }}
    		<ReturnsWithin>{{ .ReturnPolicy.ReturnsWithin }}</ReturnsWithin>
    	{{ end }}
        	{{ if .ReturnPolicy.ReturnsWithinOption }}
    		<ReturnsWithinOption>{{ .ReturnPolicy.ReturnsWithinOption }}</ReturnsWithinOption>
    	{{ end }}
        	{{ if .ReturnPolicy.ShippingCostPaidBy }}
    		<ShippingCostPaidBy>{{ .ReturnPolicy.ShippingCostPaidBy }}</ShippingCostPaidBy>
    	{{ end }}
        	{{ if .ReturnPolicy.ShippingCostPaidByOption }}
    		<ShippingCostPaidByOption>{{ .ReturnPolicy.ShippingCostPaidByOption }}</ShippingCostPaidByOption>
    	{{ end }}
        	{{ if .ReturnPolicy.WarrantyDuration }}
    		<WarrantyDuration>{{ .ReturnPolicy.WarrantyDuration }}</WarrantyDuration>
    	{{ end }}
        	{{ if .ReturnPolicy.WarrantyDurationOption }}
    		<WarrantyDurationOption>{{ .ReturnPolicy.WarrantyDurationOption }}</WarrantyDurationOption>
    	{{ end }}
        	{{ if .ReturnPolicy.WarrantyOffered }}
    		<WarrantyOffered>{{ .ReturnPolicy.WarrantyOffered }}</WarrantyOffered>
    	{{ end }}
        	{{ if .ReturnPolicy.WarrantyOfferedOption }}
    		<WarrantyOfferedOption>{{ .ReturnPolicy.WarrantyOfferedOption }}</WarrantyOfferedOption>
    	{{ end }}
        	{{ if .ReturnPolicy.WarrantyType }}
    		<WarrantyType>{{ .ReturnPolicy.WarrantyType }}</WarrantyType>
    	{{ end }}
      {{ if .ReturnPolicy.WarrantyTypeOption }}
        <WarrantyTypeOption>{{ .ReturnPolicy.WarrantyTypeOption }}</WarrantyTypeOption>
      {{ end }}  


    {{ end }}
  </Item> 
    `
}

type AddItemsChild struct {
	Item      *Item
	MessageID string
	Text      string
}
type AddItemsStruct struct {
	Children []AddItemsChild
}

func AddItemsTemplate() string {
	return `
{{ range .Children }}
  <AddItemRequestContainer>
    <MessageID>{{ .MessageID }}</MessageID>
    {{ .Text }}
  </AddItemRequestContainer>
{{ end}}
`
}
