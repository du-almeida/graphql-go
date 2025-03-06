package users

type users struct {
	ID       string `json: "id"`
	Username string `json: "name"`
	Password string `json: "password"`
}
