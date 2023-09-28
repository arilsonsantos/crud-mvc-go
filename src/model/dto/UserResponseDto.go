package dto

type UserResponseDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int8   `json:"age"`
}
