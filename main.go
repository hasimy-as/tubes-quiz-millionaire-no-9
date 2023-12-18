package main

import "fmt"

const (
	ADMIN   = "admin"
	PESERTA = "peserta"
	NMAX    = 100
)

type (
	Soal struct {
		Id           int
		Pertanyaan   string
		Pilihan      [4]string
		KunciJawaban string
	}

	Peserta struct {
		Id   int
		Nama string
		Skor int
	}
)

var (
	// Arrays
	DaftarPeserta = [NMAX]Peserta{
		{1, "Riyan", 110},
		{2, "Kinara", 190},
		{3, "Arsha", 90},
	}

	BankSoal = [NMAX]Soal{
		{1, "Berapa 1+1?", [4]string{"1", "2", "3", "4"}, "2"},
		{2, "Kapan Indonesia merdeka?", [4]string{"17 Januari 1900", "17 Februari 1902", "17 Agustus 1945", "31 Desember 2020"}, "17 Agustus 1945"},
		{3, "Siapakah capres untuk Pemilu 2024?", [4]string{"Anies, Prabowo, Ganjar", "Prabowo, Sutoyo, Anies", "Ganjar, Toyota, Sukijan", "Nurhadi, Aldo, Sujana"}, "Anies, Prabowo, Ganjar"},
		{4, "Siapa juara piala dunia 2022?", [4]string{"Zimbabwe", "Argentina", "Jerman", "Kanada"}, "Argentina"},
	}
)

// Helper untuk memilih role.
func roleCheck(role int) string {
	if role == 1 {
		return ADMIN
	}

	return PESERTA
}

func pesertaCheck() int {
	var count int

	for _, peserta := range DaftarPeserta {
		if peserta.Id != 0 {
			count++
		}
	}

	return count
}

func login() {
	var (
		id   int
		nama string
	)

	fmt.Print("\nMasukkan ID dan nama anda: ")
	fmt.Scan(&id, &nama)

	var found bool = false

	for _, peserta := range DaftarPeserta {
		if peserta.Id == id && peserta.Nama == nama {
			found = true
			break
		}
	}

	if found {
		fmt.Printf("ID %d dengan nama %s ditemukan dalam daftar peserta.\n", id, nama)
		fmt.Printf("Selamat datang %s!\n", DaftarPeserta[id-1].Nama)
	} else {
		fmt.Printf("ID %d dengan nama %s tidak ditemukan dalam daftar peserta.\n", id, nama)
		return
	}
}

// Hal yang dapat dilakukan admin, i.e. menambahkan soal, dsb.
func adminCommands() string {
	return "Disini admin melakukan tambah, edit, hapus soal dan kunci jawaban"
}

func mainMenu() {
	var answer string

	fmt.Print("Apakah anda sudah mendaftar? (y/n): ")
	fmt.Scan(&answer)

	if answer == "y" {
		login()
	} else {
		fmt.Println("Silakan mendaftar terlebih dahulu.")
		registerPeserta()
	}
}

func registerPeserta() {
	var jumlahPeserta int = pesertaCheck()

	if jumlahPeserta >= NMAX {
		fmt.Println("Kapasitas peserta sudah penuh.")
		return
	}

	var nama string
	fmt.Print("\nMasukkan nama Anda: ")
	fmt.Scan(&nama)

	DaftarPeserta[jumlahPeserta] = Peserta{
		Id:   jumlahPeserta + 1,
		Nama: nama,
		Skor: 0,
	}

	jumlahPeserta++

	fmt.Printf("Peserta %s berhasil mendaftar dengan ID %d, silakan login\n", nama, jumlahPeserta)
	login()
}

func main() {
	var (
		userType int
		role     string
	)

	fmt.Print("Silakan masukan role anda (1 untuk admin, 2 untuk peserta): ")

	// Masuk sebagai user admin/peserta.
	fmt.Scan(&userType)
	role = roleCheck(userType)

	if role == ADMIN {
		adminCommands()
	} else {
		mainMenu()
	}
}
