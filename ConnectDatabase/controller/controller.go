package controller

import "go.mongodb.org/mongo-driver/mongo"

const connectionString = "mongodb+srv://admin:i2345@cluster0.zqufr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const colName = "watchList"

// Most Important
var collection *mongo.Collection


