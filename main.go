package main

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var alphabets = "abcdefghijklmnopqrstuvwxyz"

func main() {
	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	db, err := leveldb.OpenFile("./data", o)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// batch := new(leveldb.Batch)
	// for _, i := range alphabets {
	//         for _, j := range alphabets {
	//                 key := []byte(string(i) + "-" + string(j))
	//                 val := []byte(string(j) + "-" + string(i))
	//                 batch.Put(key, val)
	//         }
	// }
	// if err := db.Write(batch, nil); err != nil {
	//         log.Println(err)
	// }

	val, err := db.Get([]byte("ab"), nil)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(val))

	log.Println("iterate through everything")
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		val := iter.Value()
		log.Println(string(key), string(val))
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		log.Fatal(err)
	}

	log.Println("iterating through prefix a")
	iter = db.NewIterator(util.BytesPrefix([]byte("a")), nil)
	for iter.Next() {
		key := iter.Key()
		val := iter.Value()
		log.Println(string(key), string(val))
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		log.Fatal(err)
	}

	s, err := db.SizeOf([]util.Range{
		util.Range{Start: []byte("a"), Limit: []byte("z")},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println("got size:", s, s.Sum())

}
