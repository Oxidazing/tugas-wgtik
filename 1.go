package main

import "fmt"

func main() {
	var uts, uas, kuis, prak int
	var total float64

	fmt.Scanln(&uts, &uas, &kuis, &prak)

	uts = uts * 35 / 100
	uas = uas * 35 / 100
	kuis = kuis * 20 / 100
	prak = prak * 10 / 100

	total = float64(uas) + float64(uts) + float64(kuis) + float64(prak)

	if 0 <= total && total <= 40 {
		fmt.Println("Nilai akhir: ", total)
		fmt.Println("Indeks mutu: E")
	}
	if 40 < total && total <= 50 {
		fmt.Println("Nilai akhir: ", total)
		fmt.Println("Indeks mutu: D")
	}
	if 50 < total && total <= 70 {
		fmt.Println("Nilai akhir: ", total)
		fmt.Println("Indeks mutu: C")
	}
	if 70 < total && total <= 85 {
		fmt.Println("Nilai akhir: ", total)
		fmt.Println("Indeks mutu: B")
	}
	if 85 < total && total <= 100 {
		fmt.Println("Nilai akhir: ", total)
		fmt.Println("Indeks mutu: A")
	}
	if total > 100 {
		fmt.Println("NILAI TIDAK VALID")
	}
}
