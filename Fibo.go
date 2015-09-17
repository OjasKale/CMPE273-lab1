package main

import "fmt"

func Fibc(num uint) uint {
    if num == 0 {
        return 0
    } else if num == 1 {
		return 1
	} else {
		return Fibc(num - 1) + Fibc(num-2)
	}
}

func main() {
	var n uint
  fmt.Println("Enter the position of the number you want in fibonacci series: ")
  fmt.Scanf("%d",&n)

	fmt.Println(Fibc(n))
}
