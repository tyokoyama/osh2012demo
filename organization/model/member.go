package model

import (
	"appengine"
	"appengine/datastore"
	"os"
	"time"
)

type Member struct {
	Id int
	Name string
	CreateDate datastore.Time
}

func (member *Member) Add(c appengine.Context) os.Error {
	parent := datastore.NewKey(c, "Group", "", int64(member.Id), nil)
	key := datastore.NewIncompleteKey(c, "Member", parent)

	member.CreateDate = datastore.SecondsToTime(time.Seconds())

	_, err := datastore.Put(c, key, member)

	return err
}

func MemberList(c appengine.Context, id int) ([]Member, os.Error) {
	var member []Member
	key := datastore.NewKey(c, "Group", "", int64(id), nil)

	q := datastore.NewQuery("Member").Ancestor(key)
	if count, err := q.Count(c); err != nil {
		return nil, err
	} else {
		member = make([]Member, 0, count)
	}

	if _, err := q.GetAll(c, &member); err != nil {
		return nil, err
	}

	return member, nil
}
