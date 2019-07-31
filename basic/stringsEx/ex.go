package stringsEx

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	s1 := []string{"Hello,", "World"}
	fmt.Println(strings.Join(s1, " "))
	s2 := strings.Split("Hello, World!", " ")
	fmt.Println(s2[1])
	s3 := strings.Fields("Hello, World!")
	fmt.Println(s3[0])

	var a interface{}
	var b string
	b = "this is string"
	a = b
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	v1 := `{
				"name":"thun"
		}`
	var data map[string]string
	json.Unmarshal([]byte(v1), &data)
	fmt.Println(data)
}
