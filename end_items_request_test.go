package gobay
import "testing"

        func TestEndItemsRequest ( t *testing.T ) {
            data := `<?xml version="1.0" encoding="utf-8"?>
<EndItemsRequest xmlns="urn:ebay:apis:eBLBaseComponents">
  <!-- Call-specific Input Fields -->
  <EndItemRequestContainer>EndItemRequestContainerType<EndingReason>EndReasonCodeType</EndingReason>
    <ItemID>ItemIDType (string)</ItemID>
    <MessageID>string</MessageID>
    <SellerInventoryID>string</SellerInventoryID>
  </EndItemRequestContainer>
  <EndItemRequestContainer>EndItemRequestContainerType<EndingReason>EndReasonCodeType</EndingReason>
    <ItemID>ItemIDType (string)</ItemID>
    <MessageID>string</MessageID>
    <SellerInventoryID>string</SellerInventoryID>
  </EndItemRequestContainer>
  
</EndItemsRequest>`
            var o EndItemsRequest
            err := o.FromXML([]byte(data))
            if err != nil {
                t.Errorf("failed loading xml %v",err)
                return
            }
    
                if o.EndItemRequestContainer[0].EndingReason != "EndReasonCodeType" {
                    t.Errorf("failed because o.EndItemRequestContainer[0].EndingReason != EndReasonCodeType '%+v' %+v",o.EndItemRequestContainer[0].EndingReason,o)
                    return
                }
            

                if o.EndItemRequestContainer[0].ItemID != "ItemIDType (string)" {
                    t.Errorf("failed because o.EndItemRequestContainer[0].ItemID != ItemIDType (string) %+v %+v",o.EndItemRequestContainer[0].ItemID,o)
                    return
                }
            

                if o.EndItemRequestContainer[0].MessageID != "string" {
                    t.Errorf("failed because o.EndItemRequestContainer[0].MessageID != string %+v %+v",o.EndItemRequestContainer[0].MessageID,o)
                    return
                }
            

                if o.EndItemRequestContainer[0].SellerInventoryID != "string" {
                    t.Errorf("failed because o.EndItemRequestContainer[0].SellerInventoryID != string %+v %+v",o.EndItemRequestContainer[0].SellerInventoryID,o)
                    return
                }
            

}
