package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json: "id"`
	Title  string  `json: "title"`
	Artist string  `json: "artist"`
	Price  float64 `json: "price"`
}

// SLICE DE TIPO ALBUM
var albums = []album{
	{ID: "1", Title: "Kind of Blue", Artist: "Miles Davis", Price: 49.99},
	{ID: "2", Title: "Abbey Road", Artist: "The Beatles", Price: 39.99},
	{ID: "3", Title: "Thriller", Artist: "Michael Jackson", Price: 29.99},
	{ID: "4", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 34.99},
	{ID: "5", Title: "Back in Black", Artist: "AC/DC", Price: 24.99},
	{ID: "6", Title: "Rumours", Artist: "Fleetwood Mac", Price: 44.99},
	{ID: "7", Title: "Born to Run", Artist: "Bruce Springsteen", Price: 27.99},
	{ID: "8", Title: "Led Zeppelin IV", Artist: "Led Zeppelin", Price: 36.99},
	{ID: "9", Title: "Hotel California", Artist: "Eagles", Price: 31.99},
	{ID: "10", Title: "Sgt. Pepper's Lonely Hearts Club Band", Artist: "The Beatles", Price: 41.99},
	{ID: "11", Title: "Darkness on the Edge of Town", Artist: "Bruce Springsteen", Price: 28.99},
	{ID: "12", Title: "Highway 61 Revisited", Artist: "Bob Dylan", Price: 33.99},
	{ID: "13", Title: "Pet Sounds", Artist: "The Beach Boys", Price: 45.99},
	{ID: "14", Title: "Who's Next", Artist: "The Who", Price: 29.99},
	{ID: "15", Title: "A Night at the Opera", Artist: "Queen", Price: 37.99},
	{ID: "16", Title: "The Wall", Artist: "Pink Floyd", Price: 39.99},
	{ID: "17", Title: "Purple Rain", Artist: "Prince", Price: 32.99},
	{ID: "18", Title: "Born to Die", Artist: "Lana Del Rey", Price: 26.99},
	{ID: "19", Title: "Exile on Main St.", Artist: "The Rolling Stones", Price: 48.99},
	{ID: "20", Title: "The Chronic", Artist: "Dr. Dre", Price: 22.99},
	{ID: "21", Title: "Appetite for Destruction", Artist: "Guns N' Roses", Price: 29.99},
	{ID: "22", Title: "The Rise and Fall of Ziggy Stardust and the Spiders from Mars", Artist: "David Bowie", Price: 35.99},
	{ID: "23", Title: "Legend", Artist: "Bob Marley & The Wailers", Price: 38.99},
	{ID: "24", Title: "Nevermind", Artist: "Nirvana", Price: 30.99},
	{ID: "25", Title: "Hunky Dory", Artist: "David Bowie", Price: 34.99},
}

// CONTROLADOR PARA DEVOLVER LOS DATOS - RESPONDER AL CLIETNE
func getAlbums(c *gin.Context) {
	// RESPONDER AL CLIENTE CON LA DATA SERILIALIZADA A TIPO JSON Y CON ESTATUS HTTP
	c.IndentedJSON(http.StatusOK, albums) // (estatus, data)

}

func main() {
	fmt.Println("Bienvenido a Vinil-API")

	//CREAR UNA INSTANCIA DE GIN
	router := gin.Default()

	//CREAR UNA RUTA
	// router.GET("/", func(c *gin.Context) { // gin.context : recibir y responder las peticiones
	// 	//Al utilizar esta ruta, vamos a contestar con un JSON
	// 	c.JSON(200, gin.H{ //JSON(estatus, cuerpo del mensaje) - Utilizar las constantes de net/http para los status
	// 		"message": "Hola bb",
	// 	})
	// })

	// CUANDO HAGAMOS LA PETICION A ESA RUTA, EJECUTAREMOS EL HANDLER
	router.GET("/albums", getAlbums)

	//LEVANTAR EL SERVIDOR
	router.Run(":8080") // Run(puerto)
}
