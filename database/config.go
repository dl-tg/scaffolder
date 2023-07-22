package database

import (
	"time"

	"github.com/boltdb/bolt"
)

var configBucket = []byte("config")
var db *bolt.DB

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(configBucket)
		return err
	})
}

func StoreConfigDir(name string, path string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(configBucket)
		err := b.Put([]byte(name), []byte(path))
		return err
	})
}

func GetConfigDir(name string) (string, error) {
	var path string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(configBucket)
		path = string(b.Get([]byte(name)))
		return nil
	})
	return path, err
}
