package peda

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func Authorization(publickeykatalogkemanan, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response CredentialUser
	var auth User
	response.Status = false

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return GCFReturnStruct(response)
	}

	tokenname := DecodeGetName(os.Getenv(publickeykatalogkemanan), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogkemanan), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogkemanan), header)

	auth.Username = tokenusername

	if tokenname == "" || tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return GCFReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, auth) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	response.Message = "Berhasil decode token"
	response.Status = true
	response.Data.No_whatsapp = tokenname
	response.Data.Username = tokenusername
	response.Data.Role = tokenrole

	return GCFReturnStruct(response)
}

func Registrasi(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	if UsernameExists(mongoenvkatalogfilm, dbname, user) {
		response.Message = "Username telah dipakai"
		return GCFReturnStruct(response)
	}

	hash, hashErr := HashPassword(user.Password)
	if hashErr != nil {
		response.Message = "Gagal hash password: " + hashErr.Error()
		return GCFReturnStruct(response)
	}

	user.Password = hash

	InsertUser(mconn, collname, user)
	response.Status = true
	response.Message = "Berhasil input data"

	return GCFReturnStruct(response)
}

func Login(privatekeykatalogkemanan, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, user) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	if !IsPasswordValid(mconn, collname, user) {
		response.Message = "Password Salah"
		return GCFReturnStruct(response)
	}

	auth := FindUser(mconn, collname, user)

	tokenstring, tokenerr := Encode(auth.No_whatsapp, auth.Username, auth.Role, os.Getenv(privatekeykatalogkemanan))
	if tokenerr != nil {
		response.Message = "Gagal encode token: " + tokenerr.Error()
		return GCFReturnStruct(response)
	}

	response.Status = true
	response.Message = "Berhasil login"
	response.Token = tokenstring

	return GCFReturnStruct(response)
}

func CreateAllform(publickeykatalogkemanan, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var film FormInputAll
	err := json.NewDecoder(r.Body).Decode(&film)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return GCFReturnStruct(response)
	}
	tokenname := DecodeGetName(os.Getenv(publickeykatalogkemanan), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogkemanan), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogkemanan), header)

	if tokenusername == "" || tokenrole == "" || tokenname == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return GCFReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	if tokenrole != "admin" && tokenrole != "dosen" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}

	CreateAllInput(mconn, collname, FormInputAll{
		Nama_dosen:  film.Nama_dosen,
		Nik:         film.Nik,
		Dosen:       film.Dosen,
		Suratall:    film.Suratall,
		Sertificate: film.Sertificate,
		Akademis:    film.Akademis,
		Status:      film.Status,
	})
	response.Status = true
	response.Message = "Berhasil input data"

	return GCFReturnStruct(response)
}

func HapusFilm(publickeykatalogkemanan, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var film FormInputAll
	err := json.NewDecoder(r.Body).Decode(&film)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return GCFReturnStruct(response)
	}
	tokenname := DecodeGetName(os.Getenv(publickeykatalogkemanan), header)

	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogkemanan), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogkemanan), header)

	if tokenusername == "" || tokenrole == "" || tokenname == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return GCFReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	if tokenrole != "admin" && tokenrole != "dosen" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}

	DeleteAllform(mconn, collname, film)
	response.Status = true
	response.Message = "Berhasil hapus data"

	return GCFReturnStruct(response)
}

func UpdateFilm(publickeykatalogkemanan, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var film FormInputAll
	err := json.NewDecoder(r.Body).Decode(&film)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return GCFReturnStruct(response)
	}
	tokenname := DecodeGetName(os.Getenv(publickeykatalogkemanan), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogkemanan), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogkemanan), header)

	if tokenusername == "" || tokenrole == "" || tokenname == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return GCFReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return GCFReturnStruct(response)
	}

	if tokenrole != "admin" && tokenrole != "dosen" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	auth3 := FindBackend(mconn, collname, film)
	oldNamaDosen := auth3.Username // Save the old value
	UpdateForm(mconn, collname, bson.M{"nama_dosen": oldNamaDosen}, film)
	response.Status = true
	response.Message = "Berhasil update data"

	return GCFReturnStruct(response)
}

func AmbilSatuFilm(publickeykatalogkemanan, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var film FormInputAll
	err := json.NewDecoder(r.Body).Decode(&film)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	auth3 := FindBackend(mconn, collname, film)

	if auth3.Username == DecodeGetName(os.Getenv(publickeykatalogkemanan), r.Header.Get("token")) {
		response.Message = "Anda tidak memiliki akses"
		GCFReturnStruct(response)
	} else {
		datafilm := FindFilm(mconn, collname, film)
		return GCFReturnStruct(datafilm)
	}
	return GCFReturnStruct(response)
}
func AmbilSemuaFilm(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	datafilmm := Getall(mconn, collname)
	return GCFReturnStruct(datafilmm)
}

// Keamanaan (Tranfer data Encrypt dan Decrypt)

func Encrypt(publickeykatalogkemanan, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	var encrypting Backend
	err := json.NewDecoder(r.Body).Decode(&encrypting)
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	encrypt := EncryptData(os.Getenv(publickeykatalogkemanan), encrypting.Autentikasitoken)
	response.Status = true
	response.Message = "Berhasil encrypt data"
	response.Data = encrypt

	return GCFReturnStruct(response)
}

func Decrypt(publickeykatalogkemanan, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	var decrypting Document
	err := json.NewDecoder(r.Body).Decode(&decrypting)
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	decrypt := DecryptData(os.Getenv(publickeykatalogkemanan), decrypting.Encrypted_Docs.Autentikasitoken)
	response.Status = true
	response.Message = "Berhasil decrypt data"
	response.Data = decrypt

	return GCFReturnStruct(response)
}

func Base64Encrypt(publickeykatalogkemanan, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	var encryptdata DocumentInput
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&encryptdata)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	encrypt := base64.StdEncoding.EncodeToString(buf.Bytes())
	response.Status = true
	response.Message = "Berhasil encrypt data"
	response.Data = encrypt

	return GCFReturnStruct(response)
}

func Base64Decrypt(publickeykatalogkemanan, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	var encryptdata Document
	err := json.NewDecoder(r.Body).Decode(&encryptdata)
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	decoded, err := base64.StdEncoding.DecodeString(encryptdata.Encrypted_Docs)
	if err != nil {
		response.Message = "Error decoding base64: " + err.Error()
		return GCFReturnStruct(response)
	}
	encrypt := string(decoded)
	response.Status = true
	response.Message = "Berhasil decrypt data"
	response.Data = encrypt

	return GCFReturnStruct(response)
}
