package util

import "fmt"

func PrintNameLengthInvalid() {
	fmt.Println("Nama Anda harus lebih dari 4 karakter!")
}

func PrintFloorInvalid() {
	fmt.Println("Lantai yang Anda masukkan harus diantara rentang 1-10!")
}

func PrintArriveAtCurrentFloor(){
	fmt.Println("Lift Sudah sampai dilokasi penjemputan, pintu lift terbuka")
	fmt.Println("Silahkan masuk, Anda akan diantar ke lantai tujuan")
}

func PrintArriveAtDestinationFloor(){
	fmt.Println("Lift Sudah sampai dilokasi tujuan, pintu lift terbuka")
	fmt.Println("Terima kasih, Anda sudah sampai di lantai tujuan")
}

func PrintWaitingLiftForPickUp(){
	fmt.Println("Mohon menunggu, lift akan segera bergerak menjemput Anda")
}
