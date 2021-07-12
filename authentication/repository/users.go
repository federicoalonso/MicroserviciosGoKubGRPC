package repository

import (
	"microservicios/db"

	"gopkg.in/mgo.v2"
)

const UsersCollection = "users"

type UsersRepository interface{}

type usersRepository struct {
	c *mgo.Collection
}

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{c: conn.DB().C(UsersCollection)}
}

func (r *usersRepository) Save() {}
