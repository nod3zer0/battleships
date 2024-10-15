package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"crypto/rand"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type user struct {
	ID       string `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
	session  string `json:"session"`
}

type userData struct {
	Username string `json:"username"`
}

type userSession struct {
	Sessionid string `json:"sessionid"`
}

type userAuth struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var db *sql.DB

func main() {

	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "battleships",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "battleships",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/addUser", addUser)
	router.POST("/user", GetUser)
	router.POST("/login", Login)
	router.POST("/signup", SignUp)

	router.Run("localhost:8081")
}

func GetUser(c *gin.Context) {
	var session userSession
	if err := c.BindJSON(&session); err != nil {
		return
	}
	log.Println(session)
	var u userData
	err := db.QueryRow("SELECT username FROM user WHERE session = ?", session.Sessionid).Scan(&u.Username)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, u)
}

func addUser(c *gin.Context) {
	db.Exec("INSERT INTO user (username, password) VALUES (?, ?)", "test", "test")
	return
}

func generateSessionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func Login(c *gin.Context) {
	var u userAuth

	if err := c.BindJSON(&u); err != nil {
		return
	}

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT username FROM user WHERE user.username = ? AND user.password = ?)", u.Username, u.Password).Scan(&exists)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !exists {
		c.IndentedJSON(http.StatusConflict, gin.H{"error": "User does not exist"})
		return
	}

	session_id := generateSessionID()
	db.Exec("UPDATE user SET session = ? WHERE user.username = ? AND user.password = ?", session_id, u.Username, u.Password)
	c.Header("Authorization", session_id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in"})
}

func SignUp(c *gin.Context) {
	var u userAuth

	if err := c.BindJSON(&u); err != nil {
		return
	}

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT username FROM user WHERE user.username = ?)", u.Username).Scan(&exists)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if exists {
		c.IndentedJSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	session_id := generateSessionID()
	db.Exec("INSERT INTO user (username, password, session) VALUES (?, ?, ?)", u.Username, u.Password, session_id)
	c.Header("Authorization", session_id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User created"})

}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// import "fmt"

// func main() {
//     fmt.Println("Hello, World!")
// }
