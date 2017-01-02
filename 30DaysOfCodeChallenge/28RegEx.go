package main

import (
	"fmt"
	"regexp"
	"sort"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)
	names := []string{}
	for i := 0; i < n; i++ {
		var firstName, emailID string
		fmt.Scan(&firstName, &emailID)
		ok, _ := regexp.MatchString("[a-z]+@gmail.com", emailID)
		if ok {
			names = append(names, firstName)
		}
	}
	sort.Strings(names)
	for _, s := range names {
		fmt.Println(s)
	}
}
