package gobay

var SiteIDMap = [...]struct {
	SiteID      string
	SiteName    string
	GlobalID    string
	TerritoryID string
	Language    string
}{
	{GlobalID: "EBAY-AT", Language: "de-AT", TerritoryID: "AT", SiteName: "eBay Austria   ", SiteID: "16"},
	{GlobalID: "EBAY-AU", Language: "en-AU", TerritoryID: "AU", SiteName: "eBay Australia ", SiteID: "15"},
	{GlobalID: "EBAY-CH", Language: "de-CH", TerritoryID: "CH", SiteName: "eBay Switzerland   ", SiteID: "193"},
	{GlobalID: "EBAY-DE", Language: "en-DE", TerritoryID: "DE", SiteName: "eBay Germany   ", SiteID: "77"},
	{GlobalID: "EBAY-ENCA", Language: "en-CA", TerritoryID: "CA", SiteName: "eBay Canada (English)  ", SiteID: "2"},
	{GlobalID: "EBAY-ES", Language: "es-ES", TerritoryID: "ES", SiteName: "eBay Spain ", SiteID: "186"},
	{GlobalID: "EBAY-FR", Language: "fr-FR", TerritoryID: "FR", SiteName: "eBay France", SiteID: "71"},
	{GlobalID: "EBAY-FRBE", Language: "fr-BE", TerritoryID: "BE", SiteName: "eBay Belgium (French)  ", SiteID: "23"},
	{GlobalID: "EBAY-FRCA", Language: "fr-CA", TerritoryID: "CA", SiteName: "eBay Canada (French)   ", SiteID: "210"},
	{GlobalID: "EBAY-GB", Language: "en-GB", TerritoryID: "GB", SiteName: "eBay UK", SiteID: "3"},
	{GlobalID: "EBAY-HK", Language: "zh-Hant", TerritoryID: "HK", SiteName: "eBay Hong Kong", SiteID: "201"},
	{GlobalID: "EBAY-IE", Language: "en-IE", TerritoryID: "IE", SiteName: "eBay Ireland   ", SiteID: "205"},
	{GlobalID: "EBAY-IN", Language: "en-IN", TerritoryID: "IN", SiteName: "eBay India ", SiteID: "203"},
	{GlobalID: "EBAY-IT", Language: "it-IT", TerritoryID: "IT", SiteName: "eBay Italy ", SiteID: "101"},
	{GlobalID: "EBAY-MOTOR", Language: "en-US", TerritoryID: "US", SiteName: "eBay Motors", SiteID: "100"},
	{GlobalID: "EBAY-MY", Language: "en-MY", TerritoryID: "MY", SiteName: "eBay Malaysia  ", SiteID: "207"},
	{GlobalID: "EBAY-NL", Language: "nl-NL", TerritoryID: "NL", SiteName: "eBay Netherlands   ", SiteID: "146"},
	{GlobalID: "EBAY-NLBE", Language: "nl-BE", TerritoryID: "BE", SiteName: "eBay Belgium (Dutch)   ", SiteID: "123"},
	{GlobalID: "EBAY-PH", Language: "en-PH", TerritoryID: "PH", SiteName: "eBay Philippines   ", SiteID: "211"},
	{GlobalID: "EBAY-PL", Language: "pl-PL", TerritoryID: "PL", SiteName: "eBay Poland", SiteID: "212"},
	{GlobalID: "EBAY-SG", Language: "en-SG", TerritoryID: "SG", SiteName: "eBay Singapore ", SiteID: "216"},
	{GlobalID: "EBAY-US", Language: "en-US", TerritoryID: "US", SiteName: "eBay United States ", SiteID: "0"},
}

var SiteCodeTypeMap = [...]struct {
	SiteValue string
	SiteID    string
	Currency  string
}{
	{SiteValue: "Australia", SiteID: "15", Currency: "AU"},
	{SiteValue: "Austria", SiteID: "16", Currency: "EU"},
	{SiteValue: "Belgium_Dutch", SiteID: "123", Currency: "EU"},
	{SiteValue: "Belgium_French", SiteID: "23", Currency: "EU"},
	{SiteValue: "Canada", SiteID: "2", Currency: "CAD"},
	{SiteValue: "CanadaFrench", SiteID: "210", Currency: "CAD"},
	{SiteValue: "France", SiteID: "71", Currency: "EU"},
	{SiteValue: "Germany", SiteID: "77", Currency: "EU"},
	{SiteValue: "HongKong", SiteID: "201", Currency: "HK"},
	{SiteValue: "India", SiteID: "203", Currency: "IN"},
	{SiteValue: "Ireland", SiteID: "205", Currency: "EU"},
	{SiteValue: "Italy", SiteID: "101", Currency: "EU"},
	{SiteValue: "Malaysia", SiteID: "207", Currency: "MY"},
	{SiteValue: "Netherlands", SiteID: "146", Currency: "EU"},
	{SiteValue: "Philippines", SiteID: "211", Currency: "PH"},
	{SiteValue: "Poland", SiteID: "212", Currency: "PL"},
	{SiteValue: "Russia", SiteID: "215", Currency: "RUB"},
	{SiteValue: "Singapore", SiteID: "216", Currency: "SG"},
	{SiteValue: "Spain", SiteID: "186", Currency: "EU"},
	{SiteValue: "Switzerland", SiteID: "193", Currency: "CH"},
	{SiteValue: "UK", SiteID: "3", Currency: "GB"},
	{SiteValue: "US", SiteID: "0", Currency: "US"},
}
