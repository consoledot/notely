package notes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/consoledot/notely/database"
	httplib "github.com/consoledot/notely/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}

	var result []Note
	cusror, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Panic(err)

	}

	if err = cusror.All(context.TODO(), &result); err != nil {
		log.Panic(err)
	}

	c.Response(true, result, "get all notes", http.StatusOK)

}

func CreateNewNotes(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {

		c.Response(false, nil, "error creating note", http.StatusBadRequest)
		return
	}

	result, err := coll.InsertOne(context.TODO(), note)
	if err != nil {
		c.Response(false, nil, "error creating note", http.StatusBadRequest)
		return
	}

	c.Response(true, result, "note added successfully", http.StatusCreated)

}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	c := httplib.C{W: w, R: r}

	// Get variables
	vars := mux.Vars(r)
	id := vars["id"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Id is not a primitive id: ", err)
	}

	note := bson.D{{Key: "_id", Value: objId}}

	result, err := coll.DeleteOne(context.TODO(), note)
	if err != nil || result.DeletedCount <= 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Note not found"))
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
		fmt.Println("Id is not a primitive id: ", err)
		w.Write([]byte("Not a valid ID"))
		return
	}

	note := bson.D{{Key: "_id", Value: objId}}

	var result Note
	err = coll.FindOne(context.TODO(), note).Decode(&result)
	if err != nil {
		fmt.Println("Not found: ", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Note not found"))
		return
	}
	c.Response(true, result, "note found", http.StatusFound)

}
