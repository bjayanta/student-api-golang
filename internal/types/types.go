package types

type Student struct {
	Id    int64  `json:"id"`
	Name  string `json:"name" validate:"required,min=3,max=20"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required"`
}