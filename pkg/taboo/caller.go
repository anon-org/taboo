package taboo

import (
	"regexp"
	"runtime"
)

// caller stores runtime caller function, line, and filename
type caller struct {
	function string
	line     int
}

// call returns *caller with defined skip frame
func call(skip int) *caller {
	pc, _, line, ok := runtime.Caller(skip)

	var fn string

	if ok {
		fn = runtime.FuncForPC(pc).Name()
		re := regexp.MustCompile("\\.func\\d$")

		fn = re.Split(fn, -1)[0]
	}

	return &caller{
		function: fn,
		line:     line,
	}
}

func (c caller) Function() string {
	return c.function
}

func (c caller) Line() int {
	return c.line
}
