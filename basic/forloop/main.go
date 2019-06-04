package main

import (
	"fmt"
	"net/url"
)

func main() {
	userURLstring := "https://www.google.com/search?newwindow=1,2&sxsrf=qwu&ei=jMH1XLW0NeyUr7wPjdyY2Ak&q=ohmy&oq=ohmy&gs_l=psy-ab.3..0l10.1575092.1575844..1576179...0.0..0.123.457.0j4....3..0....1..gws-wiz.......35i39.94t15Jbgd2s"
	userURL, _ := url.Parse(userURLstring)
	userQuery := userURL.Query()
	for _, s := range userQuery {
		fmt.Println(s[0])
	}

	arr := [][]string{{"one", "two"}, {"three"}, {"four"}}
	for _, s := range arr {
		fmt.Println(s[0])
	}
}
