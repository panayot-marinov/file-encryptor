export CONNSTR="user=postgres password=parolazabaza host=127.0.0.1 port=5432 dbname=MainDB connect_timeout=20 sslmode=disable"
cd tests
go test