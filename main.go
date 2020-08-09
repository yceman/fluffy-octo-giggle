package main

import (
            "html/template"
            "fmt"
	    "io/ioutil"
	    "log"
	    "net/http"
)

type Page struct {
       Title string
       Body []byte
}

//Save page = persistent storage

func (p *Page)  save() error {
      filename := p.Title + ".txt"
      return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
      filename := title + ".txt"
      body, err := ioutil.ReadFile(filename)
      if err != nil {
          return nil, err
      }
      return &Page{Title: title, Body: body}, nil
}

/*func renderTemplate(w http.ResponseWriter, tmpl string, p *page) {
            t, _ := template.ParseFiles(tmpl + ".html")
	    t.Execute(w, p)
}*/

func viewHandler(w http.ResponseWriter, r *http.Request) {
             title := r.URL.Path[len("/view/") : ]
	     p, _ := loadPage(title)
	     if err != nil {
	          http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		  return
	     }
	     renderTemplate(w, "view" , p)
	     /*fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)*/
}

func editHandler(w http.ResponseWriter, r *http.Request) {
             title := r.URL.PATH[ ("edit") : ]
	     p, err := loadPage(title)
	     if err != nil {
	          p = &Page{Title: title}
	     }
	     renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
       title := r.URL.Path[len("/save/") : ]
       body := r.FormValue("body")
       p := &Page{Title: title, Body: []byte(body)}
       err := p.save()
       if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
       }
       http.Redirect(w, r "/view/"+title, http.Status.Found)
}

func main() {
       http.HandleFunc("/view/", viewHandler)
       http.HandleFunc("/edit/", editHandler)
       //http.HandleFunc("/save/", saveHandler)
       log.Fatal(http.ListenAndServe(":8080", nil))
       p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
       p1.save()
       p2, _ := loadPage("TestPage")
       fmt.Println(string(p2.Body))
}

