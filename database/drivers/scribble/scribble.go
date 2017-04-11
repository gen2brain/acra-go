package driver

import (
	"bytes"

	"github.com/nanobox-io/golang-scribble"

	"github.com/gen2brain/acra-go/acra"
)

// ScribbleDB driver
type ScribbleDB struct {
	db         *scribble.Driver
	collection string
}

// NewDB return new LevelDB driver
func NewDB() *ScribbleDB {
	s := &ScribbleDB{}
	s.collection = "default"
	return s
}

// Open opens database
func (b *ScribbleDB) Open(dir string) (err error) {
	b.db, err = scribble.New(dir, nil)
	return
}

// Put puts report
func (b *ScribbleDB) Put(key string, report acra.Report) (err error) {
	err = b.db.Write(b.collection, key, report)
	return
}

// Get gets report
func (b *ScribbleDB) Get(key string) (report acra.Report, err error) {
	err = b.db.Read(b.collection, key, &report)
	return
}

// GetAll gets all reports
func (b *ScribbleDB) GetAll() (reports []acra.Report, err error) {
	records, err := b.db.ReadAll(b.collection)
	if err != nil {
		return
	}

	reports = make([]acra.Report, 0)

	for i := 0; i < len(records); i++ {
		b := []byte(records[i])
		report := acra.Report{}

		decoder := acra.NewDecoder(bytes.NewBuffer(b))

		err = decoder.Decode(&report)
		if err != nil {
			return
		}

		reports = append(reports, report)
	}

	return
}

// Delete deletes report
func (b *ScribbleDB) Delete(key string) error {
	return b.db.Delete(b.collection, key)
}

// Close closes database
func (b *ScribbleDB) Close() error {
	return nil
}
