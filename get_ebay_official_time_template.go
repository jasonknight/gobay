package gobay
//GeteBayOfficialTimeRequest
func GeteBayOfficialTimeTemplate() string {
    return `
<ErrorLanguage> {{ .Language }}</ErrorLanguage>
<MessageID> {{ .MessageID }} </MessageID>
<Version> {{ .CompatLevel }} </Version>
<WarningLevel> WarningLevelCodeType </WarningLevel>
`
}