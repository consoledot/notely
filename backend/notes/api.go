package notes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/consoledot/notely/database"
	"github.com/consoledot/notely/user"
	cryptolib "github.com/consoledot/notely/utils/crypto"
	httplib "github.com/consoledot/notely/utils/httplib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}
	tokenResponse := r.Context().Value(cryptolib.TokenResponse{}).(cryptolib.TokenResponse)
	var user user.User
	userResponse, _ := user.GetUser("email", tokenResponse.Email)
	var result []Note
	filter := bson.M{"created_by": userResponse.Id}
	cusror, err := coll.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "Error getting notes", http.StatusNoContent, nil)

	}

	if err = cusror.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		c.Response(false, nil, "Error getting notes", http.StatusNoContent, nil)
	}

	c.Response(true, result, "get all notes", http.StatusOK, nil)

}

func CreateNewNotes(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}
	tokenResponse := r.Context().Value(cryptolib.TokenResponse{}).(cryptolib.TokenResponse)
	var note Note
	var user user.User
	result, err := user.GetUser("email", tokenResponse.Email)
	if err != nil {
		c.Response(false, nil, "No user found", http.StatusForbidden, nil)
		return
	}

	note.CreatedBy = result.Id
	fmt.Println("kjhg", note)
	if err := c.GetJSONfromRequestBody(&note); err != nil {
		fmt.Println(err)
		c.Response(false, nil, "error creating note", http.StatusBadRequest, nil)
		return
	}

	_, err = coll.InsertOne(context.TODO(), note)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "error creating note", http.StatusBadRequest, nil)
		return
	}

	c.Response(true, nil, "note added successfully", http.StatusCreated, nil)

}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}

	// Get variables
	id := c.GetParamsById(`id`)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "id is not a valid id", http.StatusBadRequest, nil)
		return
	}

	note := bson.D{{Key: "_id", Value: objId}}

	result, err := coll.DeleteOne(context.TODO(), note)
	if err != nil || result.DeletedCount <= 0 {

		fmt.Println(err)
		c.Response(false, nil, "note not found", http.StatusNotFound, nil)
		return
	}
	c.Response(true, nil, "note deleted", http.StatusOK, nil)

}

func GetNote(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}

	// Get variables

	id := c.GetParamsById(`id`)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "id is not a valid id", http.StatusBadRequest, nil)
		return
	}

	note := bson.D{{Key: "_id", Value: objId}}

	var result Note
	err = coll.FindOne(context.TODO(), note).Decode(&result)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "note not found", http.StatusNotFound, nil)
		return
	}
	c.Response(true, result, "note found", http.StatusFound, nil)

}

func EditNote(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}
	id := c.GetParamsById(`id`)

	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		fmt.Println(err)
		c.Response(false, nil, "error updating note", http.StatusCreated, nil)
		return
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "id is not a valid id", http.StatusBadRequest, nil)
		return
	}
	filter := bson.D{{Key: "_id", Value: objId}}

	_ = coll.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": note})

	c.Response(true, nil, "note updated successfully", http.StatusAccepted, nil)

}
