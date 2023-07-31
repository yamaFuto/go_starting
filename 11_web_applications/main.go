package main

//1 ioutil

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// type Page struct {
// 	Title string
// 	Body []byte
// }

// //pageの情報をデータベースに書き込むイメージ
// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	//0600⇒webサーバを立ち上げたものが読み書きできる
// 	return ioutil.WriteFile(filename, p.Body, 0600)
// }

// //pageの情報をデータベースから読み込むイメージ
// func loadPage(title string) (*Page, error) {
// 	filename := title + ".txt"
// 	body, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//返り値が二つ以上あるときにはカンマで区切る
// 	return &Page{Title: title, Body: body}, nil
// }

// func main() {
// 	//データベースに保存する
// 	p1 := &Page{Title: "test", Body: []byte("This is a sample Page")}
// 	p1.save()

// 	//データベースからと取り出す
// 	p2, _ := loadPage(p1.Title)
// 	fmt.Println(string(p2.Body))
// }

//2 http.ListenAndServer

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// type Page struct {
// 	Title string
// 	Body []byte
// }

// //pageの情報をデータベースに書き込むイメージ
// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	//0600⇒webサーバを立ち上げたものが読み書きできる
// 	return ioutil.WriteFile(filename, p.Body, 0600)
// }

// //pageの情報をデータベースから読み込むイメージ
// func loadPage(title string) (*Page, error) {
// 	filename := title + ".txt"
// 	body, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//返り値が二つ以上あるときにはカンマで区切る
// 	return &Page{Title: title, Body: body}, nil
// }

// //インターフェースのhttp.ResponseWriter(responseのストラクトが入っている)
// //http.Reuest(Requestのなかみのでーたなどがはいっている) を持つハンドラーをそれぞれ第１，２引数にしなくてはいけない
// //HandleFunc()の中でこの引数たちが呼ばれることになる
// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	// /view/test
// 	// Request
// 	title := r.URL.Path[len("/view/"):]
// 	p, _ := loadPage(title)
// 	//Fprintf()⇒第一引数のio.Writerに書き込みができる
// 	//第二引数にフォーマット、第三引数以降に値を入れる
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div", p.Title, p.Body)
// }

// func main() {
// 	//　/viewに行ったときにはviewHanderにいくように指示する
// 	// /以外のURLにハンドラーを登録している
// 	http.HandleFunc("/view/", viewHandler)
// 	//第一引数でport 第二引数でハンドラーを気宇jつする
// 	//:8080は後ろに何も書かなければlocalが自動的につく、nilは/で読んだときにdefaultのnoPageが返ってくる
// 	//サーバが走っているときに何かerrorが発生したらerrorを返す
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

//3 html template

// import (
// 	// "fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"text/template"
// )

// type Page struct {
// 	Title string
// 	Body []byte
// }

// //pageの情報をデータベースに書き込むイメージ
// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	//0600⇒webサーバを立ち上げたものが読み書きできる
// 	return ioutil.WriteFile(filename, p.Body, 0600)
// }

// func loadPage(title string) (*Page, error) {
// 	filename := title + ".txt"
// 	body, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }

// func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
// 	//templateを検索
// 	t , _ := template.ParseFiles(tmpl + ".html")
// 	//作成したtemplateを使ってwにpを書き込む
// 	t.Execute(w, p)
// }

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
// 		//関数終了
// 		return
// 	}
// 	renderTemplate(w, "view", p)
// }

// func editHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/edit/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		p = &Page{Title: title}
// 	}
// 	renderTemplate(w, "edit", p)
// }

// func saveHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/save/"):]
// 	//formのnameValueから取り出すことができる
// 	body := r.FormValue("body")
// 	p := &Page{Title:title, Body: []byte(body)}
// 	err := p.save()
// 	if err != nil {
// 		//StatusInternalServerError⇒500バンコードをer.Error()の内容とともにwに格納する
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		//関数終了
// 		return
// 	}
// 	http.Redirect(w, r, "/view/"+title, http.StatusFound)
// }

// func main() {
// 	http.HandleFunc("/view/", viewHandler)
// 	http.HandleFunc("/edit/", editHandler)
// 	http.HandleFunc("/save/", saveHandler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

//4 tmeplate cashing

import (
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	//ファイルが存在するときは上書き 存在しないときは新しくファイルを作る
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//titleに合うtxtファイルからbodyを取り出しPageStructとして返す
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//事前に複数のfileを読み込んでそのその結果を関数の中で利用する
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	//通常バージョン
	// t , _ := template.ParseFiles(tmpl + ".html")
	// t.Execute(w, p)

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		//titleが存在しないため新しく、bodyが空のPageStructを作る
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

//saveボタンの発火(editページに存在する)
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//makeHandlerで省略化
	// title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title:title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// urlを操作するためのひな型
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//urlのqueryを取り出す
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			//handleFunc関数の中でエラーを表示
			http.NotFound(w, r)
			return
		}
		//取り出したqueryを格納してhandlerを呼び出すことによってhandler内でpath空queryくぉとりだす操作を書かずに済む
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}