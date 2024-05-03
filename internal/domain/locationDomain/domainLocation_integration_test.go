package locationDomain

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/dbconn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_GetLocationByID(t *testing.T) {
	config, err := dbconn.LoadConfigFile()
	if err != nil {
		t.Fatalf("Unable to load configuration: %v\n", err)
	}
	db := &dbconn.PgxDB{}
	_, err = dbconn.GetPoolInstance(config, db)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}

	repo := NewRepository(db.Pool)
	location, err := repo.GetLocationByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, location)
}

func TestRepo_GetAllLocations(t *testing.T) {
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
	locations, err := repo.GetAllLocations()

	assert.NoError(t, err)
	assert.NotNil(t, locations)
}
