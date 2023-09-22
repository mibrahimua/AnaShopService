package repository

import (
	"database/sql"
)

type ShopRepository struct {
	db *sql.DB
}

func NewShopRepository(db *sql.DB) *ShopRepository {
	return &ShopRepository{db}
}

func (u *ShopRepository) PaidCartItems(cartsId int) error {
	query := "UPDATE cart SET is_paid = TRUE WHERE id = $1"
	_, err := u.db.Query(query, cartsId)
	if err != nil {
		return err
	}

	return nil
}
