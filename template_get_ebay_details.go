package gobay

type EbayDetails struct {
	DetailNames []string
}

func GeteBayDetailsTemplate() string {
	return `
    {{ range $x,$y := .DetailNames }}
        <DetailName>{{ $y }}</DetailName>
    {{ end }}
    
`
}
