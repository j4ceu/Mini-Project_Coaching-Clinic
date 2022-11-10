package main

import (
	"Mini-Project_Coaching-Clinic/configs"
	"Mini-Project_Coaching-Clinic/route"
	"os"
)

func main() {
	configs.Init()

	e := route.NewRouter()                     // Init Router
	e.Logger.Fatal(e.Start(os.Getenv("PORT"))) // Start Server

}
