package main

import (
  . "./model"
  "github.com/ant0ine/go-json-rest/rest"
  "net/http"
  "fmt"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type Message struct {
  Body string
}

type MongoPersister struct {

}

func (mongoPersister *MongoPersister) InsertPerson() {
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)

  c := session.DB("test").C("people")
  err = c.Insert(Person{"Alex", "+55 53 8116 9639"},
                 Person{"Klan", "+55 53 8402 8510"})
  if err != nil {
    panic(err)
  }
}

func (mongoPersister *MongoPersister) GetPersonByName() {
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)

  c := session.DB("test").C("people")
  result := Person{}
  err = c.Find(bson.M{"name": "Ale"}).One(&result)
  if err != nil {
    panic(err)
  }

  fmt.Println("Phone:", result.Phone)
}

func main() {
    mongoPersister := MongoPersister{}
    handler := rest.ResourceHandler{}
    handler.SetRoutes(
      &rest.Route{"GET", "/api/v1/change_pk", func(w rest.ResponseWriter, req *rest.Request) {
        var longBook Book = Book{Title: "Harry Potter", Author: "JK Rolling", Pages: 1000}
        mongoPersister.InsertPerson() 
        w.WriteJson(&Message{
          Body: longBook.CategoryByLength(),
        })
      }},
    )
    http.ListenAndServe(":8080", &handler)
}
