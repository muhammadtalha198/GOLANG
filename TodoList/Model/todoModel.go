package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task      string             `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool               `json:"completed" bson:"completed"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// Exported IsEmpty method for the Todo struct
func (t *Todo) IsEmpty() bool {
	return t.Task == ""
}

// Define the Response struct
type Response struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
