package gobay

type GetCategoriesStruct struct {
	CategoryParent  []string
	SiteID          string
	LevelLimit      string
	ViewAllNodes    string
	DetailLevels    []string `xml:"DetailLevel"`
	OutputSelectors []string `xml:"OutputSelector"`
}

func NewGetCategoriesStruct() *GetCategoriesStruct {
	var s GetCategoriesStruct
	s.LevelLimit = "3"
	s.ViewAllNodes = "false"
	s.DetailLevels = append(s.DetailLevels, "ReturnAll")
	outputs := [...]string{
		"CategoryID",
		"CategoryLevel",
		"CategoryName",
		"CategoryParentID",
	}
	for _, v := range outputs {
		s.OutputSelectors = append(s.OutputSelectors, v)
	}
	return &s
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
