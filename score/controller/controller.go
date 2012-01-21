package controller

import (
	"appengine"
	"http"
	"model"
	"os"
	"strconv"
	"template"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/add", add)
	http.HandleFunc("/batch_insert", batchinsert)
}

var view = template.Must(template.New("view").ParseFile("view/front.html"))
func index(w http.ResponseWriter, r *http.Request) {
	/* アクセスされた時に呼び出される、コントローラ */

	c := appengine.NewContext(r)

	// データを取得
	var score []model.Score
	var err os.Error
	score, err = model.ScoreList(c)
	if err != nil {
		systemError(w, err)
		return
	}

	// ブラウザに表示
	if err := view.Execute(w, score); err != nil {
		systemError(w, err)
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	var newScore model.Score

	c := appengine.NewContext(r)

	// パラメータチェック（エラーは無視）
	no, _ := strconv.Atoi(r.FormValue("no"))
	name := template.HTMLEscapeString(r.FormValue("name"))
	japanese, _ := strconv.Atoi(r.FormValue("japanese"))
	math, _ := strconv.Atoi(r.FormValue("math"))
	english, _ := strconv.Atoi(r.FormValue("english"))

	// データを生成
	newScore.Set(no, name, japanese, math, english)

	// データを登録
	if err := newScore.AddScore(c); err != nil {
		systemError(w, err)
		return
	}

	// トップページにリダイレクト
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func systemError(w http.ResponseWriter, err os.Error) {
	http.Error(w, err.String(), http.StatusInternalServerError)
}

func batchinsert(w http.ResponseWriter, r *http.Request) {
	var newScore model.Score

	c := appengine.NewContext(r)

	no, _ := strconv.Atoi(r.FormValue("no"))
	name := template.HTMLEscapeString(r.FormValue("name"))
	japanese, _ := strconv.Atoi(r.FormValue("japanese"))
	math, _ := strconv.Atoi(r.FormValue("math"))
	english, _ := strconv.Atoi(r.FormValue("english"))

	// データを生成
	newScore.Set(no, name, japanese, math, english)

	// データを登録
	if err := newScore.AddScore(c); err != nil {
		systemError(w, err)
		return
	}
	
}
