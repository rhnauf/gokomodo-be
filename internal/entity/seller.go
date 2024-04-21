package entity

type Seller struct {
	Id            int    `db:"id"`
	Email         string `db:"email"`
	Name          string `db:"name"`
	Password      string `db:"password"`
	AddressPickup string `db:"address_pickup"`
}
