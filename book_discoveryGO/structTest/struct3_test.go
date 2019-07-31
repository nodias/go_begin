package structTest

import "fmt"

type Address struct {
	City  string
	State string
}

type Telephone struct {
	Mobile string
	Direct string
}

type Contact struct {
	Address
	Telephone
}

type Contact2 struct {
	*Address
	*Telephone
}

func ExampleContact() {
	// var c Contact
	// c.Mobile = "010-1212-3434"
	// fmt.Println(c.Telephone.Mobile)
	// c.Address.City = "San Farancisco"
	// c.State = "CA"
	// c.Direct = "N/A"
	// fmt.Println(c)

	var c2 Contact2
	fmt.Println(c2)
	c2.Address = &Address{}
	c2.Telephone = &Telephone{}
	fmt.Println(c2)
	//Output:
	//{<nil> <nil>}
	//{0xc00004a4e0 0xc00004a500}

}
