package container_slice

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContainerModel struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name      string             `json:"name" validate:"required" bson:"name"`
	Path      string             `json:"path,omitempty" bson:"path"`
	Parent    string             `json:"parent,omitempty" bson:"parent"`
	CreatedAt string             `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt string             `json:"updatedAt,omitempty" bson:"updated_at"`
	DeletedAt string             `json:"deletedAt,omitempty" bson:"deleted_at"`
}
