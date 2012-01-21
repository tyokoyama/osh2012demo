package controller

import (
	"appengine"
	"fmt"
	"http"
	"model"
	"strconv"
	"template"
)

var detailTemplate = template.Must(template.New("detail").ParseFile("view/detail.html"))

type detail_view struct {
	Group model.Group
	Member []model.Member
}

func view(w http.ResponseWriter, r *http.Request) {
	
	id, _ := strconv.Atoi(r.FormValue("id"))

	c := appengine.NewContext(r)

	var view detail_view

	// グループ情報を取得
	if group, err := model.GetGroup(c, id); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	} else {
		view.Group = group
	}

	// メンバー情報を取得
	if memberlist, err := model.MemberList(c, id); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	} else {
		view.Member = memberlist
	}

	// 詳細画面を表示
	if err := detailTemplate.Execute(w, view); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}

func addmember(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := template.HTMLEscapeString(r.FormValue("name"))
	url := fmt.Sprintf("/view?id=%d", id)

	c := appengine.NewContext(r)

	// メンバーを追加
	var member model.Member
	member.Id = id
	member.Name = name
	if err := member.Add(c); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}

	// 詳細画面にリダイレクト
	http.Redirect(w, r, url, http.StatusFound)
}
