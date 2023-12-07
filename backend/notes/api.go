package notes

import "time"

type Note struct {
	Title string 
	author string 
	createdAt time.Time 
}

var notes []Note
func AddNote (note Note){
	notes = append(notes, note)
}

func GetNote (index int32 ) Note{
	return notes[index]
}

// func DeleteNote(index int32){
// 	notes.
// }