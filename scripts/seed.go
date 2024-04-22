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
	ctx        = context.Background()
)

func seedHotel(name, location string) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    make([]primitive.ObjectID, 0, 2),
	}
	res, err := hotelStore.Insert(context.Background(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	rooms := []types.Room{
		{
			HotelID:   res.ID,
			Type:      types.SingleRoomType,
			BasePrice: 99.9,
		},
		{
			HotelID:   res.ID,
			Type:      types.DoubleRoomType,
			BasePrice: 150.9,
		},
		{
			HotelID:   res.ID,
			Type:      types.SeaSideRoomType,
			BasePrice: 200.9,
		},
		{
			HotelID:   res.ID,
			Type:      types.DeluxeRoomType,
			BasePrice: 600,
		},
		{
			HotelID:   res.ID,
			Type:      types.DeluxeRoomType,
			BasePrice: 60500,
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
	seedHotel("Beluccia", "France")
	seedHotel("The cozy hotel", "Spain")
	seedHotel("Moscow", "Moscow")
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
}
