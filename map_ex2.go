package main

// func main() {
// 	testInstance()
// }

type test struct {
	myarray [3]string
	myslice []string
	mymap   map[string]string
	mymap2  map[string]interface{}
}

func testInstance() {
	var myarray = [3]string{"arr1", "arr2", "arr3"}
	var myslice = []string{"arr1", "arr2", "arr3"}
	var mymap = map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	var mymap2 = map[string]interface{}{"mkey1": myslice}
	printAll(myarray, myslice, mymap, mymap2)
	updateAll(myarray, myslice, mymap, mymap2)
	printAll(myarray, myslice, mymap, mymap2)
}

func printAll(myarray [3]string, myslice []string, mymap map[string]string, mymap2 map[string]interface{}) {
	println("## printAll")
	println(myarray[2])
	println(myslice[2])
	println(mymap2["mkey1"].([]string)[2])
	println(mymap["key3"])
}

func updateAll(myarray [3]string, myslice []string, mymap map[string]string, mymap2 map[string]interface{}) {
	myarray[2] = "false"
	myslice[2] = "false"
	mymap["key3"] = "false"
	println("## updateAll")
	printAll(myarray, myslice, mymap, mymap2)
}
