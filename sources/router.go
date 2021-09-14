package sources

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoUploadedFile struct {
	//ID              primitive.ObjectID `bson:"_id" json:"_id"`
	Encrypted_bytes []byte `bson:"encrypted_bytes" json:"encrypted_bytes"`
}

func SetupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", Get).Methods(http.MethodGet)
	api := r.PathPrefix("/api/v1").Subrouter()
	//api.HandleFunc("/decrypt", decryptFile).Methods(http.MethodPost)
	api.HandleFunc("/upload", UploadFile).Methods(http.MethodPost)
	api.HandleFunc("/download/{fileID}", DonwloadFile).Methods(http.MethodGet)
	api.HandleFunc("/file/{fileID}", GetFile).Methods(http.MethodGet)
	api.HandleFunc("/searchFile", SearchFile).Methods(http.MethodGet)

	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/",
		http.FileServer(http.Dir("./sources/templates/styles"))))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/",
		http.FileServer(http.Dir("./sources/templates/images"))))
	r.PathPrefix("/api/v1/styles/").Handler(http.StripPrefix("/api/v1/styles/",
		http.FileServer(http.Dir("./sources/templates/styles"))))
	r.PathPrefix("/api/v1/images/").Handler(http.StripPrefix("/api/v1/images/",
		http.FileServer(http.Dir("./sources/templates/images"))))
	r.PathPrefix("/api/v1/download/styles/").Handler(http.StripPrefix("/api/v1/download/styles/",
		http.FileServer(http.Dir("./sources/templates/styles"))))
	r.PathPrefix("/api/v1/download/images/").Handler(http.StripPrefix("/api/v1/download/images/",
		http.FileServer(http.Dir("./sources/templates/images"))))

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func Get(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-type", "application/json")
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(`{"message": "get called"}`))
	tpl.ExecuteTemplate(w, "index.html", nil) //Read about nginx
}

// func decryptFile(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Decrypting file\n")
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}

// 	r.ParseMultipartForm(10 << 20) // Max of 10 megabyte files

// 	//retrieve file from posted form-data
// 	file, handler, err := r.FormFile("decrImage")
// 	if err != nil {
// 		fmt.Println("Error Retrieving file from form-data")
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()
// 	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
// 	fmt.Printf("File Size: %+v\n", handler.Size)
// 	fmt.Printf("MIME Header: %+v\n", handler.Header)

// 	extension := filepath.Ext(handler.Filename)

// 	cipherBytes, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		fmt.Println("ERROR!")
// 		fmt.Println(err)
// 	}

// 	//Decrypting fileBytes----------
// 	db := ConnectToDb()
// 	defer db.Close()

// 	fileId := strings.Split(strings.Split(handler.Filename, ".")[0], "-")[1]
// 	key, err := GetKey(db, fileId)
// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid file id\n")
// 		fmt.Fprintf(w, "Redirecting to home\n")
// 		http.Redirect(w, r, "localhost:8080", http.StatusSeeOther)
// 	}
// 	//key := []byte("passphrasewhichneedstobe32bytes!")
// 	decryptedBytes := DecryptBytes(cipherBytes, key)

// 	WriteFileOnServer("decrypted-images", "decrypt-", extension, decryptedBytes)

// 	//outputFile.Write(decryptedBytes)

// 	fmt.Fprintf(w, "Successfully decrypted file\n")
// }

func UploadFile(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Uploading file\n")

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseMultipartForm(10 << 20) // Max of 10 megabyte files

	//retrieve file from posted form-data
	file, handler, err := r.FormFile("uplImage")
	if err != nil {
		parseError(err, w, r, "Error Retrieving file from form-data")
		//fmt.Println(err)
		//remhttp.Error(w, "my own error message", http.StatusForbidden)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	//extension := filepath.Ext(handler.Filename)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	//key := []byte("passphrasewhichneedstobe32bytes!")
	key := CreateUuidKey()

	//Encrypting fileBytes----------
	encryptedBytes := EncryptBytes(fileBytes, key)

	//outputFileName := WriteFileOnServer("uploaded-images", "upload-", extension, encryptedBytes)

	//Sending data to mongodb------------------
	mongoConn, ctx, cancel, err := ConnectToMongoClient()
	if err != nil {
		fmt.Println("Error connecting to db")
		panic(err)
	}
	defer CloseConnectionMongo(mongoConn, ctx, cancel)

	var dbname = os.Getenv("MONGODB_DB")
	database := mongoConn.Database(dbname)
	collection := database.Collection("uploaded-files")

	fileToUpload := MongoUploadedFile{
		Encrypted_bytes: encryptedBytes,
	}

	result, err := collection.InsertOne(ctx, fileToUpload)
	if err != nil {
		fmt.Println("Error inserting the file")
		panic(err)
	}

	fmt.Println("Successful upload to mongodb")
	resultStr := result.InsertedID.(primitive.ObjectID).String()
	mongoObjectId := strings.Split(resultStr, "\"")[1]

	//-------------------------------------------

	//Sending data to postgredb------------------
	db := ConnectToDb()
	defer db.Close()

	//fileId := strings.Split(strings.Split(outputFileName, ".")[0], "-")[2]
	fileId := string(CreateUuidKey())
	query := "INSERT INTO \"Uploads\" (file_id, orig_file_name, encryption_key, upload_date, mongodb_id) VALUES ($1, $2, $3, $4, $5)"

	//Get current time
	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)

	_, err = db.Exec(query, fileId, handler.Filename, key, dt, mongoObjectId)
	if err != nil {
		fmt.Println("Error executing insert statement")
		panic(err)
	}
	fmt.Println("Successful upload to postgre")
	//------------------------------------

	fileUrl := "http://localhost:8080/api/v1/download/" + fileId
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	tpl.ExecuteTemplate(w, "fileurl.html", fileUrl)

	//fmt.Fprintf(w, "Successfully uploaded file\n")
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	var fileId string
	var err error
	if val, ok := pathParams["fileID"]; ok {
		// fileId, err = val
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte(`{"message": "need a number"}`))
		// 	return
		// }
		fileId = val
	}

	//Another read filePath method
	db := ConnectToDb()
	defer db.Close()

	// fileUrl, err := GetFilePath(db, fileId)
	// if err != nil {
	// 	parseError(err, w, r, "Cannot get file_path")
	// 	return
	// }

	// fmt.Println("FileUrl = " + string(fileUrl))
	// file, err := os.Open(string(fileUrl))
	// if err != nil {
	// 	parseError(err, w, r, "Cannot open file")
	// 	return
	// }

	// fileBytes, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	parseError(err, w, r, "Cannot read fileBytes")
	// 	return
	// }

	key, err := GetKey(db, fileId)
	if err != nil {
		parseError(err, w, r, "Cannot get encryption_key")
		return
	}

	mongoDbId, err := GetMongoDbId(db, fileId)
	if err != nil {
		parseError(err, w, r, "Cannot get mongodb_id")
		return
	}

	//Get file data from MongoDB---
	mongoConn, ctx, cancel, err := ConnectToMongoClient()
	if err != nil {
		fmt.Println("Error connecting to db")
		panic(err)
	}
	defer CloseConnectionMongo(mongoConn, ctx, cancel)

	database := mongoConn.Database("MongoMainDB")
	collection := database.Collection("uploaded-files")

	// convert id string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(string(mongoDbId))
	if err != nil {
		log.Println("Invalid id")
	}
	fmt.Print("objectid: ", objectId)

	var fileFromDb MongoUploadedFile
	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&fileFromDb)
	if err != nil {
		log.Fatal(err)
	}
	// var filesFiltered []byte
	// if err = filterCursor.All(ctx, &filesFiltered); err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Println("filesFiltered: ")
	//fmt.Println(fileFromDb.ID)
	//fmt.Println(fileFromDb.Encrypted_bytes)

	//-----------------------------

	//fmt.Println("key = " + string(key))
	//key := []byte("passphrasewhichneedstobe32bytes!")
	//decryptedBytes := DecryptBytes(fileBytes, key)
	decryptedBytes := DecryptBytes(fileFromDb.Encrypted_bytes, key)

	//TODO: rethink!!!
	tempFileUrl := "temp-*" + fileId + ".jpg"
	tempFile, err := ioutil.TempFile("temp", tempFileUrl)
	if err != nil {
		parseError(err, w, r, "Cannot create tempFile")
		return
	}
	//defer os.RemoveAll("temp/*")
	defer RemoveContents("temp/")
	defer tempFile.Close()

	tempFile.Write(decryptedBytes)

	decrFile, err := os.Open(tempFile.Name())
	if err != nil {
		parseError(err, w, r, "Cannpt create tempFile")
		return
	}

	fmt.Println("temp-path : " + tempFile.Name())

	orig_file_name, err := GetOrigFileName(db, fileId)
	if err != nil {
		parseError(err, w, r, "Cannot get orig_file_name")
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+string(orig_file_name))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Type", r.Header.Get("Content-Length"))

	_, err = io.Copy(w, decrFile)
	if err != nil {
		parseError(err, w, r, "Cannot copy file")
		return
	}
}

func SearchFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fileId := r.Form.Get("fileId")
	fmt.Println("fileid= " + fileId)
	destUrl := "http://localhost:8080" + "/api/v1/download/" + fileId
	fmt.Println("destUrl= " + destUrl)
	//http.Redirect(w, r, destUrl, http.StatusOK)
	http.Redirect(w, r, destUrl, http.StatusFound)
}

func DonwloadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pathParams := mux.Vars(r)

	var fileId string
	var err error
	if val, ok := pathParams["fileID"]; ok {
		fileId = val
	}

	type fileMeta struct {
		ID         string
		Name       string
		UploadDate string
	}

	var fileToDownload fileMeta
	fileToDownload.ID = fileId

	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusOK)

	db := ConnectToDb()
	defer db.Close()

	//Get file data from mongodb
	fileNameBytes, err := GetOrigFileName(db, fileId)
	if err != nil {
		fmt.Println("Cannot get orig_file_name")
		tpl.ExecuteTemplate(w, "fileNotFound.html", fileToDownload)
		return
	}
	fileToDownload.Name = string(fileNameBytes)

	uploadDateBytes, err := GetUploadDate(db, fileId)
	if err != nil {
		fmt.Println("Cannot get upload_date")
		tpl.ExecuteTemplate(w, "fileNotFound.html", fileToDownload)
		return
	}

	fmt.Println("time = ", string(uploadDateBytes))

	layout := time.RFC3339
	uploadDateTimestamp, err := time.Parse(layout, string(uploadDateBytes))
	if err != nil {
		fmt.Println("Cannot parse file time")
		tpl.ExecuteTemplate(w, "fileNotFound.html", fileToDownload)
		return
	}

	fileToDownload.UploadDate = uploadDateTimestamp.Format("2 Jan 2006 15:04")

	tpl.ExecuteTemplate(w, "downloadFile.html", fileToDownload)
}

func parseError(err error, w http.ResponseWriter, r *http.Request, message string) {
	//Alert user message
	fmt.Println("ERROR!" + message)
	http.Redirect(w, r, "localhost:8080", http.StatusFound)
}
