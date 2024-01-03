package helpers

const (
	ADMIN   = "admin"
	PESERTA = "peserta"
	NMAX    = 20
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
		{1, "Riyan", 50},
		{2, "Kinara", 100},
		{3, "Arsha", 100},
	}

	BankSoal = [NMAX]Soal{
		{1, "Berapa 1+1?", [4]string{"1", "2", "3", "4"}, "2"},
		{2, "Kapan Indonesia merdeka?", [4]string{"17 Januari 1900", "17 Februari 1902", "17 Agustus 1945", "31 Desember 2020"}, "17 Agustus 1945"},
		{3, "Siapakah capres untuk Pemilu 2024?", [4]string{"Anies, Prabowo, Ganjar", "Prabowo, Sutoyo, Anies", "Ganjar, Toyota, Sukijan", "Nurhadi, Aldo, Sujana"}, "Anies, Prabowo, Ganjar"},
		{4, "Siapa juara piala dunia 2022?", [4]string{"Zimbabwe", "Argentina", "Jerman", "Kanada"}, "Argentina"},
		{5, "Planet apa yang terdekat dengan Matahari?", [4]string{"Venus", "Bumi", "Merkurius", "Mars"}, "Merkurius"},
		{6, "Siapakah penemu teori relativitas?", [4]string{"Isaac Newton", "Albert Einstein", "Stephen Hawking", "Nikola Tesla"}, "Albert Einstein"},
		{7, "Ibu kota negara mana yang bernama Tokyo?", [4]string{"China", "Korea Selatan", "Jepang", "Indonesia"}, "Jepang"},
		{8, "Hewan apa yang dikenal sebagai raja hutan?", [4]string{"Harimau", "Serigala", "Singa", "Cheetah"}, "Singa"},
		{9, "Apa warna yang dihasilkan dari campuran warna merah dan kuning?", [4]string{"Hijau", "Ungu", "Oranye", "Biru"}, "Oranye"},
		{10, "Siapakah penulis novel 'Harry Potter'?", [4]string{"J.K. Rowling", "Stephen King", "George R.R. Martin", "J.R.R. Tolkien"}, "J.K. Rowling"},
		{11, "Apa unsur kimia dengan simbol 'O'?", [4]string{"Osmium", "Oxygen", "Orbit", "Opium"}, "Oxygen"},
		{12, "Negara mana yang memiliki ibu kota Bangkok?", [4]string{"Thailand", "Vietnam", "Malaysia", "Singapura"}, "Thailand"},
		{13, "Hewan apa yang dikenal sebagai hewan tercepat di darat?", [4]string{"Kuda", "Anjing", "Gepard", "Kelinci"}, "Gepard"},
		{14, "Apa bahasa resmi di Brasil?", [4]string{"Inggris", "Portugis", "Spanyol", "Prancis"}, "Portugis"},
		{15, "Monumen apa yang disebut juga sebagai 'The Bean' di Chicago?", [4]string{"Willis Tower", "Cloud Gate", "Millennium Park", "Navy Pier"}, "Cloud Gate"},
		{16, "Siapa pelukis di balik karya 'Mona Lisa'?", [4]string{"Vincent Van Gogh", "Pablo Picasso", "Leonardo da Vinci", "Michelangelo"}, "Leonardo da Vinci"},
		{17, "Negara manakah yang dikenal dengan sebutan 'Negeri Kiwi'?", [4]string{"Australia", "Selandia Baru", "Papua Nugini", "Fiji"}, "Selandia Baru"},
		{18, "Siapakah penemu lampu pijar?", [4]string{"Thomas Edison", "Nikola Tesla", "Benjamin Franklin", "Alexander Graham Bell"}, "Thomas Edison"},
		{19, "Di negara manakah Piramida Giza berada?", [4]string{"Meksiko", "Spanyol", "Mesir", "Peru"}, "Mesir"},
		{20, "Apa satelit alami Bumi?", [4]string{"Mars", "Venus", "Bulan", "Jupiter"}, "Bulan"},
	}
)
