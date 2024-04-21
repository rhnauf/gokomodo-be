package entity

type Buyer struct {
	Id          int    `db:"id"`
	Email       string `db:"email"`
	Name        string `db:"name"`
	Password    string `db:"password"`
	AddressSend string `db:"address_send"`
}
