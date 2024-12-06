package main

import (
	EmployeeController "backend/api/controller/employee"
	"backend/api/db"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Get .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	//get InitDB fuction
	db.InitDB()

	router := gin.Default()
	//Employee API Method       //GET
	router.GET("/employeedb", EmployeeController.GetEmployeeDB)     //GET DB
	router.GET("/employee/:id", EmployeeController.GetEmployeeByID) //GET BY ID
	router.POST("/employee", EmployeeController.PostEmployee)       //POST
	router.PUT("/employee", EmployeeController.PutEmployee)         //PUT
	router.DELETE("/employee", EmployeeController.DeleteEmployee)   //DELETE

	//Customer API Method

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
