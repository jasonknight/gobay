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
      </ShippingServiceOptions>
    {{ end }}
    {{ range $key, $value := .InternationalShippingServiceOptions }}
        <InternationalShippingServiceOption>
          <ShippingService>{{ $value.Service }}</ShippingService>
          <ShippingServiceCost currencyID="{{ $Currency }}">{{ $value.Cost }}</ShippingServiceCost>
          <ShippingServicePriority>{{ $value.Priority }}</ShippingServicePriority>
          <ShipToLocation>{{ $value.Location }}</ShipToLocation>
        </InternationalShippingServiceOption>
    {{ end }}
    </ShippingDetails>
    <DispatchTimeMax>{{ .DispatchTimeMax }}</DispatchTimeMax>
    <ReturnPolicy>
      <ReturnsAcceptedOption>ReturnsAccepted</ReturnsAcceptedOption>
      <RefundOption>MoneyBack</RefundOption>
      <ReturnsWithinOption>Days_30</ReturnsWithinOption>
      <Description>You may return items still in pristine, un-opened condition.</Description>
      <ShippingCostPaidByOption>Buyer</ShippingCostPaidByOption>
    </ReturnPolicy>
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
