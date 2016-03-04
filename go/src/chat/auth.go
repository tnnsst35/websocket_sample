package main

import (
    "net/http"
    "strings"
    "log"
    "github.com/stretchr/gomniauth"
)

type authHandler struct {
    next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
        w.Header().Set("Location", "/login")
        w.WriteHeader(http.StatusTemporaryRedirect)
    } else if err != nil {
        panic(err.Error())
    } else {
        h.next.ServeHTTP(w, r)
    }
}

func MustAuth(handler http.Handler) http.Handler {
    return &authHandler{next: handler}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    segs := strings.Split(r.URL.Path, "/")
    action := segs[2]
    provider := segs[3]
    switch action {
    case "login":
        provider, err := gomniauth.Provider(provider)
        if err != nil {
            log.Fatalln(provider, "-", err)
        }
        loginUrl, err := provider.GetBeginAuthURL(nil, nil)
        if err != nil {
            log.Fatalln(provider, "-", err)
        }
        w.Header().Set("Location", loginUrl)
        w.WriteHeader(http.StatusTemporaryRedirect)
    default:
        w.WriteHeader(http.StatusNotFound)
    }
}