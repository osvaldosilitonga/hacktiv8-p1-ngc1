package main

import (
	"fmt"
)

func main() {
	// NG Challange 1 : Variabel 1
	// Soal A
	myNum := 50
	fmt.Println("myNum =", myNum)
	// ------------------------------------

	// Soal B
	var myNum2 float32 = 51
	fmt.Println("myNum2 = ", myNum2)
	// --------------------------------

	// Soal C
	var myNumStr string = "50"
	fmt.Println("myNumStr = ", myNumStr)
	// ------------------------------------------

	// NG Challange 1 : Variabel 2
	x := 5
	y := 10
	var z = x + y
	fmt.Println("Hasil x + y = ", z)
	// ----------------------------------------

	// NG Challange 1 : CLI -> Input Nama
	var nama string
	fmt.Print("Masukkan Nama : ")
	fmt.Scanln(&nama)
	fmt.Println("Helo", nama)
	// ---------------------------------

	// NG Challange 1 : Slice
	people := []string{
		"Walt", "Jesse", "Skyler", "Saul",
	}

	// Hitung panjang slice
	fmt.Println("Panjang slice people :", len(people))
	// -------------------------------------------------

	// Tambah nilai ke slice
	people = append(people, "Hank", "Marie")
	// ---------------------------------------------

	// Panjang slice baru dan isi slice baru
	fmt.Println("Panjang slice baru :", len(people))
	fmt.Println("People baru :", people)
	// ------------------------------------------------

	// Array & Slice 2
	person := [3]map[string]string{
		{
			"name":   "Hank",
			"gender": "M",
		},
		{
			"name":   "Hank",
			"gender": "M",
		},
		{
			"name":   "Hank",
			"gender": "M",
		},
	}

	fmt.Println("Person :", person)

	// Mengganti gender jadi "F" dan menambahkan "Mrs"
	for i := 0; i < len(person); i++ {
		if person[i]["gender"] == "F" {
			person[i]["name"] = "Mrs. " + person[i]["name"]
		} else {
			person[i]["name"] = "Mr. " + person[i]["name"]
		}
	}
	fmt.Println("Person Baru :", person)
	// --------------------------------------------------
}
