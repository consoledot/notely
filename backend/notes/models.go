package notes

import "time"

type Note struct {
	Title string 
	Author string 
	CreatedAt time.Time 
}

var notes []Note
func AddNote (note Note){
	notes = append(notes, note)
}

func GetNote (index int32 ) Note{
	return notes[index]
}

func GetAllNotes () []Note{
	if len(notes) <= 0 {
		return []Note{}
	}

	return notes
}

// func DeleteNote(index int32){
// 	notes.
// }ç