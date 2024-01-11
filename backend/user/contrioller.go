package user

import (
	"context"

	db "github.com/consoledot/notely/database"
	"go.mongodb.org/mongo-driver/bson"
)

func NewUser(email string, name string, passwordHash string) *User {
	return &User{
		Email:        email,
		Name:         name,
		PasswordHash: passwordHash,
	}
}

func (user *User) CreateUser() (interface{}, error) {
	coll := db.UsersCollection()

	result, err := coll.InsertOne(context.TODO(), user)

	return result.InsertedID, err

}

func (user *User) DeleteUser() error {
	coll := db.UsersCollection()

	filter := bson.D{{Key: "_id", Value: user.Id}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	return err

}

func (user *User) DoesUserExit() bool {
	coll := db.UsersCollection()
	filter := bson.D{{Key: "email", Value: user.Email}}
	err := coll.FindOne(context.TODO(), filter).Err()

	return err != nil
}

func (user *User) GetUser() (User, error) {
	coll := db.UsersCollection()
	var n User
	filter := bson.D{{Key: "_id", Value: user.Id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&n)

	return n, err
}
