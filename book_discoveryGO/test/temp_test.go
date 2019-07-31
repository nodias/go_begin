package test

import (
	"fmt"
	"path/filepath"
)

func Example_test01() {
	var url = "http://wow.com/"
	fmt.Println(filepath.Glob("error*"))
	fmt.Println(filepath.Base(url))

	//Output:
	//wow
}