package requests

type UserRequest struct {
	CreateUserRequest struct{
		Name string
		Email string
	}
}
