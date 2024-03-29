package user

import "github.com/google/uuid"

type UserModel struct {
	ID       uuid.UUID `bson:"_id"`
	Username string    `bson:"username"`
	Password string    `bson:"password"`
}
