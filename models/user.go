package models

type User struct {
	UID            string   `json:"uid"`
	FirstName      *string  `json:"first_name"`
	SecondName     *string  `json:"second_name"`
	LastName       *string  `json:"last_name"`
	SecondLastName *string  `json:"second_last_name"`
	Gender         *string  `json:"gender"`
	Nationality    *string  `json:"nationality"`
	NationalID     *string  `json:"national_id"`
	Email          *string  `json:"email"`
	Type           *string  `json:"type"`
	Roles          []string `json:"roles"`
	FullName       string   `json:"full_name"`
}

type NonPublicUser struct {
	UID            string  `json:"uid"`
	FirstName      *string `json:"first_name"`
	SecondName     *string `json:"second_name"`
	LastName       *string `json:"last_name"`
	SecondLastName *string `json:"second_last_name"`
	Gender         *string `json:"gender"`
	Nationality    *string `json:"nationality"`
	FullName       string  `json:"full_name"`
	Cellphone      *string `json:"cellphone"`
}

func (u *User) IsProfileCompleted() bool {
	return u.FirstName != nil && u.LastName != nil && u.Email != nil
}

func (u *User) GetFullName() string {
	return *u.FirstName + " " + *u.LastName
}

func (u *User) GetFullNameWithSecondName() string {
	return *u.FirstName + " " + *u.SecondName + " " + *u.LastName
}

func (u *User) GetFullNameWithSecondLastName() string {
	return *u.FirstName + " " + *u.SecondName + " " + *u.LastName + " " + *u.SecondLastName
}

func (u *User) ToNonPublicUser() *NonPublicUser {
	return &NonPublicUser{
		UID:            u.UID,
		FirstName:      u.FirstName,
		SecondName:     u.SecondName,
		LastName:       u.LastName,
		SecondLastName: u.SecondLastName,
		FullName:       u.FullName,
	}
}
