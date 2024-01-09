package notes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/consoledot/notely/database"
	httplib "github.com/consoledot/notely/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}

	var result []Note
	cusror, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "Error getting notes", http.StatusNoContent)

	}

	if err = cusror.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		c.Response(false, nil, "Error getting notes", http.StatusNoContent)
	}

	c.Response(true, result, "get all notes", http.StatusOK)

}

func CreateNewNotes(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		fmt.Println(err)
		c.Response(false, nil, "error creating note", http.StatusBadRequest)
		return
	}

	result, err := coll.InsertOne(context.TODO(), note)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "error creating note", http.StatusBadRequest)
		return
	}

	c.Response(true, result, "note added successfully", http.StatusCreated)

}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}

	// Get variables
	id := c.GetParamsById(`id`)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "id is not a valid id", http.StatusBadRequest)
		return
	}

	note := bson.D{{Key: "_id", Value: objId}}

	result, err := coll.DeleteOne(context.TODO(), note)
	if err != nil || result.DeletedCount <= 0 {

		fmt.Println(err)
		c.Response(false, nil, "note not found", http.StatusNotFound)
		return
	}
	c.Response(true, nil, "note deleted", http.StatusOK)

}

func GetNote(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}

	// Get variables

	id := c.GetParamsById(`id`)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "id is not a valid id", http.StatusBadRequest)
		return
	}

	note := bson.D{{Key: "_id", Value: objId}}

	var result Note
	err = coll.FindOne(context.TODO(), note).Decode(&result)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "note not found", http.StatusNotFound)
		return
	}
	c.Response(true, result, "note found", http.StatusFound)

}

func EditNote(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}
	id := c.GetParamsById(`id`)

	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		fmt.Println(err)
		c.Response(false, nil, "error updating note", http.StatusCreated)
		return
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "id is not a valid id", http.StatusBadRequest)
		return
	}
	filter := bson.D{{Key: "_id", Value: objId}}

	_ = coll.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": note})

	c.Response(true, nil, "note updated successfully", http.StatusAccepted)

}
