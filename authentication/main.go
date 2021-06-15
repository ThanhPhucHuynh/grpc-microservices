package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ThanhPhucHuynh/db"
	"github.com/joho/godotenv"
)


var (
	local bool
)

func init() {
	flag.BoolVar(&local,"local",true,"run server local")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}
	// fmt.Println("hah")
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(conn.DB().CollectionNames())

	defer conn.Close()
}