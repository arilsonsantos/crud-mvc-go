package domain

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func NewUserDomain(name, email, password string, age int8) UserDomainInterface {
	return &userDomain{
		name:     name,
		email:    email,
		password: password,
		age:      age,
	}
}

func (u *userDomain) GetEmail() string {
	return u.email
}

func (u *userDomain) GetPassword() string {
	return u.password
}

func (u *userDomain) GetName() string {
	return u.name
}

func (u *userDomain) GetAge() int8 {
	return u.age
}

func (u *userDomain) SetID(id string) {
	u.id = id
}

func (u *userDomain) GetID() string {
	return u.id
}
