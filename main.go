package main

import (
	"fmt"
	"log"
	"net/http"

	r "gopkg.in/gorethink/gorethink.v3"
)

type run struct {
	Title     string
	Distance  int
	StartTime int64
	EndTime   int64
}

var runs = []run{
	{Title: "monday", Distance: 1600, StartTime: 1506294329824, EndTime: 1506295329824},
	{Title: "tuesday", Distance: 3200, StartTime: 1506294329824, EndTime: 1506295329824},
	{Title: "wednesday", Distance: 4800, StartTime: 1506294329824, EndTime: 1506295329824},
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
	})
	if err != nil {
		log.Fatalln("35", err)
	}

	router := NewRouter(session)

	initializeRoutes()

	fmt.Println("Now listening at http://localhost:8082")
	log.Fatal(http.ListenAndServe("localhost:8082", nil))
}

func initializeRoutes() {
	router
}

func createTable(client Client, name string) error {
	_, err := r.TableCreate(name).Run(client.session)
	if err != nil {
		return err
	}
	return nil
}

func initDatabase(client *Client) {
	res, err := r.Table("runs").Insert(runs).Run(client.session)
	if err != nil {
		log.Fatal("45", err)
	}

	fmt.Println(res)
}
