# debuglog
Easy log variable for debugging

## Usage

`go get -u github.com/chinaran/debuglog`

## Functions

### debuglog.Val()

print a value

### debuglog.SpewVal()

print a value using spew [https://github.com/davecgh/go-spew](https://github.com/davecgh/go-spew)

### debuglog.ToJson()

print a value (json string)

### debuglog.ToJsonPretty()

print a value (pretty json string)

## Example

```golang
package main

import "github.com/chinaran/debuglog"

type TestJson struct {
	Id   int64
	Name string
}

func main() {
	intVal := 123
	debuglog.Val(intVal)
	debuglog.Val(intVal, "prefix1")
	debuglog.Val(intVal, "prefix1", "prefix2")

	testJson := TestJson{Id: 987, Name: "alan"}

	debuglog.SpewVal(testJson, "testJson")
	// debuglog.OctUtf8Val(testJson, "testJson")

	debuglog.ToJson(testJson, "testJson")
	debuglog.ToJsonPretty(testJson, "testJson")
}
```

result:

![result](https://github.com/chinaran/my-pictures/blob/master/debuglog/result.png)
