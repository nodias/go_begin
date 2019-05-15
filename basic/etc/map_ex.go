package main

// func main() {
// 	map1()
// }

var idMap map[int]string
var idArray [3]string

func map1() {
	idMap = make(map[int]string)
	idMap[1] = "value1"
	idMap[2] = "value2"

	idArray = [3]string{"arr0", "arr1", "arr2"}
	println(idArray[2])
	arrayValueChanger()
	println(idArray[2])

	println(idMap[2])
	mapValueChanger()
	println(idMap[2])
}

func mapValueChanger() {
	println(idMap[2])
	idMap[2] = "value2change!"
}

func arrayValueChanger() {
	println(idMap[2])
	idArray[2] = "arr2change!"
}
