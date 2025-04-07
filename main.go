package main

import (
	"daily-brew/config"
	"daily-brew/models"
	"daily-brew/routes"
	"fmt"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	config.InitConfig()
	config.InitDB()
	config.InitRedis()
	config.DB.AutoMigrate(&models.Member{})
	r := routes.SetupRoutes()
	serverAddr := fmt.Sprintf("127.0.0.1:%s", config.AppConfig.ServerPort)
	log.Printf("Server running on http://localhost%s\n", serverAddr)
	r.Run(serverAddr)
}
