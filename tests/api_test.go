package tests

// import (
// 	"file-encryptor/sources"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	_ "database/sql"

// 	"github.com/gorilla/mux"
// 	_ "github.com/lib/pq" //postgres drivers for initialization
// )

// func TestGetMethod(t *testing.T) {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", sources.Get).Methods(http.MethodGet)
// 	//!!!log.Fatal(http.ListenAndServe(":8080", r))

// 	//Create a mock request
// 	req, err := http.NewRequest(http.MethodGet, "/", nil)
// 	if err != nil {
// 		t.Fatalf("Could not create request: %v\n", err)
// 	}

// 	//Create a response recorder so you can inspect the response
// 	w := httptest.NewRecorder()

// 	//Perform the request
// 	r.ServeHTTP(w, req)
// 	//fmt.Println(w.Body)

// 	//Check to see if the response was what is expected

// 	if w.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
// 	}
// }

// func TestUploadMethod(t *testing.T) {
// 	r := mux.NewRouter()
// 	api := r.PathPrefix("/api/v1").Subrouter()
// 	api.HandleFunc("/upload", sources.UploadFile).Methods(http.MethodPost)

// 	// file, err := os.Open("C:\\Users\\C5328147\\Desktop\\download.jpg")
// 	// if err != nil {
// 	// 	t.Fatalf("Could not create request: %v\n", err)
// 	// }
// 	// defer file.Close()

// 	//Create a mock request
// 	req, err := http.NewRequest(http.MethodPost, "/api/v1/upload", nil)
// 	if err != nil {
// 		t.Fatalf("Could not create request: %v\n", err)
// 	}
// 	req.Header.Set("Content-Type", "image/jpeg")

// 	//Create a response recorder so you can inspect the response
// 	w := httptest.NewRecorder()

// 	//Perform the request
// 	//r.ServeHTTP(w, req)
// 	sources.UploadFile(w, req)
// 	//fmt.Println(w.Body)

// 	//Check to see if the response was what is expected

// 	if w.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
// 	}
// }

// // func TestGetFileMethod(t *testing.T) {
// // 	r := mux.NewRouter()
// // 	api := r.PathPrefix("/api/v1").Subrouter()
// // 	api.HandleFunc("/file/{fileID}", sources.GetFile).Methods(http.MethodGet)

// // 	//Create a mock request
// // 	req, err := http.NewRequest(http.MethodGet, "/api/v1/file/927352357", nil)
// // 	if err != nil {
// // 		t.Fatalf("Could not create request: %v\n", err)
// // 	}

// // 	//Create a response recorder so you can inspect the response
// // 	w := httptest.NewRecorder()

// // 	//Perform the request
// // 	r.ServeHTTP(w, req)
// // 	//fmt.Println(w.Body)

// // 	//Check to see if the response was what is expected

// // 	if w.Code == http.StatusFound {
// // 		t.Logf("Expected to get status %d is same as %d\n", http.StatusFound, w.Code)
// // 	} else {
// // 		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusFound, w.Code)
// // 	}
// // }
