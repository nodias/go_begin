package main

// func main() {
// 	var emps = make([]Emp, 1)
// 	var emp = Emp{"babo", 1000}
// 	emps[0] = emp

// 	for _, emp := range emps {
// 		emp.salary = 2000
// 	}
// 	println(emps[0].salary)

// 	for i := 0; i < len(emps); i++ {
// 		emps[i].salary = 2000
// 	}

// 	println(emps[0].salary)

// 	var a = [1]string{"a"}
// 	b := a
// 	b[0] = "b"

// 	println(a[0])
// }

type Emp struct {
	name   string
	salary int
}
