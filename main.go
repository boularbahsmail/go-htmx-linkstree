package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Link struct {
	Icon  string
	Title string
	Url   string
}

type User struct {
	Avatar   string
	FullName string
	Bio      string
	Links    map[string][]Link
}

const avatar string = "https://pbs.twimg.com/profile_images/1746251896124219392/MGtCrEuI_400x400.jpg"
const fullName string = "Ismail Boularbah"
const bio string = "Hope is no strategy, Writing @typescript and @golang"

var links = map[string][]Link{
	"link": {
		{
			Icon:  "https://cdn-icons-png.flaticon.com/128/1051/1051326.png",
			Title: "GitHub",
			Url:   "https://github.com/boularbahsmail",
		},
		{
			Icon:  "https://cdn-icons-png.flaticon.com/128/5968/5968958.png",
			Title: "Twitter/X",
			Url:   "https://twitter.com/boularbahsmail",
		},
		{
			Icon:  "https://cdn-icons-png.flaticon.com/128/1006/1006771.png",
			Title: "Official Website",
			Url:   "https://ismailium.vercel.app",
		},
	},
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	// Rendering and serving HTML template
	tmpl := template.Must(template.ParseFiles("./index.html"))

	user := User{
		Avatar:   avatar,
		FullName: fullName,
		Bio:      bio,
		Links:    links,
	}

	tmpl.Execute(writer, user)
}

func main() {
	fmt.Println("Server running on PORT 8000...")

	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
