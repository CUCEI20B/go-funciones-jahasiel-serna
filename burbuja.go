package main

import "fmt"

func Burbuja(s []int64)  {
	swipe := func(s []int64, a, b int) {
		t := s[a]
		s[a] = s[b]
		s[b] = t
	}
	
	for limit := len(s) - 1; limit > 0; limit-- {
		for i := 0; i < limit; i++ {
			if s[i] > s[i + 1] {
				swipe(s, i, i + 1)
			}
		}
	}
}

func main()  {
	s := []int64{8, 4, 9, 1, 7, 3, 0, 6, 5, 2}
	Burbuja(s)
	fmt.Println(s)
}