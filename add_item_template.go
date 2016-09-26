package gobay

func GetItemTemplate() string {
    return `
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
    <Storefront> 
      <StoreCategoryID>{{ .StoreCategoryID }}</StoreCategoryID> 
    </Storefront> 
    <Title><![CDATA[{{ .Title }}]]></Title>
    <PictureDetails> 
      <GalleryType>Gallery</GalleryType>
      {{ range .PictureDetails }}
            <PictureURL>{{ .PictureURL }}</PictureURL>
      {{ end }}
    </PictureDetails>
      <SKU>{{ .SKU }}</SKU>
    {{ range $key, $value := .ShipToLocations }} 
        <ShipToLocations>{{ $value }}</ShipToLocations> 
    {{ end }}
    <ShippingDetails>
    {{ range $key, $value := .ShippingServiceOptions }}
      <ShippingServiceOptions>
        <ShippingService>{{ $value.ShippingService }}</ShippingService>
        <ShippingServiceCost currencyID="{{ .Currency }}">{{ $value.ShippingServiceCost }}</ShippingServiceCost>
        <ShippingServicePriority>{{ $value.ShippingServicePriority }}</ShippingServicePriority>
      </ShippingServiceOptions>
    {{ end }}
    {{ range $key, $value := .InternationalShippingServiceOptions }}
        <InternationalShippingServiceOption>
          <ShippingService>{{ $value.ShippingService }}</ShippingService>
          <ShippingServiceCost currencyID="{{ .Currency }}">{{ $value.ShippingServiceCost }}</ShippingServiceCost>
          <ShippingServicePriority>{{ $value.ShippingServicePriority }}</ShippingServicePriority>
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