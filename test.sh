rm -rf apisvr
rm -rf docs model public service
rm ../testsvr/apisvr/go.mod
go run main.go gen -n apisvr -o ../testsvr
