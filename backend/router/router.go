package router

import (
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"

	"database/sql"
	"fmt"
	"os"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func Router() {
	e := echo.New()

	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3030"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}
	e.Use(middleware.CORSWithConfig(corsConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/user", getUser)
	e.GET("/room", getRoom)
	e.GET("/chat", getChat)
	e.POST("/chat", postChat)
	e.PUT("/chat", editChat)
	e.DELETE("/chat", deleteChat)

	e.Logger.Fatal(e.Start(":8080"))

}

type User struct {
	ID       string `json:id`
	UserName string `json:user_name`
	Password string `json:password`
	Email    string `json:email`
}

type Room struct {
	RoomID   string `json:room_id`
	RoomName string `json:room_name`
}

type Chat struct {
	ChatID    string `json:chat_id`
	RoomID    string `json:room_id`
	UserID    string `json:user_id`
	Text      string `json:text`
	IsEditted bool   `json:is_editted`
}

func getUser(c echo.Context) error {
	// dbconf := "test_user:password@tcp(localhost:3306)/test_database?charset=utf8mb4"
	user := os.Getenv("MYSQL_USER")  
	pass := os.Getenv("MYSQL_PASSWORD")  
	host := os.Getenv("INSTANCE_CONNECTION_NAME")  
	name := os.Getenv("MYSQL_DATABASE")  
	 
	dbconf := fmt.Sprintf("%s:%s@%s/%s", user, pass, host, name)  

	db, err := sql.Open("mysql", dbconf)

	// Userをとってくる
	users, err := db.Query("SELECT * FROM user;")

	if err != nil {
		log.Fatalf("getUsers db.Query error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	defer users.Close()

	// for users.Next() {
	// 	u := &User{}
	// 	if err := users.Scan(&u.ID, &u.UserName, &u.Password, &u.Email); err != nil {
	// 		log.Fatalf("getUsers users.Scan error err:%v", err)
	// 	}
	// 	fmt.Println(u)
	// 	return c.JSON(http.StatusOK, u)
	// }

	res := []User{}
	for users.Next() {
		u := &User{}
		if err := users.Scan(&u.ID, &u.UserName, &u.Password, &u.Email); err != nil {
			log.Fatalf("getUsers users.Scan error err:%v", err)
		}
		res = append(res, *u)
	}
	return c.JSON(http.StatusOK, res)

	err = users.Err()
	if err != nil {
		log.Fatalf("getUsers users.Err error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "OK"})
}

func getRoom(c echo.Context) error {
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]



	user := os.Getenv("MYSQL_USER")  
	pass := os.Getenv("MYSQL_PASSWORD")  
	host := os.Getenv("INSTANCE_CONNECTION_NAME")  
	name := os.Getenv("MYSQL_DATABASE")  
	 
	dbconf := fmt.Sprintf("%s:%s@%s/%s", user, pass, host, name)  

	

	// dbconf := "uttc:hello0325@unix(/cloudsql/term0-hinata-takeda:us-central1:uttc-hina)/hackathon"

	// dbconf := "test_user:password@tcp(localhost:3306)/test_database?charset=utf8mb4"

	db, err := sql.Open("mysql", dbconf)

	if err != nil {
		log.Fatalf("getRoom db.Open error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// Roomをとってくる
	rooms, err := db.Query("SELECT * FROM room;")

	if err != nil {
		log.Fatalf("getRoom db.Query error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// 複数配列
	res := []Room{}
	for rooms.Next() {
		r := &Room{}
		if err := rooms.Scan(&r.RoomID, &r.RoomName); err != nil {
			log.Fatalf("getRooms rooms.Scan error err:%v", err)
		}
		res = append(res, *r)
	}

	return c.JSON(http.StatusOK, res)

	err = rooms.Err()
	if err != nil {
		log.Fatalf("getRooms rooms.Err error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "OK"})
}

// RoomIDで引っ掛けてチャットを取ってくる
func getChat(c echo.Context) error {
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	

	// dbconf := "test_user:password@tcp(localhost:3306)/test_database?charset=utf8mb4"
	user := os.Getenv("MYSQL_USER")  
	pass := os.Getenv("MYSQL_PASSWORD")  
	host := os.Getenv("INSTANCE_CONNECTION_NAME")  
	name := os.Getenv("MYSQL_DATABASE")  
	 
	dbconf := fmt.Sprintf("%s:%s@%s/%s", user, pass, host, name)  



	db, err := sql.Open("mysql", dbconf)

	if err != nil {
		log.Fatalf("getChat db.Open error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// RoomIDで引っ掛けてチャットを取ってくる
	roomID := c.QueryParam("roomID")

	chats, err := db.Query("SELECT * FROM chat WHERE roomid = ?;", roomID)

	if err != nil {
		log.Fatalf("getChat db.Query error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// 複数配列
	res := []Chat{}
	for chats.Next() {
		c := &Chat{}
		if err := chats.Scan(&c.ChatID, &c.RoomID, &c.UserID, &c.Text, &c.IsEditted); err != nil {
			log.Fatalf("getChat chats.Scan error err:%v", err)
		}
		res = append(res, *c)
	}

	return c.JSON(http.StatusOK, res)

	err = chats.Err()
	if err != nil {
		log.Fatalf("getChat chats.Err error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "OK"})
}

func postChat(c echo.Context) error {
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	// dbconf := "uttc:hello0325@unix(/cloudsql/term0-hinata-takeda:us-central1:uttc)/hackathon"

	// dbconf := "test_user:password@tcp(localhost:3306)/test_database?charset=utf8mb4"
	user := os.Getenv("MYSQL_USER")  
	pass := os.Getenv("MYSQL_PASSWORD")  
	host := os.Getenv("INSTANCE_CONNECTION_NAME")  
	name := os.Getenv("MYSQL_DATABASE")  
	 
	dbconf := fmt.Sprintf("%s:%s@%s/%s", user, pass, host, name)  


		// db, err := sql.Open("mysql", dbconf)

	if err != nil {
		log.Fatalf("postChat db.Open error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// chatIDで引っ掛けてチャットを取ってくる
	roomID := c.QueryParam("roomID")
	userID := c.QueryParam("userID")
	chatText := c.QueryParam("chatText")
	IsEditted := false
	_, err = db.Query("INSERT INTO chat (roomid, id, text, retext) VALUES (?, ?, ?, ?);", roomID, userID, chatText, IsEditted)

	if err != nil {
		log.Fatalf("postChat db.Query error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "OK"})
}

func editChat(c echo.Context) error {
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	// dbconf := "uttc:hello0325@unix(/cloudsql/term0-hinata-takeda:us-central1:uttc)/hackathon"

	// dbconf := "test_user:password@tcp(localhost:3306)/test_database?charset=utf8mb4"

	// db, err := sql.Open("mysql", dbconf)

	user := os.Getenv("MYSQL_USER")  
	pass := os.Getenv("MYSQL_PASSWORD")  
	host := os.Getenv("INSTANCE_CONNECTION_NAME")  
	name := os.Getenv("MYSQL_DATABASE")  
	 
	dbconf := fmt.Sprintf("%s:%s@%s/%s", user, pass, host, name)  
	
	if err != nil {
		log.Fatalf("editChat db.Open error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// chatIDで引っ掛けてチャットを取ってくる
	chatID := c.QueryParam("chatID")
	intChatID, _ := strconv.Atoi(chatID)
	chatText := c.QueryParam("chatText")
	isEditted := true
	_, err = db.Query("UPDATE chat SET text = ?, retext = ? WHERE chatid = ?;", chatText, isEditted, intChatID)

	if err != nil {
		log.Fatalf("editChat db.Query error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "OK"})
}

// チャットを削除する
func deleteChat(c echo.Context) error {
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	// dbconf := "uttc:hello0325@unix(/cloudsql/term0-hinata-takeda:us-central1:uttc)/hackathon"
	user := os.Getenv("MYSQL_USER")  
	pass := os.Getenv("MYSQL_PASSWORD")  
	host := os.Getenv("INSTANCE_CONNECTION_NAME")  
	name := os.Getenv("MYSQL_DATABASE")  
	 
	dbconf := fmt.Sprintf("%s:%s@%s/%s", user, pass, host, name)  



	db, err := sql.Open("mysql", dbconf)

	if err != nil {
		log.Fatalf("deleteChat db.Open error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// chatIDで引っ掛けてチャットを取ってくる
	chatID := c.QueryParam("chatID")
	intChatID, _ := strconv.Atoi(chatID)
	_, err = db.Query("DELETE FROM chat WHERE chatid = ?;", intChatID)

	if err != nil {
		log.Fatalf("deleteChat db.Query error err:%v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "OK"})
}
