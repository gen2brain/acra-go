package driver

import (
	"bytes"
	"errors"

	"github.com/boltdb/bolt"

	"github.com/gen2brain/acra-go/acra"
)

var (
	// ErrNoKey is returned when key doesn't exist
	ErrNoKey = errors.New("Key doesn't exist")
	// ErrNoBucket is returned when bucket doesn't exist
	ErrNoBucket = errors.New("Bucket doesn't exist")
)

// BoltDB driver
type BoltDB struct {
	db     *bolt.DB
	bucket []byte
}

// NewDB return new BoltDB driver
func NewDB() *BoltDB {
	b := &BoltDB{}
	b.bucket = []byte("default")
	return b
}

// Open opens database
func (b *BoltDB) Open(file string) (err error) {
	b.db, err = bolt.Open(file, 0600, nil)
	if err != nil {
		return
	}

	b.db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(b.bucket)
		return nil
	})

	return
}

// Put puts report
func (b *BoltDB) Put(key string, report acra.Report) error {
	var e error

	b.db.Update(func(tx *bolt.Tx) error {
		bk, err := tx.CreateBucketIfNotExists(b.bucket)
		if err != nil {
			e = err
			return err
		}

		buf := new(bytes.Buffer)
		encoder := acra.NewEncoder(buf)

		err = encoder.Encode(report)
		if err != nil {
			e = err
			return err
		}

		err = bk.Put([]byte(key), buf.Bytes())
		if err != nil {
			e = err
			return err
		}

		return nil
	})

	return e
}

// Get gets report
func (b *BoltDB) Get(key string) (acra.Report, error) {
	var e error

	var report acra.Report

	b.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket(b.bucket)
		if bk == nil {
			e = ErrNoBucket
			return e
		}

		v := bk.Get([]byte(key))
		if v == nil {
			e = ErrNoKey
			return e
		}

		decoder := acra.NewDecoder(bytes.NewBuffer(v))

		err := decoder.Decode(&report)
		if err != nil {
			e = err
			return err
		}

		return nil
	})

	return report, e
}

// GetAll gets all reports
func (b *BoltDB) GetAll() ([]acra.Report, error) {
	var e error

	reports := make([]acra.Report, 0)

	b.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket(b.bucket)
		if bk == nil {
			e = errors.New("Bucket doesn't exist")
			return e
		}

		c := bk.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			report := acra.Report{}

			decoder := acra.NewDecoder(bytes.NewBuffer(v))

			err := decoder.Decode(&report)
			if err != nil {
				e = err
				return err
			}

			reports = append(reports, report)
		}

		return nil
	})

	return reports, e
}

// Delete deletes report
func (b *BoltDB) Delete(key string) error {
	var e error

	b.db.Update(func(tx *bolt.Tx) error {
		bk, err := tx.CreateBucketIfNotExists(b.bucket)
		if err != nil {
			e = err
			return err
		}

		err = bk.Delete([]byte(key))
		if err != nil {
			e = err
		}

		return err
	})

	return e
}

// Close closes database
func (b *BoltDB) Close() error {
	return b.db.Close()
}
