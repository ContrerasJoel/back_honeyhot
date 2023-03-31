package product

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Link         string             `json:"link"`
	Title        string             `json:"title"`
	Descriptions Descriptions       `json:"descriptions"`
	Price        int                `json:"price"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt     time.Time          `bson:"update_at" json:"update_at,omitempty"`
}

type Description struct {
	Key   string
	Value string
}

type Products []*Product

type Descriptions []Description
