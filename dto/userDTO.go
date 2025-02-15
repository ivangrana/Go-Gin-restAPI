package dto

type UserDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserCreateDTO struct {
	Name     string `json:"name"`
	Password string `json:"Password"`
}

type UserUpdateDTO struct {
	Name     string `json:"name"`
	Password string `json:"Password"`
}

type UserLoginDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
