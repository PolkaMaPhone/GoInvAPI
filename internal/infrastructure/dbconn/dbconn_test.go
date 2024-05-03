package dbconn

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"path/filepath"
	"testing"
)

type MockDB struct {
	PgxDB
	mock.Mock
}

func (m *MockDB) Connect(connectionString string) error {
	args := m.Called(connectionString)
	return args.Error(0)
}

func TestLoadConfigFile(t *testing.T) {
	testCases := []struct {
		name             string
		ProjectRoot      string
		configJSON       string
		expectErr        bool
		configSampleJSON string
	}{
		{
			name:        "ValidConfig",
			ProjectRoot: "/test",
			configJSON: `{
				"DbUser": "username",
				"DbPassword": "password",
				"DbHost": "localhost",
				"DbPort": "5432",
				"DbName": "testdb",
				"DbSchema": "public"
			}`,
			configSampleJSON: `{
				"DbUser": "username",
				"DbPassword": "password",
				"DbHost": "localhost",
				"DbPort": "5432",
				"DbName": "testdb",
				"DbSchema": "public"
			}`,
			expectErr: false,
		},
		{
			name:        "InvalidConfig",
			ProjectRoot: "/test",
			configJSON: `{
				"DbUser": "username",
			}`,
			expectErr: true,
		},
		{
			name:        "EmptyRootDir",
			ProjectRoot: "",
			configJSON: `{
				"DbUser": "username",
				"DbPassword": "password",
				"DbHost": "localhost",
				"DbPort": "5432",
				"DbName": "testdb",
				"DbSchema": "public"
			}`,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.ProjectRoot != "" {
				// Create the directory within PROJECT_ROOT
				testDir := filepath.Join(os.Getenv("PROJECT_ROOT"), tc.ProjectRoot)
				err := os.MkdirAll(testDir, os.ModePerm)
				assert.NoError(t, err)

				// Create the config.json and config.json.sample files
				err = os.WriteFile(filepath.Join(testDir, "config.json"), []byte(tc.configJSON), 0644)
				assert.NoError(t, err)
				err = os.WriteFile(filepath.Join(testDir, "config.json.sample"), []byte(tc.configSampleJSON), 0644)
				assert.NoError(t, err)

				// Set the PROJECT_ROOT environment variable to the test directory
				err = os.Setenv("PROJECT_ROOT", testDir)
				assert.NoError(t, err)

				defer func() {
					err := os.Remove(filepath.Join(testDir, "config.json"))
					err = os.Remove(filepath.Join(testDir, "config.json.sample"))
					if err != nil {
						t.Fatalf("Failed to clean up: %v", err)
					}
				}()
			}
		})
	}
}

func TestGetPoolInstance(t *testing.T) {
	tests := []struct {
		name        string
		configFile  string
		expectError bool
		ProjectRoot string
	}{
		{
			name:        "connection fails",
			ProjectRoot: "/test",
			configFile: `{
				"DbUser": "postgres",
				"DbPassword": "wrongpassword",
				"DbHost": "localhost",
				"DbPort": "5432",
				"DbName": "testdb",
				"DbSchema": "public"
			}`,
			expectError: true,
		},
		{
			name: "successful connection",
			configFile: `{
				"DbUser": "postgres",
				"DbPassword": "password",
				"DbHost": "localhost",
				"DbPort": "5432",
				"DbName": "testdb",
				"DbSchema": "public"
			}`,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary directory within PROJECT_ROOT
			testDir := filepath.Join(os.Getenv("PROJECT_ROOT"))
			err := os.MkdirAll(testDir, os.ModePerm)
			assert.NoError(t, err)

			// Create the config.json and config.json.sample files
			err = os.WriteFile(filepath.Join(testDir, "config.json"), []byte(tt.configFile), 0644)
			assert.NoError(t, err)
			err = os.WriteFile(filepath.Join(testDir, "config.json.sample"), []byte(tt.configFile), 0644)
			assert.NoError(t, err)

			// Set the PROJECT_ROOT environment variable to the test directory
			err = os.Setenv("PROJECT_ROOT", testDir)
			assert.NoError(t, err)

			// Load the configuration from the file
			config, err := LoadConfigFile()
			assert.NoError(t, err)

			// Create a mock DB and set its expected behavior
			mockDB := new(MockDB)
			if tt.expectError {
				mockDB.On("Connect", mock.Anything).Return(fmt.Errorf("mock error"))
			} else {
				// Create a new pgxpool.Pool
				//goland:noinspection SpellCheckingInspection
				pool, err := pgxpool.New(context.Background(), "host=localhost user=yourusername password=yourpassword dbname=yourdbname port=5432")
				if err != nil {
					t.Fatalf("Failed to create pool: %v", err)
				}
				// Assign the pool to the Pool field of the MockDB object
				mockDB.Pool = pool
				mockDB.On("Connect", mock.Anything).Return(nil)
			}

			// Call the GetPoolInstance function with the loaded configuration and the mock DB
			db1, err := GetPoolInstance(config, mockDB)

			// Check if the function returns an error as expected
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, db1)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, db1)

				// Call GetPoolInstance again and check that it returns the same instance
				db2, err := GetPoolInstance(config, mockDB)
				assert.NoError(t, err)
				assert.NotNil(t, db2)
				assert.Equal(t, db1, db2)
			}

			// Assert that the Connect method was called with the correct arguments
			mockDB.AssertCalled(t, "Connect", mock.Anything)
			defer func() {
				err := os.Remove(filepath.Join(testDir, "config.json"))
				err = os.Remove(filepath.Join(testDir, "config.json.sample"))
				err = os.Remove(testDir)
				if err != nil {
					t.Fatalf("Failed to clean up: %v", err)
				}
			}()
		})
	}
}
