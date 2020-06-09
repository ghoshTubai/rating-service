package leonidas

import (
	"runtime"
	"strings"
)

func WhereAmI(depthList ...int) (functionName string, lineNumber int, fileName string) {
	var depth int
	if depthList == nil {
		depth = 2
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return chopFunc(runtime.FuncForPC(function).Name()), line, chopPath(file)
}

func chopPath(ori string) string {
	i := strings.LastIndex(ori, "/")
	if i == -1 {
		return ori
	}
	return ori[i+1:]
}

func chopFunc(ori string) string {
	i := strings.LastIndex(ori, "/")
	if i == -1 {
		return ori
	}
	return ori[i+1:]
}
