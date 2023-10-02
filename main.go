package main

import (
	"context"
	"database/sql"
	"github.com/arilsonsantos/crud-mvc-go.git/src/config/database"
	"github.com/arilsonsantos/crud-mvc-go.git/src/config/routes"
	"github.com/arilsonsantos/crud-mvc-go.git/src/controller"
	"github.com/arilsonsantos/crud-mvc-go.git/src/repository"
	"github.com/arilsonsantos/crud-mvc-go.git/src/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	dbConn, _ := database.NewSqliteConnection(ctx)
	initDB(dbConn)

	router := gin.Default()
	userController := initDependencies(dbConn)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8585"); err != nil {
		log.Fatal(err)
	}

}

func initDependencies(dbConn *sql.DB) controller.UserControllerInterface {
	userRepository := repository.NewUserRepository(dbConn)
	userDomainService := service.NewUserServiceInterface(userRepository)
	return controller.NewUserControllerInterface(userDomainService)
}

func initDB(dbConn *sql.DB) {
	_, _ = dbConn.Exec(`
		CREATE TABLE IF NOT EXISTS usuario (
    		id TEXT PRIMARY KEY,
    		nome TEXT,
    		email TEXT,
    		senha TEXT,
    		idade INTEGER
    	)
	`)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return v.X + v.Y
}

func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
