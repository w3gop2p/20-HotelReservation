package main

import (
	"20-HotelReservation/db"
	"20-HotelReservation/db/fixtures"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
/*
client       *mongo.Client
roomStore    db.RoomStore
hotelStore   db.HotelStore
userStore    db.UserStore
bookingStore db.BookingStore
ctx          = context.Background()
*/
)

/*
	func seedUser(isAdmin bool, fname, lname, email string, password string) *types.User {
		user, err := types.NewUserFromParams(types.CreateUserParams{
			Email:     email,
			FirstName: fname,
			LastName:  lname,
			Password:  password,
		})
		if err != nil {
			log.Fatal(err)
		}
		user.IsAdmin = isAdmin
		insertedUser, err := userStore.InsertUser(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s -> %s\n", user.Email, api.CreateTokenFromUser(user))
		return insertedUser
	}

func seedHotel(name string, location string, rating int) *types.Hotel {

		hotel := types.Hotel{
			Name:     name,
			Location: location,
			Rooms:    []primitive.ObjectID{},
			Rating:   rating,
		}

		rooms := []types.Room{
			{
				Size:      "small",
				BasePrice: 99.9,
			},
			{
				Size:      "normal",
				BasePrice: 123.1,
			},
			{
				Size:      "large",
				BasePrice: 423.42,
			},
			{
				Size:      "kingsize",
				BasePrice: 324.21,
			},
		}
		insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
		if err != nil {
			log.Fatal(err)
		}

		for _, room := range rooms {
			room.HotelID = insertedHotel.ID
			_, err := roomStore.InsertRoom(ctx, &room)
			if err != nil {
				log.Fatal(err)
			}
		}
		return insertedHotel
	}

	func seedRoom(size string, ss bool, price float64, hotelID primitive.ObjectID) *types.Room {
		room := &types.Room{
			Size:    size,
			Seaside: ss,
			Price:   price,
			HotelID: hotelID,
		}
		insertedRoom, err := roomStore.InsertRoom(context.Background(), room)
		if err != nil {
			log.Fatal(err)
		}
		return insertedRoom
	}

	func seedBooking(userID, roomID primitive.ObjectID, from, till time.Time) {
		booking := &types.Booking{
			UserID:   userID,
			RoomID:   roomID,
			FromDate: from,
			TillDate: till,
		}
		resp, err := bookingStore.InsertBooking(context.Background(), booking)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("booking:", resp.ID)
	}
*/
func main() {
	/*var err error*/
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var (
		ctx           = context.Background()
		mongoEndPoint = os.Getenv("MONGO_DB_URL")
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoEndPoint))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client)

	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Booking: db.NewMongoBookingStore(client),
		Room:    db.NewMongoRoomStore(client, hotelStore),
		Hotel:   hotelStore,
	}

	user := fixtures.AddUser(store, "kim", "lee", false)
	fmt.Println(user)

	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println(admin)

	hotel := fixtures.AddHotel(store, "Simple Hotel", "Casandra", 5, nil)
	fmt.Println(hotel)

	room := fixtures.AddRoom(store, "large", true, 88.44, hotel.ID)
	fmt.Println(room)

	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))
	fmt.Println(booking)

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("random hotel name %d", i)
		location := fmt.Sprintf("location %d", i)
		fixtures.AddHotel(store, name, location, rand.Int()+1, nil)
	}

	/*seedUser(false, "james", "bond", "james@bond.com", "parola")
	seedUser(true, "admin", "admin", "admin@admin.com", "parola")
	alice := seedUser(false, "Alice", "Anderson", "alice@mail.com", "parola")

	seedHotel("Belluccia", "France", 3)
	seedHotel("The cozy hotel", "The Netherlands", 4)
	hotel = seedHotel("Awesome Hotel", "Moldavia", 5)

	room := seedRoom("small", true, 100, hotel.ID)
	seedRoom("medium", true, 150, hotel.ID)
	seedRoom("large", false, 200, hotel.ID)

	seedBooking(alice.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2))*/

}

/*func init() {
	var err error
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
	bookingStore = db.NewMongoBookingStore(client)

}*/
