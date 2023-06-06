package sources

import (
	"database/sql"
	"fmt"
)

func getValue(db *sql.DB, fileId string, columnName string) ([]byte, error) {
	var val []byte
	row := db.QueryRow("SELECT "+columnName+" FROM \"Uploads\" where file_id=$1", fileId)
	if err := row.Scan(&val); err != nil {
		fmt.Println("ERROR! Cannot execute query!")
		return val, err
	}
	return val, nil
}

func GetMongoDbId(db *sql.DB, fileId string) ([]byte, error) {
	return getValue(db, fileId, "mongodb_id")
}

func GetKey(db *sql.DB, fileId string) ([]byte, error) {
	return getValue(db, fileId, "encryption_key")
}

func GetFilePath(db *sql.DB, fileId string) ([]byte, error) {
	return getValue(db, fileId, "file_path")
}

func GetOrigFileName(db *sql.DB, fileId string) ([]byte, error) {
	return getValue(db, fileId, "orig_file_name")
}
