package handler

import (
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ppworks/go-playground/asset"
	"github.com/ppworks/go-playground/helper"
	"github.com/ppworks/go-playground/posts"
)

// PostNewHandleFunc is handler of "/"
func PostNewHandleFunc(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.NewRandom()
	var token string
	if err != nil {
		token = ""
	}
	token = id.String()

	crsfTokenCookie := http.Cookie{
		Name:     "csrf_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &crsfTokenCookie)

	t := template.Must(template.New("").Funcs(helper.AppHelper()).ParseFiles(
		"templates/layouts/application.html",
		"templates/posts/new.html"))

	t.ExecuteTemplate(w, "layout", struct {
		asset.ManifestFile
		BodyCSS   string
		CSRFToken string
	}{
		asset.ManifestFile{
			AppJS:  asset.Path("app.js"),
			AppCSS: asset.Path("app.css"),
		}, "", token,
	})
}

// PostsHandleFunc is hanlder of "/posts"
func PostsHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createHandleFunc(w, r)
	case http.MethodGet:
		postsHandleFunc(w, r)
	default:
		http.NotFound(w, r)
		return
	}
}

func createHandleFunc(w http.ResponseWriter, r *http.Request) {
	csrfTokenCookie, err := r.Cookie("csrf_token")

	if err != nil {
		w.WriteHeader(422)
		return
	}

	if csrfTokenCookie.Value != r.PostFormValue("csrf_token") {
		println("Can't verify CSRF token authenticity")
		w.WriteHeader(422)
		return
	}

	crsfTokenCookie := http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0),
	}

	http.SetCookie(w, &crsfTokenCookie)

	post := posts.NewPost()
	post.Author = r.PostFormValue("author")
	post.Content = r.PostFormValue("content")
	err = post.Upsert()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	http.Redirect(w, r, "/posts/", 301)
}

func postsHandleFunc(w http.ResponseWriter, r *http.Request) {
	data, err := posts.Posts(0, 100)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	t := template.Must(template.New("").Funcs(helper.AppHelper()).ParseFiles(
		"templates/layouts/application.html",
		"templates/posts/index.html"))

	t.ExecuteTemplate(w, "layout", struct {
		asset.ManifestFile
		BodyCSS string
		Posts   []*posts.Post
	}{
		asset.ManifestFile{
			AppJS:  asset.Path("app.js"),
			AppCSS: asset.Path("app.css"),
		}, "", data,
	})
}
