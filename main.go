package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"shortlinks/models"
)

var db *sql.DB

func main() {

	db = connect()
	defer db.Close()

	fmt.Println("Listening on port 3000")

	http.HandleFunc("/", redirect)
	http.HandleFunc("/auth/", auth)
	http.HandleFunc("/add-link/", addLink)
	http.HandleFunc("/get-popular/", getPopular)
	http.HandleFunc("/get-excel/", getExcel)

	http.ListenAndServe(":3000", nil)

}

func getExcel(w http.ResponseWriter, r *http.Request) {
}
func getPopular(w http.ResponseWriter, r *http.Request) {

	table, err := models.GetPopularTable(db)
	if err != nil {
		fmt.Fprintf(w, "{error:can't get popular links}")
		return
	}

	fmt.Fprintf(w, "<table>")
	for _, v := range table {
		fmt.Fprintf(w, fmt.Sprintf("%v\n", v.GetString()))
	}
	fmt.Fprintf(w, "</table>")
}

func redirect(w http.ResponseWriter, r *http.Request) {

	shortLink := r.URL.Query().Get("r")
	if shortLink == "" {
		fmt.Fprintf(w, "{error:paremeter r is required}")
		return

	}

	fullLink, err := models.LinkGetFull(shortLink, db)

	if err != nil {
		fmt.Fprintf(w, "{error:can't get the full link}")
		return
	}

	if fullLink == "" {
		fmt.Fprintf(w, "{error:full is empty}")
		return
	}

	fmt.Printf("Redirect to %s\n", fullLink)

	statErr := models.RecordFollowing(shortLink, fullLink, db)
	if statErr != nil {
		fmt.Println("Didn't record the folowing into statistics")
	}

	http.Redirect(w, r, fullLink, 302)

}

func addLink(w http.ResponseWriter, r *http.Request) {

	var link models.Link

	formShort := r.FormValue("short")
	formFull := r.FormValue("full")
	formToken := r.FormValue("token")
	formExpirationTime := r.FormValue("expiration_time")
	expirationTimeInt, err := strconv.Atoi(formExpirationTime)

	if formExpirationTime != "" && err != nil {
		fmt.Fprintf(w, "{error:can't parse the date}")
		return
	}

	if formToken == "" {
		fmt.Fprintf(w, "{error:token is required}")
		return
	}

	if formFull == "" {
		fmt.Fprintf(w, "{error:full is required}")
		return
	}

	token, err := models.TokenCheck(formToken, db)

	if err != nil {
		fmt.Fprintf(w, "{error:something is wrong with token")
		return
	}

	if token.Ownerid == 0 {
		fmt.Fprintf(w, "{error:token doesn't exist or expired")
		return
	}

	if expirationTimeInt > 0 {
		link.ExpirationTime = time.Now().AddDate(0, 0, expirationTimeInt)
	}

	link.Ownerid = token.Ownerid
	link.Short = formShort
	link.Full = formFull

	linkErr := models.LinkAdd(link, db)

	if linkErr != nil {
		fmt.Fprintf(w, "{error:can't create a new short link")
		return
	}

	fmt.Printf("link = %v\n", link)

	newShortLink, shortLinkErr := models.LinkGetShort(link.Full, db)

	if shortLinkErr != nil {
		fmt.Fprintf(w, "{error:can't get the new short link}")
		return
	}

	if newShortLink == "" {
		fmt.Fprintf(w, "{error:something wrong, have to solve this problem...}")
		return
	}

	fmt.Fprint(w, fmt.Sprintf("{short_link:http://localhost:3000/?r=%v}", newShortLink))

}

func auth(w http.ResponseWriter, r *http.Request) {

	login := r.FormValue("login")
	password := r.FormValue("password")

	if login == "" || password == "" {
		fmt.Fprintf(w, "{error:login and password are required}")
		return
	}

	checkAuthRes := models.UserCheckAuth(login, password, db)

	if checkAuthRes != true {
		fmt.Fprintf(w, "{error:login or password is incorrect}")
		return
	}

	user := models.UserGetByLogin(login, db)

	t, err := models.TokenGetNew(user, db)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "{error:token can't be recieved}")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("{token:%v}", t.Token))

}
