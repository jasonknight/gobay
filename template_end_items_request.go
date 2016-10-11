package gobay
func EndItemsTemplate() string {
	return`  <!-- Call-specific Input Fields -->
  <EndItemRequestContainer> EndItemRequestContainerType
    <EndingReason> EndReasonCodeType </EndingReason>
    <ItemID> ItemIDType (string) </ItemID>
    <MessageID> string </MessageID>
    <SellerInventoryID> string </SellerInventoryID>
  </EndItemRequestContainer>
  <EndItemRequestContainer> EndItemRequestContainerType
    <EndingReason> EndReasonCodeType </EndingReason>
    <ItemID> ItemIDType (string) </ItemID>
    <MessageID> string </MessageID>
    <SellerInventoryID> string </SellerInventoryID>
  </EndItemRequestContainer>
  
`
}
