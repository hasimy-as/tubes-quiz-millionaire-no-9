package peserta

import (
	"fmt"

	common "github.com/hasimy-as/tubes-quiz-millionaire-no-9/bin/helpers"
)

// Diakses oleh main.go
func MainMenuPeserta(loggedIn ...bool) {
	if loggedIn != nil {
		tanyaJawab(0)
	} else {
		var answer string

		fmt.Print("Apakah anda sudah mendaftar? (y/n): ")
		fmt.Scan(&answer)

		if answer == "y" {
			loginPeserta()
		} else {
			fmt.Println("Silakan mendaftar terlebih dahulu.")
			registerPeserta()
		}
	}
}

/**
 * Fungsi-fungsi yang digunakan oleh peserta.
 */

func loginPeserta() {
	var (
		id   int
		nama string
	)

	fmt.Print("\nMasukkan ID dan nama anda: ")
	fmt.Scan(&id, &nama)

	var found bool = false

	for _, peserta := range common.DaftarPeserta {
		if peserta.Id == id && peserta.Nama == nama {
			found = true
			break
		}
	}

	if found {
		fmt.Printf("ID %d dengan nama %s ditemukan dalam daftar peserta.\n", id, nama)
		fmt.Printf("Selamat datang %s!\n", common.DaftarPeserta[id-1].Nama)
		var loggedIn bool = true

		MainMenuPeserta(loggedIn)
	} else {
		fmt.Printf("ID %d dengan nama %s tidak ditemukan dalam daftar peserta.\n", id, nama)
		return
	}
}

func registerPeserta() {
	var jumlahPeserta int = pesertaCheck()

	if jumlahPeserta >= common.NMAX {
		fmt.Println("Kapasitas peserta sudah penuh.")
		return
	}

	var nama string
	fmt.Print("\nMasukkan nama Anda: ")
	fmt.Scan(&nama)

	common.DaftarPeserta[jumlahPeserta] = common.Peserta{
		Id:   jumlahPeserta + 1,
		Nama: nama,
		Skor: 0,
	}

	jumlahPeserta++

	fmt.Printf("Peserta %s berhasil mendaftar dengan ID %d, silakan login\n", nama, jumlahPeserta)
	loginPeserta()
}

func pesertaCheck() int {
	var count int

	for _, peserta := range common.DaftarPeserta {
		if peserta.Id != 0 {
			count++
		}
	}

	return count
}

// Mesti disesuaikan lagi buat nampilin soal acak, lebih dari 1.
func tanyaJawab(soalIndex int) {
	if soalIndex < 0 || soalIndex >= common.NMAX || common.BankSoal[soalIndex].Id == 0 {
		fmt.Println("Pertanyaan tidak ditemukan.")
		return
	}

	var soal = common.BankSoal[soalIndex]
	fmt.Printf("\nPertanyaan: %s\n", soal.Pertanyaan)

	var options = []string{"a", "b", "c", "d"}
	for i, pilihan := range soal.Pilihan {
		fmt.Printf("%s. %s\n", options[i], pilihan)
	}

	var jawabanUser string
	fmt.Print("\nMasukkan jawaban Anda (a, b, c, atau d): ")
	fmt.Scan(&jawabanUser)

	var userAnswerIndex = -1
	for i, opt := range options {
		if jawabanUser == opt {
			userAnswerIndex = i
			break
		}
	}

	if userAnswerIndex == -1 {
		fmt.Println("Jawaban tidak valid.")
		return
	}

	fmt.Printf("\nJawaban anda: %s\n", soal.Pilihan[userAnswerIndex])

	if soal.Pilihan[userAnswerIndex] == soal.KunciJawaban {
		fmt.Println("Jawaban anda benar!")
	} else {
		fmt.Println("Jawaban anda salah.")
		fmt.Printf("Jawaban yang benar: %s\n", soal.KunciJawaban)
	}
}
