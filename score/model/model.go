package model

import (
	"appengine"
	"appengine/datastore"
	"os"
	"time"
)

type Score struct {
	No int
	Name string
	Japanese int `datastore:"kokugo,"`
	Math int `datastore:"sansuu,"`
	English int `datastore:"eigo,"`

	CreateDate datastore.Time

	Japanese_noindex int `datastore:",noindex"`
	Math_noindex int `datastore:",noindex"`
	English_noindex int `datastore:",noindex"`

	Average float32 `datastore:"-"`                // データストアに保存しない
}

func (score *Score)Load(c <-chan datastore.Property) os.Error {
	if err := datastore.LoadStruct(score, c); err != nil {
		return err
	}

	// 平均点を設定する。
	score.Average = float32(score.Japanese + score.Math + score.English) / 3

	return nil
}

func (score *Score)Save(c chan<- datastore.Property) os.Error {

	// インデックスがない値にコピー
	score.Japanese_noindex = score.Japanese
	score.Math_noindex = score.Math
	score.English_noindex = score.English
	
	datastore.SaveStruct(score, c)

	return nil
}

func (score *Score)Set(no int, name string, japanese, math, english int) {
	score.No = no
	score.Name = name
	score.Japanese = japanese
	score.Math = math
	score.English = english

	score.CreateDate = datastore.SecondsToTime(time.Seconds())

	// 格納されない値も計算
	score.Average = float32(japanese + math + english) / 3
}

/* 成績を登録する */
func (score Score)AddScore(c appengine.Context) os.Error {

	// 登録時はKeyは自動採番
	key := datastore.NewIncompleteKey(c, "Score", nil)
	if _, err := datastore.Put(c, key, &score); err != nil {
		return err
	}

	return nil
}

/* 成績リストを取得する */
func ScoreList(c appengine.Context) ([]Score, os.Error) {
	var scorelist []Score

//	q := datastore.NewQuery("Score").Order("kokugo")
//	q := datastore.NewQuery("Score").Order("-kokugo").Order("sansuu").Order("-eigo").Limit(30)
//	q := datastore.NewQuery("Score").Order("kokugo").Offset(0).Limit(1)
//	q := datastore.NewQuery("Score").Filter("sansuu =", 75)
//	q := datastore.NewQuery("Score").Filter("sansuu >=", 75).Filter("eigo =", 75)
//	q := datastore.NewQuery("Score").Filter("sansuu >=", 75).Filter("sansuu <=", 85)
//	q := datastore.NewQuery("Score").Order("Math_noindex")
//	q := datastore.NewQuery("Score").Filter("Japanese_noindex >=", 75)
//	q := datastore.NewQuery("Score").Filter("kokugo >=", 75).Order("sansuu")
//	q := datastore.NewQuery("Score").Filter("kokugo >=", 75).Order("sansuu").Order("kokugo")
//	q := datastore.NewQuery("Score").Filter("kokugo >=", 75).Order("-kokugo").Order("sansuu")
	q := datastore.NewQuery("Score").Filter("Name = ", "横").Order("Name")
	if count, err := q.Count(c); err != nil {
		return nil, err
	} else {
		scorelist = make([]Score, 0, count)
	}

	if _, err := q.GetAll(c, &scorelist); err != nil {
		return nil, err
	}

	return scorelist, nil
}

