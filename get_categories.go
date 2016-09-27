package gobay

type GetCategoriesStruct struct {
	CategoryParent     []string
	SiteID             string
	CategoryLevelLimit string
	ViewAllNodes       string
	DetailLevels       []string
	OutputSelectors    []string
}

func GetAllCategoriesTemplate() string {
	return `
<CategorySiteID>{{ .SiteID }}</CategorySiteID>
<LevelLimit>{{ .CategoryLevelLimit }}</LevelLimit>
<ViewAllNodes>{{ .ViewAllNodes }}</ViewAllNodes>
{{ range $k,$v := .DetailLevels }}
    <DetailLevel>$v</DetailLevel>
{{ end }}
{{ range $k,$v := .OutputSelectors }}
    <OutputSelector>$v</OutputSelector>
{{ end }}
`
}
func GetChildCategoriesTemplate() string {
	return `
{{ range $k,$v := .CategoryParents }}
<CategoryParent>$v</CategoryParent>
{{ end }}
<CategorySiteID>{{ .SiteID }}</CategorySiteID>
<LevelLimit>{{ .CategoryLevelLimit }}</LevelLimit>
<ViewAllNodes>{{ .ViewAllNodes }}</ViewAllNodes>
{{ range $k,$v := .DetailLevels }}
    <DetailLevel>$v</DetailLevel>
{{ end }}
{{ range $k,$v := .OutputSelectors }}
    <OutputSelector>$v</OutputSelector>
{{ end }}
`
}
