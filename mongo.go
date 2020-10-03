package main

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

/////////////////////////////////////////////
// Password manipulations
/////////////////////////////////////////////

// HashString - encoding string
func HashString(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash - checking if entered password == encoded password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/////////////////////////////////////////////
// User manipulations
/////////////////////////////////////////////

func checkFromDB(username string, password string) bool {

	// Getting the collection (table)
	collection := Client.Database("courses").Collection("users")

	// Filter by name
	filter := bson.M{"username": username}

	var result user

	// Writing result of filtration to result var
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return false
	}

	// Checking if entered password is the same with hashed
	if CheckPasswordHash(password, result.Password) {
		return true
	}

	log.Println("Wrong password")
	return false

}

func addUserToDB(user user) {

	collection := Client.Database("courses").Collection("users")

	// Inserting user to DB
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err)
	}
	log.Println(insertResult)
}

func checkUserExist(username string) bool {

	collection := Client.Database("courses").Collection("users")

	filter := bson.M{"username": username}

	var result user
	var ret bool

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		ret = true
	}
	return ret
}

/////////////////////////////////////////////
// Article manipulations
/////////////////////////////////////////////

func getArticleFromDB() []article {

	collection := Client.Database("courses").Collection("articles")

	// Here's an array in which we store the decoded documents
	var results []article

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// Create a value into which the single document can be decoded
		var elem article
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Closing the cursor
	cur.Close(context.TODO())

	return results
}

func insertArticleToDB(a article) {

	collection := Client.Database("courses").Collection("articles")

	insertResult, err := collection.InsertOne(context.TODO(), a)
	if err != nil {
		log.Println(err)
	}
	log.Println(insertResult)
}

// Adding comment to DB
func commentToDB(comtitle, commentStr, time, name string) error {

	collection := Client.Database("courses").Collection("articles")
	filter := bson.M{"title": comtitle}

	update := bson.M{
		"$push": bson.M{"comment": comment{ComTime: time, ComContent: commentStr, ComName: name}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Println(err)
		return errors.New("Failed to update")
	}
	log.Println(updateResult)
	return nil
}
