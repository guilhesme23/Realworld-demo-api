package schemas

type User struct {
	Username string `json:"username" unique:"true"`
	Email    string `json:"email" unique:"true"`
	Password string `json:"password"`
}
