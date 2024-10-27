package main

import (
	"context"
	"fmt"
	"payment/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	db := database.Connect()
	var result bson.M
	if err := db.DB.Client().Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}

/*
	models -> controller -> routes.

Design and Implement a backend service for Wallet system supporting
Setup wallet
Credit / Debit transactions
Fetching transactions on wallet
Get wallet details

*/
