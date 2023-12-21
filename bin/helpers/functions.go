package helpers

// Helper untuk memilih role.
func RoleCheck(role int) string {
	if role == 1 {
		return ADMIN
	}

	return PESERTA
}
