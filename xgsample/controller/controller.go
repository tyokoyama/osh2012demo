package controller

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"http"
	"os"
)

func init() {
	http.HandleFunc("/", handler)
}

type Count struct {
	Counter int
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// 配列でカウント用の構造体とKeyを保持
	var counts []Count
	var keys []*datastore.Key

	counts = make([]Count, 6)                       // 配列の生成
	keys = make([]*datastore.Key, len(counts))

	for i := 0; i < len(keys); i++ {
		kind := fmt.Sprintf("Count%d", i)
		keys[i] = datastore.NewKey(c, kind, "mycounter", 0, nil)
	}

	// Cross Group Transactionを使う
	// XG: true Cross Group／ false Single Group
	options := new(datastore.TransactionOptions)
	options.XG = true
	err := datastore.RunInTransaction(c, func(c appengine.Context) os.Error {
		for pos, key := range keys {
			// とりあえず、Get
			if getErr := datastore.Get(c, key, &counts[pos]); getErr != nil && getErr != datastore.ErrNoSuchEntity {
				return getErr
			}

			// カウントアップしてPUT
			counts[pos].Counter++

			if _, putErr := datastore.Put(c, key, &counts[pos]); putErr != nil {
				return putErr
			}
		}

		return nil
	}, options)  // optionsがnilでもSingle Group

	// 結果を出力
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "text/html")
		for pos, count := range counts {
			fmt.Fprintf(w, "Count%d: %d<br>", pos, count.Counter)
		}
	}
}

