package gobay

import "testing"

func TestEndItemsResult(t *testing.T) {
	data := `<?xml version="1.0" encoding="utf-8"?>
<EndItemsResponse xmlns="urn:ebay:apis:eBLBaseComponents">
  <EndItemResponseContainer>
    <EndTime>2015-10-01T21:57:01.000Z</EndTime>
    <CorrelationID>1</CorrelationID>
  </EndItemResponseContainer>
  <EndItemResponseContainer>
    <EndTime>2015-10-01T21:57:01.000Z</EndTime>
    <CorrelationID>2</CorrelationID>
  </EndItemResponseContainer>
</EndItemsResponse>`
	var o EndItemsResult
	err := o.FromXML([]byte(data))
	if err != nil {
		t.Errorf("failed loading xml %v", err)
		return
	}

	if o.EndItemResultContainer[0].EndTime != "2015-10-01T21:57:01.000Z" {
		t.Errorf("failed because o.EndItemResultContainer[0].EndTime != 2015-10-01T21:57:01.000Z %+v %+v", o.EndItemResultContainer[0].EndTime, o)
		return
	}

	if o.EndItemResultContainer[0].CorrelationID != 1 {
		t.Errorf("failed because o.EndItemResultContainer[0].CorrelationID != 1 %+v %+v", o.EndItemResultContainer[0].CorrelationID, o)
		return
	}

}
