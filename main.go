package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/adhityaf/elevator-simulation/util"
)

const (
	maxFloor = 10
	minFloor = 1
)

func main() {
	for{
		var (
			name                                        string
			randomFloor, currentFloor, destinationFloor int
		)
		// generate random number
		rand.Seed(time.Now().UnixNano())
		randomFloor = rand.Intn(maxFloor-minFloor) + minFloor

		fmt.Println("Selamat datang di Menara Sejati")
		
		// loop if name length invalid
		for !(util.IsNameLengthValid(name)) {
			fmt.Print("Tolong Masukkan Nama Anda [Minimal 5 Karakter] : ")
			fmt.Scanln(&name)
			if !(util.IsNameLengthValid(name)) {
				util.PrintNameLengthInvalid()
				continue
			}
		}
		
		fmt.Printf("Hello %s! Selamat datang!\n", name)
		fmt.Printf("Saat ini lift sedang berada dilantai : %d\n", randomFloor)

		// loop if current floor invalid
		for !(util.IsFloorValid(currentFloor)) { 
			fmt.Print("Silahkan masukkan lantai penjemputan (1-10) : ")
			fmt.Scanln(&currentFloor)
			if !(util.IsFloorValid(currentFloor)) {
				util.PrintFloorInvalid()
				continue
			}
		}

		// loop if destination floor invalid
		for !(util.IsFloorValid(destinationFloor)) {
			fmt.Print("Silahkan masukkan lantai tujuan (1-10) : ")
			fmt.Scanln(&destinationFloor)
			if !(util.IsFloorValid(destinationFloor)) {
				util.PrintFloorInvalid()
				continue
			}
		}

		if currentFloor == destinationFloor {
			fmt.Println("Anda sudah berada dilantai tujuan, lift tidak perlu menjemput")
		} else if randomFloor > currentFloor {
			util.PrintWaitingLiftForPickUp()

			// Loop when lift is going down
			util.LiftGoDown(randomFloor, currentFloor)

			util.PrintArriveAtCurrentFloor()

			if currentFloor > destinationFloor {
				// Loop when lift is going down
				util.LiftGoDown(currentFloor, destinationFloor)
			} else {
				// Loop when lift is going up
				util.LiftGoUp(currentFloor, destinationFloor)
			}

			util.PrintArriveAtDestinationFloor()
		} else {
			util.PrintWaitingLiftForPickUp()

			// Loop when lift is going up
			util.LiftGoUp(randomFloor, currentFloor)

			util.PrintArriveAtCurrentFloor()

			if currentFloor > destinationFloor {
				// Loop when lift is going down
				util.LiftGoDown(currentFloor, destinationFloor)
			} else {
				// Loop when lift is going up
				util.LiftGoUp(currentFloor, destinationFloor)
			}

			util.PrintArriveAtDestinationFloor()
		}

		var choice string
		fmt.Print("Ingin melanjutkan program? (y/n) : ")
		fmt.Scanln(&choice)
		if choice == "n" || choice == "N" {
			fmt.Println("Terima kasih sudah menggunakan program ini!")
			break
		}
		// continue program and clear screen
		fmt.Print("\033[H\033[2J")
	}
}
