package main

import "fmt"

func Josephus(people, count int32) int32 {
	if people == 1 {
		return 0
	} else {
		return (Josephus(people-1, count) + count) % people
	}
}

func main() {
	fmt.Println(Josephus(41, 3) + 1)
}
