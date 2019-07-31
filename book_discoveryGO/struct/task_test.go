package task

import (
	"encoding/json"
	"fmt"
)

type ByteSize float64

const (
	KB ByteSize = 1 << (10 * (iota + 1))
	MB
	GB
	TB
	PB
)

func ExampleConst() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)

	//Output:
	//.
}

// func ExampleOverdue() {
// 	a := deadline(time.Now().Add(time.Hour * (-1 * 4))) //past
// 	b := deadline(time.Now().Add(time.Hour * (1 * 4)))  //future
// 	var c *deadline

// 	fmt.Println(a.Overdue())
// 	fmt.Println(b.Overdue())
// 	fmt.Println(c.Overdue())

// 	//Output:
// 	//true
// 	//false
// 	//false
// }

// func ExampleTaskOverdue() {
// 	task := []Task{{"title", DONE, Deadline(time.Now().Add(time.Hour * (-1 * 4)))}}
// 	task.Overdue()
// }

// func Example_MarshalJSON() {
// 	t := Task{
// 		"Laundry",
// 		DONE,
// 		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
// 		1,
// 	}

// 	b, err := json.Marshal(t)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(b))

// 	t1 := Task{}
// 	err = json.Unmarshal(b, &t1)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(t1)

// 	//Output:
// 	//{"Title":"Laundry","Status":"DONE","Deadline":1439739780}
// 	//{Laundry 2 2015-08-17 00:43:00 +0900 KST}
// }

type MyStruct struct {
	Title    string `json:"title"`
	Internal string `json:"-"`
	Value    string `json:",omitempty"`
	ID       int64  `json:",string`
}

func Example_Marshal() {
	ms := MyStruct{"title", "internal", "", 16}
	b, err := json.Marshal(ms)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	ms2 := MyStruct{}
	json.Unmarshal(b, &ms2)
	fmt.Println(ms2)
	//Output:
	//
}

func Example_mapMarshalJSON() {
	b, _ := json.Marshal(map[string]interface{}{
		"Name": "John",
		"Age":  16,
	})
	fmt.Println(string(b))
	//Output:
	//
}

//두개의 필드가 같을땐 둘 다 안나옴
type Blocks struct {
	VisibleField2  string
	InvisibleField string
}
type Fields struct {
	VisibleField   string `json:"visibleField"`
	InvisibleField string `json:"invisibleField"`
}

func ExampleOmitFields() {
	f := &Fields{"a", "b"}
	b, _ := json.Marshal(struct {
		*Fields
		InvisibleField string `json:"invisibleField,omitempty"`
		AdditionField  string `json:"additionField"`
	}{Fields: f, AdditionField: "c"})
	fmt.Println(string(b))
	//Output:
}

type MyFields struct {
	*Fields
	*Blocks
	InvisibleField string
	Wow            string
}

func ExamplePrac() {
	f := &Fields{"a", "b"}
	h := &Blocks{"a", "b"}

	m := &MyFields{f, h, "in", "wow"}
	b, _ := json.Marshal(m)
	// b2, _ := json.Marshal(f)
	fmt.Println(string(b))
	// fmt.Println(string(b2))

	//Output:
}

// func ExampleTask_String() {
// 	fmt.Println(Task{"Laundry", DONE, nil, 1})
// 	//Output:
// }

func ExampleIncludeSubTasks_String() {
	fmt.Println(IncludeSubTasks(Task{
		Title:    "Laundry",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{{
			Title:    "Wash",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []Task{
				{"Put", DONE, nil, 2, nil},
				{"Detergent", TODO, nil, 2, nil},
			},
		}, {
			Title:    "Dry",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}, {
			Title:    "Fold",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}))
	//Output:
}
