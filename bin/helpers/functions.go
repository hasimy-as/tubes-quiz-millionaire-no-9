package helpers

// Helper untuk memilih role.
func RoleCheck(role int) string {
	if role == 1 {
		return ADMIN
	}

	return PESERTA
}

func SimpleShuffle(N, totalSoalQuiz int) []int {
	var (
		i       int
		numbers = make([]int, totalSoalQuiz)
	)

	for i = range numbers {
		numbers[i] = i
	}

	for i = range numbers {
		swapIndex := i + ((i*73 + 41) % 100 % (totalSoalQuiz - i))
		numbers[i], numbers[swapIndex] = numbers[swapIndex], numbers[i]
	}

	return numbers[:N]
}
