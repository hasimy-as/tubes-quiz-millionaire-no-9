package helpers

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
	DaftarPeserta = [100]Peserta{
		{1, "Riyan", 110},
		{2, "Kinara", 190},
		{3, "Arsha", 90},
	}

	BankSoal = [100]Soal{
		{1, "Berapa 1+1?", [4]string{"1", "2", "3", "4"}, "2"},
		{2, "Kapan Indonesia merdeka?", [4]string{"17 Januari 1900", "17 Februari 1902", "17 Agustus 1945", "31 Desember 2020"}, "17 Agustus 1945"},
		{3, "Siapakah capres untuk Pemilu 2024?", [4]string{"Anies, Prabowo, Ganjar", "Prabowo, Sutoyo, Anies", "Ganjar, Toyota, Sukijan", "Nurhadi, Aldo, Sujana"}, "Anies, Prabowo, Ganjar"},
		{4, "Siapa juara piala dunia 2022?", [4]string{"Zimbabwe", "Argentina", "Jerman", "Kanada"}, "Argentina"},
	}
)
