package models

import "time"

type ContactOne struct {
	Name string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required,max=20,min=1"`
	Phone string `json:"phone" validate:"required,len=10"`
	Email string `json:"email" validate:"required,email"`
	Time time.Time `json:"time"`
}


type Contacts struct{
    Name string "bson:`name`"
    Content string "bson:`content`"
    Phone string "bson:`phone`"
    Email string "bson:`email`"
    Time time.Time "bson:`time`"
}
