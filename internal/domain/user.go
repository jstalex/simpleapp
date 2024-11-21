package domain

type User struct {
	Name       string `json:"name,omitempty"`
	Surname    string `json:"surname,omitempty"`
	Patronymic string `json:"patronymic,omitempty"`
	Email      string `json:"email,omitempty"`
}
