package words

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	fmt.Println("Hi Noi")
	words := strings.Fields(s) // HL
	r := map[string]int{}
	for _, w := range words {
		r[w] = r[w] + 1 // HL
	}
	return r

}

func Split(t string) map[string]int {
	m := make(map[string]int)
	someString := t
	if t == "" {
		someString = "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	}

	someString = strings.Replace(someString, ",", "", -1)
	someString = strings.Replace(someString, ".", "", -1)

	words := strings.Fields(someString)

	for _, v := range words {
		//fmt.Println(i, v)
		if m[v] == 0 {
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
	}

	for i, v := range m {
		fmt.Println(i, v)
	}

	return m
}
