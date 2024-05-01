package dbconn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Connect(connectionString string) error {
	args := m.Called(connectionString)
	return args.Error(0)
}

func TestLoadConfigFile(t *testing.T) {
	testCases := []struct {
		name        string
		projectRoot string
		configJSON  string
		expectErr   bool
	}{
		{
			name:        "ValidConfig",
			projectRoot: "/test",
			configJSON: `{
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
			projectRoot: "/test",
			configJSON: `{
				"DbUser": "username",
			}`,
			expectErr: true,
		},
		{
			name:        "EmptyRootDir",
			projectRoot: "",
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
			if tc.projectRoot != "" {
				err := os.MkdirAll(tc.projectRoot, os.ModePerm)
				assert.NoError(t, err)
				err = os.WriteFile(tc.projectRoot+"/config.json", []byte(tc.configJSON), 0644)
				assert.NoError(t, err)
			}
			err := os.Setenv("PROJECT_ROOT", tc.projectRoot)
			assert.NoError(t, err)

			_, err = LoadConfigFile()
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tc.projectRoot != "" {
				err = os.RemoveAll(tc.projectRoot)
				assert.NoError(t, err)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		configFile  string
		expectError bool
	}{
		{
			name: "connection fails",
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
			// Create a temporary directory and a configuration file in it
			tmpDir, err := os.MkdirTemp("", "")
			assert.NoError(t, err)
			err = os.WriteFile(tmpDir+"/config.json", []byte(tt.configFile), 0644)
			assert.NoError(t, err)

			// Set the PROJECT_ROOT environment variable to the path of the temporary directory
			err = os.Setenv("PROJECT_ROOT", tmpDir)
			assert.NoError(t, err)

			// Load the configuration from the file
			config, err := LoadConfigFile()
			assert.NoError(t, err)

			// Create a mock DB and set its expected behavior
			mockDB := new(MockDB)
			if tt.expectError {
				mockDB.On("Connect", mock.Anything).Return(fmt.Errorf("mock error"))
			} else {
				mockDB.On("Connect", mock.Anything).Return(nil)
			}

			// Call the New function with the loaded configuration and the mock DB
			db, err := New(config, mockDB)

			// Check if the function returns an error as expected
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, db)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, db)
			}

			// Assert that the Connect method was called with the correct arguments
			mockDB.AssertCalled(t, "Connect", mock.Anything)

			// Remove the temporary directory
			err = os.RemoveAll(tmpDir)
			assert.NoError(t, err)
		})
	}
}
