package gobay

import (
	"fmt"
)

func init() {
	globalDebugLevel = DBG_NONE
	globalDebugFunction = func(lvl int, s string) {
		if globalDebugLevel == DBG_NONE {
			return
		}
		if !(globalDebugLevel >= lvl) {
			return
		}
		fmt.Printf("[Gobay] %s\n", s)
	}
}

func SiteIDToCode(id string) string {
	for _, sm := range SiteCodeTypeMap {
		if sm.SiteID == id {
			return sm.SiteValue
		}
	}
	return "UNKNOWN"
}
