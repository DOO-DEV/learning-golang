package main

import (
	"fmt"
	clear "learning-golang/projects/clock/clear_screen"
	"time"
)

type placeholders [5]string

func main() {
	zero := placeholders{
		"###",
		"# #",
		"# #",
		"# #",
		"###",
	}
	one := placeholders{
		"## ",
		" # ",
		" # ",
		" # ",
		"###",
	}
	two := placeholders{
		"###",
		"  #",
		"###",
		"#  ",
		"###",
	}
	three := placeholders{
		"###",
		"  #",
		"###",
		"  #",
		"###",
	}
	four := placeholders{
		"# #",
		"# #",
		"###",
		"  #",
		"  #",
	}
	five := placeholders{
		"###",
		"#  ",
		"###",
		"  #",
		"###",
	}
	six := placeholders{
		"###",
		"#  ",
		"###",
		"# #",
		"###",
	}
	seven := placeholders{
		"###",
		"  #",
		"  #",
		"  #",
		"  #",
	}
	eight := placeholders{
		"###",
		"# #",
		"###",
		"# #",
		"###",
	}
	nine := placeholders{
		"###",
		"# #",
		"###",
		"  #",
		"###",
	}
	colon := placeholders{
		"   ",
		" # ",
		"   ",
		" # ",
		"   ",
	}

	digits := [...]placeholders{zero, one, two, three, four, five, six, seven, eight, nine}

	

	for {
		now := time.Now()
		h1, h2 := numberDestructure(now.Hour())
		m1, m2 := numberDestructure(now.Minute())
		s1, s2 := numberDestructure(now.Second())
		display([8]placeholders{digits[h1], digits[h2], colon, digits[m1], digits[m2], colon, digits[s1], digits[s2]})
		
		time.Sleep(time.Second)
		clear.Clear()
		fmt.Printf("\033[%d;%dH", 0, 0)
	}

}

func display(digits [8]placeholders) {
	for i := 0; i < 5; i++ {
		for digit := range digits {
			fmt.Print(digits[digit][i], "  ")
		}
		fmt.Println()
	}
}

func numberDestructure(num int) (int, int) {
	a := num / 10
	b := num % 10
	return a, b
}
