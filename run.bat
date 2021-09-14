setx POSTGRES_HOST 				"127.0.0.1"
setx POSTGRES_PORT 				"5432"
setx POSTGRES_DB 				"MainDB"
setx POSTGRES_USER 				"postgres"
setx POSTGRES_PASSWORD 			"parolazabaza"
setx POSTGRES_CONNECT_TIMEOUT 	"20"
setx POSTGRES_SSLMODE 			"disable"

REM setx CONNSTR "user=postgres password=parolazabaza host=127.0.0.1 port=5432 dbname=MainDB connect_timeout=20 sslmode=disable"

setx MONGODB_HOST 				"127.0.0.1"
setx MONGODB_PORT 				"27017"
setx MONGODB_DB 				"MongoMainDB"
setx MONGODB_USER 				"mongo"
setx MONGODB_PASSWORD 			"parolazabaza"
setx MONGODB_READ_PREFERENCE 	"primary"
setx MONGODB_APPNAME 			"MongoDB%20Compass"
setx MONGODB_DIRECT_CONNECTION 	"true"
setx MONGODB_SSL 				"false"


REM setx MONGODB_CONNSTR "mongodb://localhost:27017/mongo'-uparolazabaza-p?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"

setx SERVER_HOST 	"localhost"
setx SERVER_PORT 	"80"

go run main.go