package gobay

type GetEbayDetailsStruct struct {
	DetailNames []string
}

func NewGetEbayDetailsStruct() *GetEbayDetailsStruct {
	var s GetEbayDetailsStruct
	var dnames = []string{
		"CountryDetails",
		"CurrencyDetails",
		"ProductDetails",
		"ReturnPolicyDetails",
		"SiteDetails",
	}

	for _, dn := range dnames {
		s.DetailNames = append(s.DetailNames, dn)
	}
	return &s
}

func GeteBayDetailsTemplate() string {
	return `
    {{ range $x,$y := .DetailNames }}
        <DetailName>{{ $y }}</DetailName>
    {{ end }}
    
`
}
