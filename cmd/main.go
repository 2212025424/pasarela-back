package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pasarelapago.go/router"
	"pasarelapago.go/storage/postgres"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	loadEnv()
	startWithEcho()
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Print("No se cargaron las variables de entorno")
		panic(err)
	}
}

func startWithEcho() {
	db, err := postgres.New()

	if err != nil {
		log.Print("No se pudo cargar la base de datos")
	}

	e := echo.New()

	e.Use(middleware.Logger()) // Imprimir actividades
	e.Use(middleware.CORS())   // Permitir conexiones desde otros dominios
	/*e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://www.cobranza-legal.us", "https://api.paypal.com"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))*/                       // Permitir conexiones desde otros dominios
	e.Use(middleware.Recover()) // Recuperar aplicacion de los panic de los handler

	router.Product(e, db)
	router.Order(e, db)
	router.Subscription(e, db)
	router.Invoice(e, db)
	router.PayPal(e, db)

	port := os.Getenv("HTTP_PORT")

	if port == "" {
		port = ":8080"
	}

	//log.Print(e.Start(port))
	log.Print(e.StartTLS(port, os.Getenv("PUBLIC_KEY"), os.Getenv("PRIVATE_KEY")))
}
