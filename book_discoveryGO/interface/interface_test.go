package interfaceTest

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

/**
type Interface interface{
	//Len is the number of elements in the collection.
	Len() int
	//Less reports whether the element with
	//index i should sort before the element with index j.
	Less(i, j int) bool
	//Swap swaps the elements with indexes i and j.
	Swap(i,j int)
}
*/

type CaseInsensitive []string

func (c CaseInsensitive) Len() int {
	return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) ||
		(strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}

func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func ExampleCaseInsensitive_sort() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	sort.Sort(apple)
	fmt.Println(apple)
	//Output:
	//
}

/**
type Interface interface{
	sort.Interface
	Push(x interface{}) //add x as element Len()
	Pop() interface{} //remove and return element Len() -1
}
*/

func (c *CaseInsensitive) Push(x interface{}) {
	*c = append(*c, x.(string))
}

func (c *CaseInsensitive) Pop() interface{} {
	len := c.Len()
	last := (*c)[len-1]
	*c = (*c)[:len-1]
	return last
}

func ExampleCaseInsensitive_heap() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		// fmt.Println(heap.Pop(&apple))
		popped := heap.Pop(&apple)
		s := popped.(string)
		fmt.Println(s)
	}
	//Output:
}
