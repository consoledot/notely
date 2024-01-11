package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string             `bson:"title" json:"title"`
	Author    string             `bson:"author" json:"author"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}
