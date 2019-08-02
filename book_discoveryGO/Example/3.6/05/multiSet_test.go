package multiSet

import "fmt"

func ExampleMultiSet() {
	m := NewMultiSet()
	fmt.Println(m.String())
	fmt.Println(m.Count("3"))
	m.Insert( "3" )
	m.Insert("3" )
	m.Insert("3" )
	m.Insert("3" )
	fmt.Println(m.String())
	m.Insert("4" )
	m.Insert("2" )
	m.Erase("3")
	fmt.Println(m.Count("3"))
	//Output:
	//{ }
	//0
	//{ 3 3 3 3 }
	//3
}
