package admin

import "fmt"

// Diakses oleh main.go
func AdminCommands() {
	var choice int
	fmt.Println("1. Tambah Soal")
	fmt.Println("2. Edit Soal")
	fmt.Println("3. Hapus Soal")
	fmt.Print("Masukkan nomor untuk aksi yang ingin Anda lakukan: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		tambahSoal()
	case 2:
		editSoal()
	case 3:
		hapusSoal()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

/**
 * Fungsi-fungsi yang digunakan oleh admin.
 * Hal yang dapat dilakukan admin, i.e. menambahkan soal, dsb.
 */

// Kemampuan Admin untuk Tambah Soal ke dalam array
func tambahSoal() {
	fmt.Print("tambahSoal")
}

// Kemampuan Admin untuk edit Soal dalam array
func editSoal() {
	fmt.Print("editsoal")
}

// Kemampuan Admin untuk Hapus Soal dari array
func hapusSoal() {
	fmt.Print("hapussoal")

}
