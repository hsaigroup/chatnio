package cli

import (
	"chat/auth"
	"chat/connection"
	"fmt"
	"strings"
)

func CreateInvitationCommand(args []string) {
	db := connection.ConnectMySQL()

	var (
		t     = GetArgString(args, 0)
		num   = GetArgInt(args, 1)
		quota = GetArgFloat32(args, 2)
	)

	resp, err := auth.GenerateInvitations(db, num, quota, t)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.Join(resp, "\n"))
}
