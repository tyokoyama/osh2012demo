package controller

import (
	"appengine"
	"http"
	"model"
	"strconv"
	"template"
)

var groupAddTemplate = template.Must(template.New("front").ParseFile("view/addGroup.html"))

type view_group struct {
	URL string
	Group model.Group
}

func add(w http.ResponseWriter, r *http.Request) {
	var view view_group
	view.URL = "/add/add"

	// 勉強会情報の登録画面を表示する。
	if err := groupAddTemplate.Execute(w, view); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}	
}

func edit(w http.ResponseWriter, r *http.Request) {
	var view view_group
	view.URL = "/add/update"

	c := appengine.NewContext(r)

	id, _ := strconv.Atoi(r.FormValue("id"))

	// 勉強会情報を取得
	if group, err := model.GetGroup(c, id); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	} else {
		view.Group = group
	}

	// 勉強会情報の登録画面を表示する。
	if err := groupAddTemplate.Execute(w, view); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}
