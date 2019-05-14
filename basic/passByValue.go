package main

// func main() {
// 	a()
// }

func a() {
	var arr = [2]string{"t", "t2"} //array, slice, map을 각각 선언
	var sli = []string{"t", "t2"}
	var map1 = map[string][]string{"t": sli}

	println("## initial variables")
	println("# array")
	println("object addr : ", &arr)
	println("value addr : ", &arr[0], ", value : ", arr[0])
	println("value2 addr : ", &arr[1], ", value : ", arr[1])
	println("# slice")
	println("object addr : ", &sli)
	println("value addr : ", &sli[0], ", value : ", sli[0]) // same
	println("value2 addr : ", &sli[1], ", value : ", sli[1])
	println("# map")
	println("object addr : ", &map1)
	println("value addr : ", map1["t"], ", value : ", map1["t"])
	print("\n")
	b(arr, sli, map1) //B를 호출하여 전달한 variables를 출력한 후 변경
	print("\n")
	println("## changed variables")
	println("# array")
	println("object addr : ", &arr)
	println("value addr : ", &arr[0], ", value : ", arr[0])
	println("# slice")
	println("object addr : ", &sli)
	println("value addr : ", &sli[0], ", value : ", sli[0]) // same
	println("# map")
	println("object addr : ", &map1)
	println("value addr : ", map1["t"], ", value : ", map1["t"])
}

func b(arr [2]string, sli []string, map1 map[string][]string) {
	println("## passed variables")
	println("# array")
	println("object addr : ", &arr)
	println("value addr : ", &arr[0], ", value : ", arr[0])
	println("# slice")
	println("object addr : ", &sli)
	println("value addr : ", &sli[0], ", value : ", sli[0]) // same
	println("# map")
	println("object addr : ", &map1)
	println("value addr : ", map1["t"], ", value : ", map1["t"])
	//값 변경
	arr[0] = "f"
	sli[0] = "f"
}
