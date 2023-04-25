package datasources

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/azambranapr/rmocs_collector_api/src/core/services"
	"github.com/azambranapr/rmocs_collector_api/src/data/repositories/batchRepository"
	"github.com/azambranapr/rmocs_collector_api/src/domain/repositories/batchService"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	_Client *mongo.Client
	_Ctx    context.Context
	_Cancel context.CancelFunc
)
var (
	port        int
	password    string
	user        string
	DataBase    string
	server      string
	drive       string
	Collections interface{}
)
var uri = fmt.Sprintf("%s://%s:%s@%s:%d", drive, user, password, server, port)

func New() {
	config, err := services.LoadConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	password = config.Db.Password
	user = config.Db.User
	DataBase = config.Db.Database
	drive = config.Server.Drive
	server = config.Server.Address
	port = config.Server.Port
	Collections = config.Db.Collections

	//echo -n Hello | base64 en Linux

	sDec, err := base64.StdEncoding.DecodeString(user)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return
	}

	user = string(sDec)
	sDec, err = base64.StdEncoding.DecodeString(password)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return
	}
	password = string(sDec)

	uri = fmt.Sprintf("%s://%s:%s@%s", drive, user, password, server)

	connect()
}

func connect() {

	fmt.Println("Enter connect")
	var err error

	_Ctx = context.TODO()
	_Client, err = mongo.Connect(_Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err = ping(_Client, _Ctx); err != nil {
		panic(err)
		//log.Fatalf("can't do ping: %v", err)
	}
	fmt.Println("connected successfully")

}

func ping(client *mongo.Client, ctx context.Context) error {

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("ping successfully")
	return nil
}

func DAO() (batchService.IdbService, error) {
	return batchRepository.NewRepository(_Client, _Ctx, _Cancel), nil
}
