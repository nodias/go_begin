package interfaceTest

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	task "github.com/nodias/go_begin/book_discoveryGO/structPrac"
)

type MySlice []string

func (m MySlice) len() int {
	return len(m)
}

//
func (m MySlice) less(i, j int) bool {
	return strings.ToLower(m[i]) < strings.ToLower(m[j]) || m[i] < m[j] && strings.ToLower(m[i]) < strings.ToLower(m[j])
}

func (m MySlice) swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func Example_MySlice() {
	mySlice := MySlice{
		"sli1", "sli2", "Sli3",
	}
	mySlice_length := mySlice.len()
	fmt.Println(mySlice_length)
	mySlice_less := mySlice.less(0, 1)
	fmt.Println(mySlice_less)
	fmt.Println(mySlice)
	mySlice.swap(0, 1)
	mySlice_less = mySlice.less(0, 1)
	fmt.Println(mySlice_less)
	fmt.Println(mySlice)
	//Output:
}

func Example_bool() {
	fmt.Println("A" > "a")
	//Output:
	//
}

type FileSystem interface {
	Rename(oldpath, newpath string) error
	Remove(name string) error
}

type OSFileSystem struct{}

func (fs OSFileSystem) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (fs OSFileSystem) Remove(name string) error {
	return os.Remove(name)
}

func Example() {
	t := task.Task{
		Title:    "Laundry",
		Status:   task.DONE,
		Deadline: nil,
	}
	fmt.Println(Join(",", 1, "two", 3, t))
	//Output:
}

func Join(sep string, a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}
	str := make([]string, len(a))
	for i := range a {
		switch x := a[i].(type) {
		case string:
			str[i] = x
		case int:
			str[i] = strconv.Itoa(x)
		case fmt.Stringer:
			str[i] = x.String()
		}
	}
	return strings.Join(str, sep)
}
