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
    r := newRoom()
    r.tracer = trace.New(os.Stdout)
    http.Handle("/", &templateHandler{filename: "chat.html"})
    http.Handle("/room", r)
    go r.run()
    // Webサーバーを開始します
    // log.Println("Webサーバーを開始します　ポート：", *addr)
    if err := http.ListenAndServe(*addr, nil); err != nil {
        log.Fatal("ListenAndServe", err)
    }
}