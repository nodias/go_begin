package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func Example_A() {
	person := Person{
		482194269219498,
		"nodias@naver.com",
		FRIENDS,
		ResponseErr{nil},
	}

	// person.Cheat()
	person_json, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(person_json))

	person_unmarshal := Person{}
	err = json.Unmarshal(person_json, &person_unmarshal)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(person_unmarshal)
	//Output:
}

func Example_err() {
	person := Person2{
		nil,
	}

	// person.Cheat()
	b, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	person_unmarshal := Person2{}
	err = json.Unmarshal(b, &person_unmarshal)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(person_unmarshal)

	//Output:
}

// func Example_ResponseErr_MarshalJSON() {
// 	re := ResponseErr{Cheated}
// 	b, err := re.MarshalJSON()
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println(string(b))

// 	re2 := ResponseErr{nil}
// 	err = re2.UnmarshalJSON(b)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(reflect.DeepEqual(re2, re))
// 	fmt.Printf("re2 : %v \n re : %v", re2, re)

// 	//Output:
// }

func Example_stringMarshal() {
	s := "wow"
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	var s2 string
	err = json.Unmarshal(b, &s2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
	//Output:
}
func Example_intMarshal() {
	s := 46200
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	var s2 int
	err = json.Unmarshal(b, &s2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
	//Output:
}

type ResponseErr2 struct {
	Err error `json:err`
}

func Example_interface() {
	var a interface{}
	a = ResponseErr2{nil}
	fmt.Println(reflect.TypeOf(a))
	switch a.(type) {
	case string:
		fmt.Println("string")
	case int64:
		fmt.Println("int64")
	case int32:
		fmt.Println("int32")
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case ResponseErr:
		fmt.Println("ResponseErr")
	case ResponseErr2:
		fmt.Println("ResponseErr2")
	case interface{}:
		fmt.Println("interface{]")
	default:
		fmt.Println(reflect.TypeOf(a))
	}

	//Output:
}

func Example_Unmarshal_nil() {
	var v interface{}
	fmt.Println(json.Unmarshal([]byte("null"), &v))
	fmt.Println(json.Unmarshal([]byte("nil"), &v))
	//Output:
}
