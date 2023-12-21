package helpers

// Helper untuk memilih role.
func RoleCheck(role int) string {
	if role == 1 {
		return ADMIN
	}

	return PESERTA
}

// Helper untuk mengacak soal.
func AcakIndeks(index, totalSoal int) int {
	return (index*index + 1) % totalSoal

}
