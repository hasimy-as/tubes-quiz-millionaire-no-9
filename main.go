package main

import "fmt"

const (
	ADMIN   = "admin"
	PESERTA = "peserta"
	NMAX    = 100
)

type Soal struct {
	Id           int
	Pertanyaan   string
	Pilihan      [4]string
	KunciJawaban string
}

type Peserta struct {
	Id   int
	Nama string
	Skor int
}

var BankSoal = [NMAX]Soal{
	{1, "Berapa 1+1?", [4]string{"1", "2", "3", "4"}, "2"},
	{2, "Kapan Indonesia merdeka?", [4]string{"17 Januari 1900", "17 Februari 1902", "17 Agustus 1945", "31 Desember 2020"}, "17 Agustus 1945"},
	{3, "Siapakah capres untuk Pemilu 2024?", [4]string{"Anies, Prabowo, Ganjar", "Prabowo, Sutoyo, Anies", "Ganjar, Toyota, Sukijan", "Nurhadi, Aldo, Sujana"}, "2"},
	{4, "Berapa 1+1?", [4]string{"1", "2", "3", "4"}, "2"},
}

var DaftarPeserta = [NMAX]Peserta{
	{1, "Riyan", 110},
	{2, "Kinara", 190},
	{3, "Arsha", 90},
}

// Login untuk role.
func login(role int) string {
	if role == 1 {
		return "admin"
	}

	return "peserta"
}

// Hal yang dapat dilakukan admin, i.e. menambahkan soal, dsb.
func adminCommands() string {
	return "Disini admin melakukan tambah, edit, hapus soal dan kunci jawaban"
}

func mainMenu() string {
	return "Disini peserta melakukan pilihan soal dan melihat skor"
}

func main() {
	var (
		userType int
		role     string
	)

	fmt.Print("Silakan masukan role anda (1 untuk admin, 2 untuk peserta): ")

	// Masuk sebagai user admin/peserta.
	fmt.Scan(&userType)
	role = login(userType)
	fmt.Printf("Anda masuk sebagai %s\n", role)

	if role == ADMIN {
		adminCommands()
	} else {
		mainMenu()
	}
}
