package sources

import (
	"context"
	"os"
	"time"

	// import 'mongo-go-driver' package libraries

	// for BSON ObjectID

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func ConnectToMongoClient() *mongo.Client {
// 	os.Setenv("MONGODB_CONNSTR", "mongodb://mongo:parolazabaza@127.0.0.1:27035/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")
// 	var connStr = os.Getenv("MONGODB_CONNSTR")
// 	print("MONGODB_CONNSTR = " + connStr)

// 	options := options.Client()
// 	options.ApplyURI(connStr)

// 	client, err := mongo.Connect(context.Background(), options) //Only checking arguments
// 	if err != nil {
// 		print("Cannot connect to server")
// 		panic(err)
// 	}

// 	return client
// }

func CloseConnectionMongo(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func ConnectToMongoClient() (*mongo.Client, context.Context,
	context.CancelFunc, error) {
	//os.Setenv("MONGODB_CONNSTR", "mongodb://localhost:27017/mongo'-uparolazabaza-p?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")
	//var connStr = os.Getenv("MONGODB_CONNSTR")

	var host = os.Getenv("MONGODB_HOST")
	var port = os.Getenv("MONGODB_PORT")
	var user = os.Getenv("MONGODB_USER")
	var password = os.Getenv("MONGODB_PASSWORD")
	var readPreference = os.Getenv("MONGODB_READ_PREFERENCE")
	var appname = os.Getenv("MONGODB_APPNAME")
	var directConnection = os.Getenv("MONGODB_DIRECT_CONNECTION")
	var ssl = os.Getenv("MONGODB_SSL")

	var connStr = "mongodb://" + host + ":" + port + "/" +
		user + "'-u" +
		password + "-p" +
		"?" + "readPreference=" + readPreference +
		"&" + "appname=" + appname +
		"&" + "directConnection=" + directConnection +
		"&" + "ssl=" + ssl

	//fmt.Println("MONGODB_CONNSTR = " + connStr)

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	return client, ctx, cancel, err
}
