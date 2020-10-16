package models

import "time"

type Contender struct{
	Surname 		string
	Lastname 		string
	Email			string
	Birthdate		time.Time
	Password		string
	OwnDescription 	string
}

func NewContender(surname string, lastname string, email string, birthdate time.Time, password string, ownDescription string) *Contender {
	return &Contender{Surname: surname, Lastname: lastname, Email: email, Birthdate: birthdate, Password: password, OwnDescription: ownDescription}
}

