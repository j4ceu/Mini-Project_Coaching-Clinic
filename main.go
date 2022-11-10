package main

import (
	"Mini-Project_Coaching-Clinic/configs"
	"Mini-Project_Coaching-Clinic/route"
	"os"
)

func main() {
	configs.Init()

	e := route.NewRouter() // Init Router
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port)) // Start Server

}
