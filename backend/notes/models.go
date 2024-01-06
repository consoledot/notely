package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Author    string             `bson:"author"`
	CreatedAt time.Time          `bson:"created_at"`
}

var notes []Note

func AddNote(note Note) {
	notes = append(notes, note)
}

func GetNote(index int32) Note {
	return notes[index]
}

func GetAllNotes() []Note {
	if len(notes) <= 0 {
		return []Note{}
	}

	return notes
}
