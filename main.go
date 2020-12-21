package main

import (
	"fmt"
	"time"
)

func main() {
	// for keeping things easy to read and type-safe
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [10]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	// get current time
	now := time.Now()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()

	for line := range digits[0] {
		// Print a line for each placeholder in digits
		for digit := range digits {
			fmt.Print(digits[digit][line], "  ")
		}
		fmt.Println()
	}
}
