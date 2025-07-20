package authorize

type In struct {
	Email    string
	Password string
}

type Out struct {
	Token string
}
