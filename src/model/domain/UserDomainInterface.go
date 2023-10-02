package domain

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	SetID(string2 string)
	GetID() string
	EncryptPassword()
}
