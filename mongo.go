package peda

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func GetAllBackend(mongoconn *mongo.Database, collection string) []Backend {
	sidang := atdb.GetAllDoc[[]Backend](mongoconn, collection)
	return sidang
}

func UpdateFrontend(mongoconn *mongo.Database, collection string, filter bson.M, sidangdata FormInputAll) interface{} {
	filter = bson.M{"Nama_dosen": sidangdata.Nama_dosen}
	return atdb.ReplaceOneDoc(mongoconn, collection, filter, sidangdata)
}

func FindFrontend(mongoconn *mongo.Database, collection string, userdata FormInputAll) Frontend {
	filter := bson.M{"Nama_dosen": userdata.Nama_dosen}
	return atdb.GetOneDoc[Frontend](mongoconn, collection, filter)
}

func FindUserUser(mongoconn *mongo.Database, collection string, userdata User) User {
	filter := bson.M{
		"username": userdata.Username,
	}
	return atdb.GetOneDoc[User](mongoconn, collection, filter)
}

func FindUser(mongoconn *mongo.Database, collection string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	return atdb.GetOneDoc[User](mongoconn, collection, filter)
}
func FindNik(mongoconn *mongo.Database, collection string, userdata User) User {
	filter := bson.M{"nik": userdata.Username}
	return atdb.GetOneDoc[User](mongoconn, collection, filter)
}

func FindUserByname(mongoconn *mongo.Database, collection string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	return atdb.GetOneDoc[User](mongoconn, collection, filter)
}
func DeleteFrondent(mongoconn *mongo.Database, collection string, sidangdata Frontend) interface{} {
	filter := bson.M{"npm": sidangdata.Npm}
	return atdb.DeleteOneDoc(mongoconn, collection, filter)
}

func CreateFronent(mongoconn *mongo.Database, collection string, sidangdata Frontend) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, sidangdata)
}
func InsertAllinput(mconn *mongo.Database, collname string, datafilm FormInputAll) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datafilm)
}

func CreateAllInput(mongoconn *mongo.Database, collection string, sidangdata FormInputAll) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, sidangdata)
}

func InsertUser(mongoconn *mongo.Database, collection string, userdata User) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, userdata)
}

func UsernameExists(mongoenvkatalogfilm, dbname string, userdata User) bool {
	mconn := SetConnection(mongoenvkatalogfilm, dbname).Collection("user")
	filter := bson.M{"username": userdata.Username}

	var user User
	err := mconn.FindOne(context.Background(), filter).Decode(&user)
	return err == nil
}

func IsPasswordValid(mconn *mongo.Database, collname string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, collname, filter)
	hashChecker := CheckPasswordHash(userdata.Password, res.Password)
	return hashChecker
}
func GetAllUser(mconn *mongo.Database, collname string) []User {
	user := atdb.GetAllDoc[[]User](mconn, collname)
	return user
}

func DeleteAllform(mconn *mongo.Database, collname string, datafilm FormInputAll) interface{} {
	filter := bson.M{"nama_dosen": datafilm.Nama_dosen}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

func UpdateForm(mongoconn *mongo.Database, collection string, filter bson.M, sidangdata FormInputAll) interface{} {
	filter = bson.M{"nama_dosen": sidangdata.Nama_dosen}
	return atdb.ReplaceOneDoc(mongoconn, collection, filter, sidangdata)
}
func FindName(mongoconn *mongo.Database, collection string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	return atdb.GetOneDoc[User](mongoconn, collection, filter)
}

func FindBackend(mongoconn *mongo.Database, collection string, userdata FormInputAll) User {
	filter := bson.M{"nama_dosen": userdata.Nama_dosen}
	return atdb.GetOneDoc[User](mongoconn, collection, filter)
}
func FindFilm(mconn *mongo.Database, collname string, datafilm FormInputAll) FormInputAll {
	filter := bson.M{"nama_dosen": datafilm.Nama_dosen}
	return atdb.GetOneDoc[FormInputAll](mconn, collname, filter)
}

func Getall(mongoconn *mongo.Database, collection string) FormInputAll {
	sidang := atdb.GetAllDoc[FormInputAll](mongoconn, collection)
	return sidang
}
