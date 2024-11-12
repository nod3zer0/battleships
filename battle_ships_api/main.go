package main

//TODO split into multiple files

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
	Username string `json:"username"`
	Password string `json:"password"`
	Session  string `json:"session"`
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

type ship_db struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type part_db struct {
	ID     string `json:"id"`
	ShipID string `json:"shipid"`
	Pos_x  int    `json:"pos_x"`
	Pos_y  int    `json:"pos_y"`
}

type ship_api struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Rotate   int     `json:"rotate"`
	Position []int   `json:"position"`
	Parts    [][]int `json:"parts"`
}

type ship_game_api struct {
	Ship_id   string `json:"id"`
	Player_id string `json:"player_id"`
	Game_id   string `json:"game_id"`
	Rotate    int    `json:"rotate"`
	Position  []int  `json:"position"`
}

type ship_game struct {
	ID        string `json:"id"`
	Ship_id   string `json:"ship_id"`
	Player_id string `json:"player_id"`
	Game_id   string `json:"game_id"`
	Turn      int    `json:"turn"`
	Pos_x     int    `json:"pos_x"`
	Pos_y     int    `json:"pos_y"`
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
	router.GET("/ships", getShips)
	router.POST("/create_game", createGame)
	router.POST("/save_ships", addShips)
	router.GET("/get_games", getGames)

	router.Run("localhost:8081")
}

type game_db struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func getGames(c *gin.Context) {

	row, err := db.Query("SELECT game.id, game.name, user.username FROM game LEFT JOIN user ON game.p1_id = user.id WHERE game.p2_id is null")

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "No games found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	defer row.Close()

	var games []game_db
	for row.Next() {
		var game game_db
		if err := row.Scan(&game.ID, &game.Name, &game.Username); err != nil {
			log.Fatal(err)
		}
		games = append(games, game)
	}

	c.IndentedJSON(http.StatusOK, games)

}

func createGame(c *gin.Context) {

	var session userSession
	if err := c.BindJSON(&session); err != nil {
		return
	}

	//get user id from session
	var u user
	err := db.QueryRow("SELECT id FROM user WHERE session = ?", session.Sessionid).Scan(&u.ID)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println(u)
	//create game
	db.Exec("INSERT INTO game (name,p1_id) VALUES ('test',?)", u.ID)

}

func addShips(c *gin.Context) {

	type addShipsAPI struct {
		Game_id      string          `json:"id"`
		User_session string          `json:"user_session"`
		Ships        []ship_game_api `json:"ships"`
	}

	var data addShipsAPI
	if err := c.BindJSON(&data); err != nil {
		return
	}

	//get user id from session
	var u user
	err := db.QueryRow("SELECT id FROM user WHERE session = ?", data.User_session).Scan(&u.ID)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println(u)
	log.Println(data)

	for _, ship := range data.Ships {
		_, err := db.Exec("INSERT INTO ship_game (ship_id,player_id,game_id,turn,pos_x,pos_y) VALUES (?,?,?,?,?,?)", ship.Ship_id, u.ID, data.Game_id, ship.Rotate, ship.Position[0], ship.Position[1])
		if err != nil {
			log.Fatal(err)
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Ships added"})
}

func getShips(c *gin.Context) {
	var ships []ship_db
	rows, err := db.Query("SELECT * FROM ship")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var ship ship_db
		if err := rows.Scan(&ship.ID, &ship.Name); err != nil {
			log.Fatal(err)
		}
		ships = append(ships, ship)
	}

	var parts []part_db
	rows, err = db.Query("SELECT * FROM ship_part")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var part part_db
		if err := rows.Scan(&part.ID, &part.ShipID, &part.Pos_x, &part.Pos_y); err != nil {
			log.Fatal(err)
		}
		parts = append(parts, part)
	}

	var ships_api []ship_api
	for _, ship := range ships {
		var parts_api [][]int
		for _, part := range parts {
			if part.ShipID == ship.ID {
				parts_api = append(parts_api, []int{part.Pos_x, part.Pos_y})
			}
		}
		ships_api = append(ships_api, ship_api{ID: ship.ID, Name: ship.Name, Parts: parts_api, Rotate: 0, Position: []int{0, 0}})
	}

	c.IndentedJSON(http.StatusOK, ships_api)
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
