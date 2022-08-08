package main

import (
	  
 "carlamissiona/golang-barbers/bootstrap"
  "log"
)

func main() {
  
	app := bootstrap.NewApplication()  
	log.Fatal(app.Listen(":5000"))
 
}
   