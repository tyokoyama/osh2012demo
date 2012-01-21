package model

import (
	"appengine"
	"appengine/datastore"
	"os"
	"time"
)

type Counter struct {
	GroupCount int `datastore:,noindex`
}

type Group struct {
	Id int `datastore:,noindex`
	Name string
	Place string
	CreateDate datastore.Time
}

// 勉強会の追加
func (group *Group)Add(c appengine.Context) os.Error {

	count := new(Counter)
	countKey := datastore.NewKey(c, "Counter", "mycounter", 0, nil)
	
	countErr := datastore.RunInTransaction(c, func(c appengine.Context) os.Error {
		err := datastore.Get(c, countKey, count)
		if err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}

		count.GroupCount++

		_, err = datastore.Put(c, countKey, count)
		
		return err
	}, nil)

	if countErr != nil {
		return countErr
	}

	group.Id = count.GroupCount
	group.CreateDate = datastore.SecondsToTime(time.Seconds())

	key := datastore.NewKey(c, "Group", "", int64(group.Id), nil)

	_, err := datastore.Put(c, key, group)

	return err
}

// 勉強会の更新
func (group *Group)Put(c appengine.Context) os.Error {

	group.CreateDate = datastore.SecondsToTime(time.Seconds())

	key := datastore.NewKey(c, "Group", "", int64(group.Id), nil)

	_, err := datastore.Put(c, key, group)

	return err
}

// 勉強会リストを取得
func GroupList(c appengine.Context) ([]Group, os.Error) {
	var grouplist []Group

	q := datastore.NewQuery("Group")
	if count, err := q.Count(c); err != nil {
		return nil, err
	} else {
		grouplist = make([]Group, 0, count)
	}

	if _, err := q.GetAll(c, &grouplist); err != nil {
		return nil, err
	}

	return grouplist, nil

}

func GetGroup(c appengine.Context, id int) (Group, os.Error) {
	var group Group

	key := datastore.NewKey(c, "Group", "", int64(id), nil)

	err := datastore.Get(c, key, &group)

	return group, err
}