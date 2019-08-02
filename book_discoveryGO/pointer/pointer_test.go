package pointer

import "fmt"

func addInt(a int) {
	a++
}

func Example_Int(){
	a:= 15
	addInt(a)
	fmt.Println(a)
	//Output:
	//
}

func addArray(a [3]int){
	a[0] = 1
}

func Example_Array(){
	a:= [3]int{0,0,0}
	addArray(a)
	fmt.Println(a)
	//Output:
	//
}

func addSlice(a []int){
	a = append(a, 1)
	a[0] = 1
}

func Example_Slice(){
	a:= []int{0,0,0}
	addSlice(a)
	fmt.Println(a)
	//Output:
	//
}


func addMap(a map[string]int){
	a["first"] = 1
	a["wow"] = 2
}

func Example_Map(){
	a := map[string]int{"first":0}
	addMap(a)
	fmt.Println(a)
	//Output:
	//
}
