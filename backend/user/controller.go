package user

import (
	"context"
	"strings"

	db "github.com/consoledot/notely/database"
	cryptolib "github.com/consoledot/notely/utils/crypto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUser(email string, name string, passwordHash string) *User {
	return &User{
		Email:        email,
		Name:         name,
		PasswordHash: passwordHash,
	}
}

func (user *User) CreateUser() (string, error) {
	coll := db.UsersCollection()

	result, err := coll.InsertOne(context.TODO(), user)
	// id := primitive.ObjectID(result.InsertedID)
	userId, _ := result.InsertedID.(primitive.ObjectID)
	strId := userId.String()
	hexID := strings.TrimPrefix(strings.TrimSuffix(strId, ")"), "ObjectID(")
	return hexID, err

}

func (user *User) DeleteUser() error {
	coll := db.UsersCollection()

	filter := bson.D{{Key: "_id", Value: user.Id}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	return err

}
func (user *User) GetUser(key string, value string) (User, error) {
	coll := db.UsersCollection()
	var n User
	filter := bson.D{{Key: key, Value: value}}
	err := coll.FindOne(context.TODO(), filter).Decode(&n)

	return n, err
}

func (user *User) DoesUserExit() bool {

	_, err := user.GetUser("email", user.Email)

	return err == nil
}

func (user *User) DoesPassWordMatch() bool {
	result, _ := user.GetUser("email", user.Email)

	return cryptolib.CompareHashWithText(result.PasswordHash, user.PasswordHash)

}
