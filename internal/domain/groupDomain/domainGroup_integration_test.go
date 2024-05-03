package groupDomain

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/dbconn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_GetGroupByID(t *testing.T) {
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
	group, err := repo.GetGroupByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, group)
}

func TestRepo_GetAllGroups(t *testing.T) {
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
	groups, err := repo.GetAllGroups()

	assert.NoError(t, err)
	assert.NotNil(t, groups)
}
