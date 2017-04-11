package driver

import (
	"bytes"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/gen2brain/acra-go/acra"
)

// LevelDB driver
type LevelDB struct {
	db *leveldb.DB
}

// NewDB return new LevelDB driver
func NewDB() *LevelDB {
	return &LevelDB{}
}

// Open opens database
func (b *LevelDB) Open(file string) (err error) {
	b.db, err = leveldb.OpenFile(file, nil)
	return
}

// Put puts report
func (b *LevelDB) Put(key string, report acra.Report) (err error) {
	buf := new(bytes.Buffer)
	encoder := acra.NewEncoder(buf)

	err = encoder.Encode(report)
	if err != nil {
		return
	}

	err = b.db.Put([]byte(key), buf.Bytes(), nil)

	return
}

// Get gets report
func (b *LevelDB) Get(key string) (report acra.Report, err error) {
	v, err := b.db.Get([]byte(key), nil)
	if err != nil {
		return
	}

	decoder := acra.NewDecoder(bytes.NewBuffer(v))

	err = decoder.Decode(&report)

	return
}

// GetAll gets all reports
func (b *LevelDB) GetAll() (reports []acra.Report, err error) {
	reports = make([]acra.Report, 0)

	c := b.db.NewIterator(nil, nil)
	for c.Next() {
		v := c.Value()

		report := acra.Report{}

		decoder := acra.NewDecoder(bytes.NewBuffer(v))

		err = decoder.Decode(&report)
		if err != nil {
			return
		}

		reports = append(reports, report)
	}

	c.Release()
	err = c.Error()

	return
}

// Delete deletes report
func (b *LevelDB) Delete(key string) error {
	return b.db.Delete([]byte(key), nil)
}

// Close closes database
func (b *LevelDB) Close() error {
	return b.db.Close()
}
