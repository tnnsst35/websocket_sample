package main

import (
    "net/http"
    "strings"
    "log"
    "github.com/stretchr/gomniauth"
    gomniauthcommon "github.com/stretchr/gomniauth/common"
    "github.com/stretchr/objx"
    "io"
    "crypto/md5"
    "fmt"
)

type ChatUser interface {
    UniqueID() string
    AvatarURL() string
}

type chatUser struct {
    gomniauthcommon.User
    uniqueID string
}

func (u chatUser) UniqueID() string {
    return u.uniqueID
}

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
    case "callback":
        provider, err := gomniauth.Provider(provider)
        if err != nil {
            log.Fatalln(provider, "-", err)
        }
        creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
        if err != nil {
            log.Fatalln(provider, "-", err)
        }
        user, err := provider.GetUser(creds)
        if err != nil {
            log.Fatalln(provider, "-", err)
        }

        chatUser := &chatUser{User: user}
        m := md5.New()
        io.WriteString(m, strings.ToLower(user.Email()))
        chatUser.uniqueID = fmt.Sprintf("%x", m.Sum(nil))
        avatarURL, err := avatars.GetAvatarURL(chatUser)
        if err != nil {
            log.Fatalln("GetAvatarURLに失敗しました",  "-", err)
        }
        authCookieValue := objx.New(map[string]interface{}{
            "userid": chatUser.uniqueID,
            "name": user.Name(),
            "avatar_url": avatarURL,
        }).MustBase64()
        http.SetCookie(w, &http.Cookie{
            Name: "auth",
            Value: authCookieValue,
            Path: "/"})
        w.Header()["Location"] = []string{"/chat"}
        w.WriteHeader(http.StatusTemporaryRedirect)
    default:
        w.WriteHeader(http.StatusNotFound)
    }
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{
        Name: "auth",
        Path: "/",
        MaxAge: -1})
    w.Header().Set("Location", "/login")
    w.WriteHeader(http.StatusTemporaryRedirect)
}
