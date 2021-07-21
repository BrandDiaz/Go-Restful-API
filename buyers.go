package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type buyersResource struct{}

func (rs buyersResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)

	r.Route("/{date}", func(r chi.Router){
		r.Use(PostContext)
		r.Get("/", rs.Get)
	})

	return r
}

func (rs buyersResource) List(w http.ResponseWriter, r *http.Request){
	resp, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _,err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostContext(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		context := context.WithValue(r.Context(), "date", chi.URLParam(r, "date"))
		next.ServeHTTP(w, r.WithContext(context))
	})
}

func (rs buyersResource) Get(w http.ResponseWriter, r *http.Request) {
  date := r.Context().Value("date").(string)

  if len(date) < 1 {
	date = strconv.FormatInt(time.Now().Unix(), 10)
  }

  resp, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers?date=" + date)
  fmt.Println(resp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  defer resp.Body.Close()

  w.Header().Set("Content-Type", "application/json")

  if _, err := io.Copy(w, resp.Body); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
