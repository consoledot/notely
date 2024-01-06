package notes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/consoledot/notely/database"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	w.Header().Set("Content-Type", "application/json")
	// id := r.URL.Query()["id"]
	// if id != nil {
	// 	index, err := strconv.Atoi(id[0])
	// 	if err == nil && index < len(GetAllNotes()) {
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode(GetAllNotes()[index])
	// 	} else {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		http.Error(w, "Note not available or doesn't exist", http.StatusBadRequest)
	// 	}
	// 	return
	// }

	var result []Note
	cusror, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Panic(err)

	}

	if err = cusror.All(context.TODO(), &result); err != nil {
		log.Panic(err)
	}

	fmt.Println(result)
	// fmt.Println(cusror)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func CreateNewNotes(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	w.Header().Set("Content-Type", "application/json")
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := coll.InsertOne(context.TODO(), note)
	if err != nil {
		log.Panic("Error creating note ", err)
	}
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// AddNote(note)
	fmt.Println(result)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Note added successfully"))
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Note deleted"))

}

func GetNote(w http.ResponseWriter, r *http.Request) {
	var coll = db.NotesCollection()
	// Get variables
	vars := mux.Vars(r)
	id := vars["id"]
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
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(result)

}
