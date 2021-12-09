package route

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reddit/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Collection = model.ConnectDB()

func GetBoards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var boards []model.Reddit
	cur, err := Collection.Find(context.TODO(), bson.M{})

	if err != nil {
		model.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var board model.Reddit
		err := cur.Decode(&board) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		boards = append(boards, board)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(boards)

}

func GetBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var board model.Reddit
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := Collection.FindOne(context.TODO(), filter).Decode(&board)

	if err != nil {
		model.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(board)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var board model.Reddit

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&board)

	// insert our book model.
	result, err := Collection.InsertOne(context.TODO(), board)

	if err != nil {
		model.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var board model.Reddit

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&board)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"title", board.Title},
			{"board", board.Board},
			{"author", bson.D{
				{"firstname", board.Author.FirstName},
				{"lastname", board.Author.LastName},
			}},
		}},
	}

	err := Collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&board)

	if err != nil {
		model.GetError(err, w)
		return
	}

	board.Id = id

	json.NewEncoder(w).Encode(board)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := Collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		model.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}
