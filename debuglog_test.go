package debuglog

import "testing"

type TestJson struct {
	Id   int64
	Name string
}

func TestDebuglog(t *testing.T) {
	intVal := 123
	Val(intVal)
	Val(intVal, "prefix1")
	Val(intVal, "prefix1", "prefix2")

	testJson := TestJson{Id: 987, Name: "alan"}

	SpewVal(testJson, "testJson")
	OctUtf8Val(testJson, "testJson")

	ToJson(testJson, "testJson")
	ToJsonPretty(testJson, "testJson")
}
