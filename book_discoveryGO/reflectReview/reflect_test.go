package reflect

import (
	"fmt"
	"reflect"
)

//func ExampleNewMap() {
//	mapmock := map[string]int{"mykey": 15}
//	res := NewMap("mykey", 15)
//	rtype:= reflect.TypeOf(res)
//	rmap := reflect.MakeMap(rtype)
//	rmap["mykey"]= 15
//	fmt.Println(resMap)
//
//	if !reflect.DeepEqual(resMap, mapmock) {
//		fmt.Println("not equal")
//		return
//	}
//	fmt.Println("equal")
//	return
//
//	//Output:
//	//equal
//}

func ExampleReflect(){
	a:= map[string]int64{"mykey":64}
	//b:= 15

	x := reflect.Value{}
	y := reflect.ValueOf(a)
	z1 := y.Type()
	z2 := reflect.TypeOf(a)
	z3 := z2.Kind()
	z4:= y.Kind()

	//test := reflect.Value{z}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z1)
	fmt.Println(z2)
	fmt.Println(z3)
	fmt.Println(z4)
	//fmt.Println(z4)
	//Unlike reflection.Value, reflection.Type && reflection.Kind is type comparable
	fmt.Println(z1==z2)
	fmt.Println(z3==z4)

	fmt.Println("--------------------------------------------------")
	//Output:
	//fail
}
