package album_handler

import (
	album_model "EdmundsBankai/golang-intro/gin-tutorial/models"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateAlbumFields(album album_model.Album) {
	if album.ID == "" {
		panic("Missing id field")
	}
	if album.Title == "" {
		panic("Missing title")
	}
	if album.Artist == "" {
		panic("Missing artist")
	}
	if album.Price == 0 {
		panic("Nothing in life is free")
	}
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, album_model.Albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum album_model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	ValidateAlbumFields(newAlbum)
	for _, album := range album_model.Albums {
		if newAlbum.ID == album.ID {
			log.Fatal("album with id already exists")
			return
		}
	}
	album_model.Albums = append(album_model.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func DeleteAlbum(c *gin.Context, id string) {
	if id == "" {
		log.Fatal("¿Cuál?")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var newAlbums []album_model.Album
	var deleted album_model.Album
	for _, album := range album_model.Albums {
		if album.ID != id {
			newAlbums = append(newAlbums, album)
			continue
		}
		deleted = album
	}
	album_model.Albums = newAlbums
	c.IndentedJSON(http.StatusOK, deleted)
}

func UpdateAlbum(c *gin.Context) {
	var album2Delete album_model.Album
	if album2Delete.ID == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := c.BindJSON(&album2Delete); err != nil {
		return
	}
	DeleteAlbum(c, album2Delete.ID)
	PostAlbum(c)
	c.IndentedJSON(http.StatusOK, album2Delete.ID)
}

func PatchAlbum(c *gin.Context) []album_model.Album {
	var updates album_model.Album

	if err := c.BindJSON(&updates); err != nil {
		log.Fatal("Body is not of type Album")
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if updates.ID == "" {
		log.Fatal("Missing id field")
		c.AbortWithStatus(http.StatusBadRequest)
	}

	var album *album_model.Album      // found album
	var albums = album_model.Albums   // current albums
	var newAlbums []album_model.Album // updated albums

	for _, al := range albums {
		if al.ID == updates.ID {
			album = &al
			if updates.Title != "" {
				al.Title = updates.Title
			}
			if updates.Artist != "" {
				al.Artist = updates.Artist
			}
			if updates.Price != 0 {
				al.Price = updates.Price
			}
		}
		newAlbums = append(newAlbums, al)
	}

	if album.ID == "" {
		log.Fatal("Album with id not found")
		c.AbortWithStatus(http.StatusBadRequest)
	}

	album_model.Albums = newAlbums

	return album_model.Albums
}
