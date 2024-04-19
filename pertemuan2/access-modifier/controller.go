package controller

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ReadUserList() users []user {
	users = make([]user, 2)

	users[0] = user{
		ID:    1,
		Name:  "Ksatria Faikar Nasywaan",
		Email: "faikarksatria@gmail.com",
	}

	users[1] = user{
		ID:   2,
		Name: "M. Alirizki",
		Email: "malirizki@gmail.com",
	}

	return
}


