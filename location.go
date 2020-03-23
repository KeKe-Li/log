package log

import (
	"path"
	"runtime"
	"strconv"
	"strings"
)

func callerLocation(skip int) string {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return "???"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return trimFileName(file) + ":" + strconv.Itoa(line)
	}
	return trimFuncName(fn.Name()) + "(" + trimFileName(file) + ":" + strconv.Itoa(line) + ")"
}

func trimFuncName(name string) string {
	return path.Base(name)
}

func trimFileName(name string) string {
	i := strings.Index(name, "/src/") + len("/src/")
	if i >= len("/src/") && i < len(name) /* BCE */ {
		name = name[i:]
	}
	i = strings.LastIndex(name, "/vendor/") + len("/vendor/")
	if i >= len("/vendor/") && i < len(name) /* BCE */ {
		return name[i:]
	}
	return name
}
