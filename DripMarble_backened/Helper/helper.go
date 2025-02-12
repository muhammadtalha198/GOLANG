package helpers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Database struct to hold MongoDB collection references
type Database struct {
	UserCollection     *mongo.Collection
	StakeCollection    *mongo.Collection
	WithdrawCollection *mongo.Collection
	ContractCollection *mongo.Collection
}
