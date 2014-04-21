package service 

import (
  "../model"
 "fmt"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

func mongodbNewPerson() {
  newPerson := model.Person{"Ale", "1234556"}
  fmt.Println(newPerson.Name)

  session, err := mgo.Dial("192.168.2.48")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)

  c := session.DB("test").C("people")
  err = c.Insert(&model.Person{"Ale", "+55 53 8116 9639"},
                 &model.Person{"Cla", "+55 53 8402 8510"})
  if err != nil {
    panic(err)
  }

  result := model.Person{}
  err = c.Find(bson.M{"name": "Ale"}).One(&result)
  if err != nil {
    panic(err)
  }

  fmt.Println("Phone:", result.Phone)
}
