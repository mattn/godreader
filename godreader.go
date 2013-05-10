package main

import (
	"database/sql"
	"github.com/mattn/go-mobileagent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mattn/godcrawler"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type tmplValue struct {
	Root  interface{}
	Value interface{}
}

func getTmplName(req *http.Request) string {
	userAgent := req.Header.Get("User-Agent")
	if mobileagent.IsMobile(userAgent) {
		return "mobile"
	}
	return "iphone"
}

func main() {
	db, err := sql.Open("sqlite3", "./godreader.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	crawler := godcrawler.New(db)
	go crawler.Run()

	root := "/godreader/"

	fmap := template.FuncMap{
		"html":  func(text string) template.HTML {
			return template.HTML(text)
		},
	}
	tmpls := map[string]*template.Template{}
	tmpls["mobile"], err = template.New("mobile").Funcs(fmap).ParseGlob(filepath.Join(filepath.Dir(os.Args[0]), "tmpl/mobile", "*.t"))
	if err != nil {
		log.Fatal("mobile ", err.Error())
	}
	tmpls["iphone"], err = template.New("iphone").Funcs(fmap).ParseGlob(filepath.Join(filepath.Dir(os.Args[0]), "tmpl/iphone", "*.t"))
	if err != nil {
		log.Fatal("iphone ", err.Error())
	}

	http.HandleFunc(root+"assets/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/"+r.URL.Path[len(root+"asserts"):])
	})

	http.HandleFunc(root, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == root {
			entries, err := crawler.Entries(50)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			tmpls[getTmplName(r)].ExecuteTemplate(w, "entries", tmplValue{
				Root: root,
				Value: entries,
			})
		} else {
			id := r.URL.Path[len(root):]
			entry, err := crawler.Entry(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			err = tmpls[getTmplName(r)].ExecuteTemplate(w, "entry", tmplValue{
				Root: root,
				Value: entry,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
		}
	})

	http.ListenAndServe(":10089", nil)
}
