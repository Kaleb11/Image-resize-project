package initiator

import (
	"fmt"
	"os"
)

const (
	dbbURL = "host=%v user=%v password=%v dbname=%v port=%v sslmode=disable"
)

var (
	dbUser  = os.Getenv("CR_USER")
	dbName  = os.Getenv("CR_NAME")
	dbPass  = os.Getenv("CR_PASS")
	dbHost  = os.Getenv("CR_HOST")
	dbPort  = os.Getenv("CR_PORT")
	imgPath = os.Getenv("imagepath")
	dbURL   = fmt.Sprintf(dbbURL, dbHost, dbUser, dbPass, dbName, dbPort)
)
