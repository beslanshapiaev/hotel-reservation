package main

import (
	"context"
	"fmt"
	"log"

	"github.com/beslanshapiaev/hotel-reservation/db"
	"github.com/beslanshapiaev/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	userStore  db.UserStore
	ctx        = context.Background()
)

func seedUser(firstname, lastname, email string) {
	user, _ := types.NewUserFromParams(types.CreateUserParams{
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Password:  "Secret123",
	})

	_, err := userStore.InsertUser(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
}

func seedHotel(name, location string, rating int) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    make([]primitive.ObjectID, 0, 2),
		Rating:   rating,
	}
	res, err := hotelStore.Insert(context.Background(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	rooms := []types.Room{
		{
			HotelID: res.ID,
			Size:    "small",
			Price:   99.9,
		},
		{
			HotelID: res.ID,
			Size:    "medium",
			Price:   150.9,
		},
		{
			HotelID: res.ID,
			Size:    "kingsize",
			Price:   200.9,
		},
		{
			HotelID: res.ID,
			Size:    "medium",
			Price:   600,
		},
		{
			HotelID: res.ID,
			Size:    "kingsize",
			Price:   60500,
		},
	}
	for _, r := range rooms {
		_, err = roomStore.InsertRoom(context.Background(), &r)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	seedHotel("Beluccia", "France", 3)
	seedHotel("The cozy hotel", "Spain", 4)
	seedHotel("Moscow", "Moscow", 5)
	seedUser("beslan", "shapiaev", "mishastalinaa@yandex.ru")
	fmt.Println("seeding the database")
}

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBNAME).Drop(context.Background()); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client, db.DBNAME)
}
