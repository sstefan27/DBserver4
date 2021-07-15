package rest

import (
	"encoding/json"
	"fmt"
	"go/problem4/entity"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Login(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body
	bodyBytes, _ := ioutil.ReadAll(reqBody)
	var credentials entity.Credentials

	json.Unmarshal(bodyBytes, &credentials)
	session, _ := store.Get(r, "cookie-name")
	session.ID = credentials.ID
	session.Values["userID"] = credentials.ID
	session.Values["authenticated"] = true
	session.Save(r, rw)
}
func Welcome(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	session, _ := store.Get(r, "cookie-name")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(rw, "acces forbidden", http.StatusForbidden)
		return
	}
	fmt.Fprintln(rw, "Welcome ", name)
}

func Logout(rw http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Save(r, rw)
}
