package main
import (
	"fmt" 
	"math")
func main()  {
	var age int = 18
	if age >= 18 {
		fmt.Println("Yes, you can vote")
	} else {
		fmt.Println("No, you can't vote")
	}

	switch age { 
	case 16: fmt.Println("u are 16")
	case 18: fmt.Println("u are 18")
	}

	//Arrays
	var Player[5] string
	Player[0] = "yomi"
	Player[1] = "ope"

	fmt.Println(Player[1])

	//shorter way to declare and array
	Team := [2]string{"madrid","barca"}
	fmt.Println(Team[1])

	for _, myteam :=range Team{
		fmt.Println(myteam)
	}

	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
}
