export CONNSTR="user=postgres password=parolazabaza host=127.0.0.1 port=5432 dbname=MainDB connect_timeout=20 sslmode=disable"
export MONGODB_CONNSTR="mongodb://localhost:27017/mongo'-uparolazabaza-p?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
go run main.go