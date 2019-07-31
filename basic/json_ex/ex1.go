package main

func main() {
	// 	ex1()
}

// type Person struct {
// 	Number int    `json:"number"`
// 	Name   string `json:"name"`
// 	Age    int    `json:"age"`
// 	Gender int    `json:"gender"`
// }

// const (
// 	MAN   = 1
// 	WOMAN = 2
// )

// func ex1() {
// 	var persons = []Person{{1, "jomin", 27, MAN}, {2, "jodang", 4, WOMAN}}
// 	b, _ := json.Marshal(persons)
// 	fmt.Println(string(b))

// 	var jsonPersons = new([]Person)
// 	json.Unmarshal(b, jsonPersons)

// 	ioutil.WriteFile("easyJsonSample.json", b, os.FileMode(0644))
// 	x, _ := ioutil.ReadFile("easyJsonSample.json")
// 	fmt.Println(string(x))
// }
