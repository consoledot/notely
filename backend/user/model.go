package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		Email        string             `json:"email" bson:"email"`
		Name         string             `json:"name" bson:"name"`
		PasswordHash string             `json:"password_hash" bson:"password_hash"`
		CreatedAt    time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	}
)
