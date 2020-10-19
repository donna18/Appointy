package main

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
	"time"
	"html/template"
	//"os"
    
)

type Meeting struct {
    Id int `json:"id"`
    Title string `json:"title"`
	Participant string `json:"participant"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
}

type Meetings []Meeting

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	http.HandleFunc("/allmeetings", allmeetings)
	http.HandleFunc("/login",login)
    log.Fatal(http.ListenAndServe(":10000", nil))
}



func allmeetings(w http.ResponseWriter, r *http.Request){
	meetings:=Meetings{
		
	}
	fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        y, _ := template.ParseFiles("allmeetings.gtpl")
		y.Execute(w, nil)
		fmt.Println("ID:", r.Form["id"])
		fmt.Println("Title:", r.Form["title"])
		fmt.Println("Participant:", r.Form["participant"])
		fmt.Println("Start Time:", r.Form["start"])
		fmt.Println("End Time:", r.Form["end"])
    } else {
        r.ParseForm()
        // logic part of log in
        
		fmt.Println("ID:", r.Form["id"])
		fmt.Println("Title:", r.Form["title"])
		fmt.Println("Participant:", r.Form["participant"])
		fmt.Println("Start Time:", r.Form["start"])
		fmt.Println("End Time:", r.Form["end"])

        
    }
	fmt.Println("Endpoint Hit: returnAllMeetings")
	json.NewEncoder(w).Encode(meetings)
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
		fmt.Println("Hi username:", r.Form["username"])
        fmt.Println("Your password:", r.Form["password"])
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		/*fmt.Println("ID:", r.Form["id"])
		fmt.Println("Title:", r.Form["title"])
		fmt.Println("Participant:", r.Form["participant"])
        */
    }
}

func main() {

    handleRequests()
}