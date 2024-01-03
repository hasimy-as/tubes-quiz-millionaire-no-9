package peserta

import (
	"fmt"

	common "github.com/hasimy-as/tubes-quiz-millionaire-no-9/bin/helpers"
)

// Diakses oleh main.go
func MainMenuPeserta(loggedIn ...bool) {
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

/**
 * Fungsi-fungsi yang digunakan oleh peserta.
 */

func loginPeserta() {
	var (
		id               int
		nama             string
		found            bool = false
		pesertaTerdaftar common.Peserta
	)

	fmt.Print("\nMasukkan ID dan nama anda: ")
	fmt.Scan(&id, &nama)

	for _, peserta := range common.DaftarPeserta {
		if peserta.Id == id && peserta.Nama == nama {
			pesertaTerdaftar = peserta
			found = true
		}
	}

	if found {
		fmt.Printf("ID %d dengan nama %s ditemukan dalam daftar peserta.\n", id, nama)
		fmt.Printf("Selamat datang %s!\n", pesertaTerdaftar.Nama)

		bridgeToTanyaJawab(pesertaTerdaftar)
	} else {
		fmt.Printf("ID %d dengan nama %s tidak ditemukan dalam daftar peserta.\n", id, nama)
		return
	}
}

func registerPeserta() {
	var (
		nama          string
		jumlahPeserta int = pesertaCheck()
	)

	if jumlahPeserta >= common.NMAX {
		fmt.Println("Kapasitas peserta sudah penuh.")
		return
	}

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

func bridgeToTanyaJawab(pesertaTerdaftar common.Peserta) {
	var jumlahSoal int
	fmt.Printf("\nSkor terakhir anda adalah %d\n", pesertaTerdaftar.Skor)

	fmt.Print("\nBerapa jumlah soal yang anda inginkan? Silakan masukan: ")
	fmt.Scan(&jumlahSoal)

	tanyaJawab(jumlahSoal, pesertaTerdaftar)
}

// Perlu penyesuaian buat randomize soal quiz yang tampil, dan
// jangan nampilin soal yang udah sebelumnya ditampilkan.
func tanyaJawab(N int, pesertaTerdaftar common.Peserta) {
	var (
		n         int
		randSlice []int
		pilihan   string
		soal      common.Soal

		selesaiJawab  = false
		totalSoalQuiz = len(common.BankSoal)
	)

	if N < 1 || N > totalSoalQuiz {
		fmt.Printf("Jumlah soal tidak valid. Maksimal soal adalah %d soal.\n", totalSoalQuiz)
		return
	}

	randSlice = common.SimpleShuffle(N, totalSoalQuiz)

	for n = 0; n < N; n++ {
		soal = common.BankSoal[randSlice[n]]

		if soal.Id == 0 {
			fmt.Printf("Soal no.%d tidak tersedia.\n", randSlice[n]+1)
			continue
		}

		fmt.Printf("\nPertanyaan %d: %s\n", n+1, soal.Pertanyaan)
		var options = []string{"a", "b", "c", "d"}
		for i, pilihan := range soal.Pilihan {
			fmt.Printf("%s. %s\n", options[i], pilihan)
		}

		var jawabanUser string
		fmt.Print("\nMasukkan jawaban Anda (a, b, c, atau d): ")
		fmt.Scan(&jawabanUser)

		var userAnswerIndex = -1
		var validAnswerFound = false

		for i, opt := range options {
			if jawabanUser == opt {
				userAnswerIndex = i
				validAnswerFound = true
			}
		}

		if !validAnswerFound {
			fmt.Println("Jawaban tidak valid.")
		} else {
			fmt.Printf("\nJawaban anda: %s\n", soal.Pilihan[userAnswerIndex])

			if soal.Pilihan[userAnswerIndex] == soal.KunciJawaban {
				fmt.Println("Jawaban anda benar!")
				pesertaTerdaftar.Skor = pesertaTerdaftar.Skor + 10

				fmt.Printf("Skor anda: %d\n", pesertaTerdaftar.Skor)
			} else {
				fmt.Println("Jawaban anda salah.")
				fmt.Printf("Jawaban yang benar: %s\n", soal.KunciJawaban)

				// Supaya skor tidak negatif.
				if pesertaTerdaftar.Skor > 0 {
					pesertaTerdaftar.Skor = pesertaTerdaftar.Skor - 10
				}

				fmt.Printf("Skor anda: %d\n", pesertaTerdaftar.Skor)
			}
		}
	}

	for !selesaiJawab {
		fmt.Print("\nApakah anda ingin mengulang kuis? (y/n): ")
		fmt.Scan(&pilihan)

		if pilihan == "n" {
			fmt.Println("Terima kasih telah berpartisipasi dalam kuis millionaire!")
			selesaiJawab = true
		} else if pilihan == "y" {
			bridgeToTanyaJawab(pesertaTerdaftar)
			selesaiJawab = true
		} else {
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
