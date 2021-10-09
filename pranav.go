package insta

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	neturl "net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)
type Person struct {
    Name string
    Password  string
    Email string
    Id int
}
type Post struct{
	Id int
	Caption string
	ImgURL string
	time time.Time
}
func initiate(){
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/instadb")
    if err != nil {
        panic(err.Error())
    }
    create, err := db.Query("CREATE TABLE users(name varchar,password varchar,email varchar,Id int)")
    if err != nil {
        panic(err.Error())
    db.Close()
    }
}
var c=0
func newUser(w http.ResponseWriter, r *http.Request) {
    // Declare a new Person struct.
    var p Person

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/instadb")
    if err != nil {
        panic(err.Error())
    }
    p.id=c+1
    insert,err:=db.Query("INSERT into users VALUES(p.Name,p.Password,p.Email,p.Id)")
    if err != nil {
        panic(err.Error())
    }
    db.Close()
}

func getUser(w http.ResponseWriter, r *http.Request)
{
	
	var p Person
	vars:=mux.Vars(r)
	ids:=vars["Id"]
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/instadb")
	result,err:=db.Query("SELECT id,name,password,email from users where id=ids")
	err = result.Scan(&p.Id, &p.Name, &p.Password, &p.Email)
	if err!=nul
	{
		panic(err.Error())
	}
