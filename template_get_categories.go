package gobay

type GetCategoriesStruct struct {
	CategoryParent  []string
	SiteID          string
	LevelLimit      string
	ViewAllNodes    string
	DetailLevels    []string `xml:"DetailLevel"`
	OutputSelectors []string `xml:"OutputSelector"`
}

func GetAllCategoriesTemplate() string {
	return `
<CategorySiteID>{{ .SiteID }}</CategorySiteID>
<LevelLimit>{{ .LevelLimit }}</LevelLimit>
<ViewAllNodes>{{ .ViewAllNodes }}</ViewAllNodes>
{{ range $k,$v := .DetailLevels }}
    <DetailLevel>{{$v}}</DetailLevel>
{{ end }}
{{ range $k,$v := .OutputSelectors }}
    <OutputSelector>{{$v}}</OutputSelector>
{{ end }}
`
}
func GetChildCategoriesTemplate() string {
	return `
{{ range $k,$v := .CategoryParents }}
<CategoryParent>$v</CategoryParent>
{{ end }}
<CategorySiteID>{{ .SiteID }}</CategorySiteID>
<LevelLimit>{{ .LevelLimit }}</LevelLimit>
<ViewAllNodes>{{ .ViewAllNodes }}</ViewAllNodes>
{{ range $k,$v := .DetailLevels }}
    <DetailLevel>{{$v}}</DetailLevel>
{{ end }}
{{ range $k,$v := .OutputSelectors }}
    <OutputSelector>{{$v}}</OutputSelector>
{{ end }}
`
}
