// // package chat

// // import (
// // 	"database/sql"
// // 	"fmt"
	
// // )





// package main
 
// import (
// 	"net/http"
// 	"backend/models"
// 	"net/http"
// 	"fmt"
// 	"database/sql"
// 	_ "github.com/go-sql-driver/mysql"
// 	"encoding/json"
// 	"log"
// 	"github.com/julienschmidt/httprouter"
// )

 
// func (app *application) routes() http.Handler {
// 	router := httprouter.New()
	
// 	router.HandlerFunc(http.MethodGet, "/general", app.getGeneralTalks)
// 	router.HandlerFunc(http.MethodGet, "/random", app.getRandomTalks)
// 	router.HandlerFunc(http.MethodPost, "/login", app.loginHandler)

// 	router.HandlerFunc(http.MethodPost, "/new", app.postGeneral)
// 	router.HandlerFunc(http.MethodPost, "/delete", app.deleteGeneral)
// 	router.HandlerFunc(http.MethodPost, "/edit", app.editGeneral)
 
// 	return app.enableCORS(router)
// }



 
// func (app *application) enableCORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Headers", "*")
//    		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
 
// 		next.ServeHTTP(w, r)
// 	})
// }


 
// func (app *application) getLoginUsers(w http.ResponseWriter, r *http.Request) {

// 	type UserList []models.User
// 	var ul UserList
	

// 	db, err := sql.Open("mysql", "root:root@tcp(mysql_container:3306)/react-go-app?charset=utf8mb4")

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer db.Close()

// 	rows, err := db.Query("SELECT id,user FROM users")

// 	if err != nil {
// 		fmt.Println("データベース接続失敗")
// 		panic(err.Error())
// 	} else {
// 		fmt.Println("データベース接続成功")
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var bar models.User
// 		err := rows.Scan(&bar.ID, &bar.User)
// 		ul = append(ul, bar)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}
// 	err = rows.Err()

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	app.writeJSON(w, http.StatusOK, ul, "user")
// }


// type Chat struct {
// 	chatid    int
// 	roomid    string
// 	id  int
// 	text string}


// type Room struct {
// 	roomid    int
// 	roomname string}

		
// type User struct {
// 	id    int
// 	username string
// 	password string
// 	email string}

	
// func ReadAll(db *sql.DB) {
// 	var chat []Chat
// 	rows, err := db.Query("select * from chat;")
// 	if err != nil {
// 		panic(err)
// 	}
// 	for rows.Next() {
// 		chat := Chat{}
// 		err = rows.Scan(&chat.chatid, &chat.roomid, &chat.id, &chat.text)
// 		if err != nil {
// 			panic(err)
// 		}
// 		chats :=  append(chats, chat)
// 	}
// 	rows.Close()


// func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {

// 	var payload UserInfo
// 	err := json.NewDecoder(r.Body).Decode(&payload)
// 	if err != nil {
// 		log.Println(err)
// 		app.errorJSON(w, err)
// 		return
// 	}


 
// 	db, err := sql.Open("mysql", "root:root@tcp(mysql_container:3306)/react-go-app?parseTime=true")

// 	if err != nil {
// 		panic(err.Error())
// 	} 

// 	var bar UserData


// 	err = db.QueryRow("select pass,id,user from users where email = ?",payload.Email).Scan(&bar.Pass,&bar.Id,&bar.User)
//     if err != nil {
//         log.Fatal(err)
//     }
// 	defer db.Close()

// 	status := models.Status{
// 	}
// 	user := models.User {


// 	}

// 	if bar.Pass == payload.Pass {
// 		status = models.Status{
// 			Status:          true,
// 		}
// 		user = models.User{
// 			ID:bar.Id,
// 			User: bar.User,
// 		}
// 	}else {
// 		status = models.Status {
// 			Status:          false,
// 		}
// 	}
	

	
// 	type ReturnData struct {
// 		Status       models.Status   `json:"status"`
// 		User         models.User   `json:"user`
// 	}
// 	returns := ReturnData{
// 		Status:status,
// 		User:user,
// 	}

// 	err = app.writeJSON(w, http.StatusOK, returns, "status")


// }

// package chat

// import (
// 	"database/sql"
// 	"fmt"
// )

// type User struct {
// 	id    int
// 	username string
// 	passward  string
// 	email string
// }

// func ReadAll(db *sql.DB) {
// 	var Users []User
// 	rows, err := db.Query("select * from user")
// 	if err != nil {
// 		panic(err)
// 	}
// 	for rows.Next() {
// 		User := User{}
// 		err = rows.Scan(&User.id, &User.username, &User.passward, &User.email)
// 		if err != nil {
// 			panic(err)
// 		}
// 		Users = append(Users, User)
// 	}
// 	rows.Close()

// 	fmt.Println(Users)
// }