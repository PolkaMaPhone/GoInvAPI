package category

import (
	"testing"

	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/dbconn"
	"github.com/stretchr/testify/assert"
)

func TestRepo_GetCategoryByID(t *testing.T) {
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
	category, err := repo.GetCategoryByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, category)
}

func TestRepo_GetAllCategories(t *testing.T) {
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
	categories, err := repo.GetAllCategories()

	assert.NoError(t, err)
	assert.NotNil(t, categories)
}
