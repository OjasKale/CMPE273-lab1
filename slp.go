package Slp


import (
	"fmt"
	"time"
)

func sleep(ts int) {

	<-time.After(time.Second*time.Duration(ts))
}

func main()	{
	t1 := time.Now().Second()
	fmt.Print(t1)
	fmt.Print("\nSleeping...\n")
	sleep(20)
	t2 := time.Now().Second()
	fmt.Print(t2)
	if t2-t1 == 20{
		fmt.Print("Wait Success!")
	}
}
