package main

import "fmt"

func Burbuja(s []int64)  {
	for limit := len(s) - 1; limit > 0; limit-- {
		for i := 0; i < limit; i++ {
			if s[i] > s[i + 1] {
				s[i], s[i + 1] = s[i + 1], s[i]
			}
		}
	}
}

func main()  {
	s := []int64{8, 4, 9, 1, 7, 3, 0, 6, 5, 2}
	Burbuja(s)
	fmt.Println(s)
}