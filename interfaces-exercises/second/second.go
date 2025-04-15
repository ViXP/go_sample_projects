package second

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) String() string {
	return fmt.Sprintf("User name: %s, email: %s\n", u.name, u.email)
}

func Run() {
	u := user{
		name:  "John Doe",
		email: "johndoe@example.com",
	}

	fmt.Println(u)
}
