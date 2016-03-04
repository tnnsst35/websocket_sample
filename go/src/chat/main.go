package main

import (
    "log"
    "net/http"
    "text/template"
    "path/filepath"
    "sync"
    "flag"
    "trace"
    "os"
    "github.com/stretchr/gomniauth"
    "github.com/stretchr/gomniauth/providers/facebook"
    "github.com/stretchr/gomniauth/providers/github"
    "github.com/stretchr/gomniauth/providers/google"
)

type templateHandler struct {
    once sync.Once
    filename string
    templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    t.once.Do(func() {
        t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
    })
    t.templ.Execute(w, r)
}

const (
    authSecurityKey = ""
    facebookClientId = ""
    facebookPrivateKey = ""
    githubClientId = ""
    githubPrivateKey = ""
    googleClientId = ""
    googlePrivateKey = ""
)

func main() {
    /* Step1
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`
            <html>
                <head>
                    <title>チャット</title>
                </head>
                <body>
                    チャットしましょう！
                </body>
            </html>
        `))
    })
    */
    var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
    flag.Parse()
    gomniauth.SetSecurityKey(authSecurityKey);
    gomniauth.WithProviders(
        facebook.New(facebookClientId, facebookPrivateKey, "http://localhost:8080/auth/callback/facebook"),
        github.New(githubClientId, githubPrivateKey, "http://localhost:8080/auth/callback/github"),
        google.New(googleClientId, googlePrivateKey, "http://localhost:8080/auth/callback/google"),
    );
    r := newRoom()
    r.tracer = trace.New(os.Stdout)
    http.Handle("/login", &templateHandler{filename: "login.html"})
    http.HandleFunc("/auth/", loginHandler)
    http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
    http.Handle("/room", r)
    go r.run()
    // Webサーバーを開始します
    // log.Println("Webサーバーを開始します　ポート：", *addr)
    if err := http.ListenAndServe(*addr, nil); err != nil {
        log.Fatal("ListenAndServe", err)
    }
}