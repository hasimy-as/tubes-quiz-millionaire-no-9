package main

import (
	"fmt"

	helpers "github.com/hasimy-as/tubes-quiz-millionaire-no-9/bin/helpers"

	admin "github.com/hasimy-as/tubes-quiz-millionaire-no-9/bin/modules/admin"
	peserta "github.com/hasimy-as/tubes-quiz-millionaire-no-9/bin/modules/peserta"
)

func main() {
	var (
		userType int
		role     string
	)

	fmt.Print("Silakan masukan role anda (1 untuk admin, 2 untuk peserta): ")

	// Masuk sebagai user admin/peserta.
	fmt.Scan(&userType)
	role = helpers.RoleCheck(userType)

	if role == helpers.ADMIN {
		admin.AdminCommands()
	} else {
		peserta.MainMenuPeserta()
	}
}
