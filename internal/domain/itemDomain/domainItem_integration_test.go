package itemDomain

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
	pool, err := dbconn.GetPoolInstance(config, db)
	_, err = dbconn.GetPoolInstance(config, db)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	db.Pool = pool
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
	pool, err := dbconn.GetPoolInstance(config, db)
	_, err = dbconn.GetPoolInstance(config, db)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	db.Pool = pool
	repo := NewRepository(db.Pool)
	items, err := repo.GetAllItems()

	assert.NoError(t, err)
	assert.NotNil(t, items)
}
func TestRepo_GetAllItemsWithCategory(t *testing.T) {
	config, err := dbconn.LoadConfigFile()
	if err != nil {
		t.Fatalf("Unable to load configuration: %v\n", err)
	}
	db := &dbconn.PgxDB{}
	pool, err := dbconn.GetPoolInstance(config, db)
	_, err = dbconn.GetPoolInstance(config, db)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	db.Pool = pool
	repo := NewRepository(db.Pool)
	items, err := repo.GetAllItemsWithCategories()

	assert.NoError(t, err)
	assert.NotNil(t, items)
}

func TestRepo_GetItemByIDWithCategory(t *testing.T) {
	config, err := dbconn.LoadConfigFile()
	if err != nil {
		t.Fatalf("Unable to load configuration: %v\n", err)
	}
	db := &dbconn.PgxDB{}
	pool, err := dbconn.GetPoolInstance(config, db)
	_, err = dbconn.GetPoolInstance(config, db)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	db.Pool = pool
	repo := NewRepository(db.Pool)
	item, err := repo.GetItemByIDWithCategory(1)

	assert.NoError(t, err)
	assert.NotNil(t, item)
}
