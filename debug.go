package gobay

const (
	DBG_DEBUG = iota
	DBG_INFO
	DBG_WARN
	DBG_ERROR
	DBG_NONE
)

type DebugFunc func(lvl int, s string)

var globalDebugFunction DebugFunc
var globalDebugLevel int

func SetDebugLevel(l int) {
	globalDebugLevel = l
}

func SetDebugFunction(f DebugFunc) {
	globalDebugFunction = f
}

