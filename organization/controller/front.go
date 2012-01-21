package controller

import (
	"appengine"
	"http"
	"model"
	"template"
)

func init() {
	http.HandleFunc("/", front)
	http.HandleFunc("/add", add)
	http.HandleFunc("/add/add", groupAdd)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/view", view)
	http.HandleFunc("/view/add", addmember)
}

var frontTemplate = template.Must(template.New("front").ParseFile("view/front.html"))

func front(w http.ResponseWriter, r *http.Request) {
	// 勉強会情報とメンバー情報を取得する。
	var grouplist []model.Group

	c := appengine.NewContext(r)

	if list, err := model.GroupList(c); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	} else {
		grouplist = list
	}

	// Viewを表示する。
	if err := frontTemplate.Execute(w, grouplist); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}

func groupAdd(w http.ResponseWriter, r *http.Request) {
	// 勉強会情報を登録する。
	var group model.Group

	c := appengine.NewContext(r)

	group.Name = r.FormValue("name")
	group.Place = r.FormValue("place")

	if err := group.Add(c); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}

	// トップページにリダイレクト
	http.Redirect(w, r, "/", http.StatusFound)
}
