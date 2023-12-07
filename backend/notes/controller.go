package notes

import (
	"encoding/json"
	"net/http"
	"strconv"
)


func NotesController (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet{
		id := r.URL.Query()["id"]
		if id != nil{
			index, err := strconv.Atoi(id[0])
			if err == nil && index < len(GetAllNotes()){
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(GetAllNotes()[index])
			}else{
				w.WriteHeader(http.StatusBadRequest)
				http.Error(w, "Note not available or doesn't exist", http.StatusBadRequest)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(GetAllNotes())
		return 
	}

	if r.Method == http.MethodPost{
		var note Note
		err := json.NewDecoder(r.Body).Decode(&note)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		AddNote(note)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Note added successfully"))
		return
	}
	http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	
}