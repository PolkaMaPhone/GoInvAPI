package item

import (
	"testing"

	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/dbconn"
	"github.com/stretchr/testify/assert"
)

func TestRepo_GetItemByID(t *testing.T) {
	config, err := dbconn.LoadConfigFile()
	if err != nil {
		t.Fatalf("Unable to load configuration: %v\n", err)
	}
	db := &dbconn.PgxDB{}
	_, err = dbconn.New(config, db)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}

	repo := NewRepository(db.Pool)
	item, err := repo.GetItemByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, item)
}

func TestRepo_GetAllItems(t *testing.T) {
	config, err := dbconn.LoadConfigFile()
	if err != nil {
		t.Fatalf("Unable to load configuration: %v\n", err)
	}
	db := &dbconn.PgxDB{}
	_, err = dbconn.New(config, db)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}

	repo := NewRepository(db.Pool)
	items, err := repo.GetAllItems()

	assert.NoError(t, err)
	assert.NotNil(t, items)
}
