package main

import (
	"Mini-Project_Coaching-Clinic/configs"
	"Mini-Project_Coaching-Clinic/route"
)

func main() {
	configs.Init()

	e := route.NewRouter()           // Init Router
	e.Logger.Fatal(e.Start(":8080")) // Start Server

	
}
