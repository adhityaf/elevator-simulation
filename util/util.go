package util

import "fmt"

func IsNameLengthValid(name string) bool {
	if len(name) < 5 {
		return false
	}

	return true
}

func IsFloorValid(floor int) bool {
	if floor >= 1 && floor <= 10 {
		return true
	}

	return false
}

func LiftGoUp(current, destination int) {
	for i := current + 1; i <= destination; i++ {
		fmt.Printf("Lift naik ke lantai : %d\n", i)
	}
}

func LiftGoDown(current, destination int) {
	for i := current - 1; i >= destination; i-- {
		fmt.Printf("Lift turun ke lantai : %d\n", i)
	}
}
