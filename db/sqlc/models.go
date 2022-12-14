// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import ()

type Order struct {
	Uuid     int64 `json:"Uuid"`
	UserUuid int64 `json:"UserUuid"`
	Quantity int64 `json:"Quantity"`
}

type OrderProduct struct {
	OrderUuid   int64 `json:"OrderUuid"`
	ProductUuid int64 `json:"ProductUuid"`
}

type Product struct {
	Uuid        int64   `json:"Uuid"`
	Description string  `json:"Description"`
	Price       float32 `json:"Price"`
	InStock     int32   `json:"InStock"`
}

type User struct {
	Uuid           int64   `json:"Uuid"`
	FirstName      string  `json:"FirstName"`
	MiddleName     string  `json:"MiddleName"`
	LastName       string  `json:"LastName"`
	FullName       string  `json:"FullName"`
	Gender         string  `json:"Gender"`
	Age            int16   `json:"Age"`
	Balance        float32 `json:"Balance"`
	Username       string  `json:"Username"`
	HashedPassword string  `json:"HashedPassword"`
}

type UserToUser struct {
	FirstUserUuid  int64 `json:"FirstUserUuid"`
	SecondUserUuid int64 `json:"SecondUserUuid"`
}
