package strings

import (
	"fmt"
	"strconv"
)

func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}

	fmt.Printf("\n")

	for _, r := range s {
		fmt.Printf("%x:", r)
	}

	fmt.Printf("\n")

	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	//Output:
	//ea:b0:80:eb:82:98:eb:8b:a4:
	//ac00:b098:b2e4:
	//eab080eb8298eb8ba4
	//ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	//Output:
	//각나다
}

func Example_strCat() {
	a := "wow"
	b := " unbelievable"
	a += b
	fmt.Println(a)
	//Output:
	//wow unbelievable
}

func Example_strCat2() {
	a := "wow"
	b := " unbelievable"
	c := a + b
	fmt.Println(c)

	s := fmt.Sprint(a, " unbelievable")
	fmt.Println(s)
	//Output:
	//wow unbelievable
	//wow unbelievable
}

func Example_strCat3() {
	var i int
	var k int64
	var f float64
	var s string
	var err error
	i, err = strconv.Atoi("350")
	f, err = strconv.ParseFloat("3.14", 64)

	fmt.Println(err)
	fmt.Println(i)
	k, err = strconv.ParseInt("cc7fdd", 16, 32)
	fmt.Println(k)
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32)
	fmt.Println(k)
	fmt.Println(f)
	s = strconv.Itoa(340)
	fmt.Println(s)
	s = strconv.FormatInt(13402077, 16)
	fmt.Println(s)
	t := strconv.FormatInt(100, 2)
	fmt.Println(t)
	//Output:
	//
}

func Example_scan() {
	var num int
	fmt.Sscanf("57", "%d", &num)
	fmt.Println(num)
	var s string
	s = fmt.Sprint(3.14)
	fmt.Println(s)
	//Output:
	//
}
