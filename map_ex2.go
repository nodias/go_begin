package main

func main() {
	testInstance()
}

type test struct {
	myarray [3]string
	myslice []string
	mymap   map[string]string
}

func testInstance() {
	var myarray = [3]string{"arr1", "arr2", "arr3"}
	var myslice = []string{"arr1", "arr2", "arr3"}
	var mymap = map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	var mymap2 = map[string]interface{}{"mkey1": myslice}

	var t1 = test{myarray, myslice, mymap}

	println("t1")
	println(t1.myarray[2])
	println(t1.myslice[2])
	println(t1.mymap["key3"])
	println("var")
	println(myarray[2])
	println(myslice[2])
	println(mymap["key3"])

	// println(myarray)
	println(myslice)
	println(mymap2["mkey1"])
}
