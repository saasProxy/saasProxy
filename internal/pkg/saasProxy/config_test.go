package saasProxy

import (
	_ "net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfiguration(t *testing.T) {
	tests := []struct {
		name           string
		fileContent    string
		expectedConfig Configuration
		expectedError  bool
	}{
		{
			name: "Valid configuration file",
			fileContent: `
				Port = 8080
				Destination = "/destination"
				[[webhooks]]
					incoming_slug = "/webhook1"
					http_response_code = 200
					response_body = "OK"
					request_verb = "GET"
					[webhooks.headers]
						Content-Type = "application/json"
				[[webhooks]]
					incoming_slug = "/webhook2"
					http_response_code = 404
					response_body = "Not Found"
					request_verb = "POST"
					[webhooks.headers]
						Content-Type = "text/plain"
			`,
			expectedConfig: Configuration{
				Port:        8080,
				Destination: "/destination",
				Webhooks: []Webhook{
					{
						IncomingSlug:     "/webhook1",
						HttpResponseCode: 200,
						ResponseBody:     "OK",
						RequestVerb:      "GET",
						Headers: map[string]string{
							"Content-Type": "application/json",
						},
					},
					{
						IncomingSlug:     "/webhook2",
						HttpResponseCode: 404,
						ResponseBody:     "Not Found",
						RequestVerb:      "POST",
						Headers: map[string]string{
							"Content-Type": "text/plain",
						},
					},
				},
			},
			expectedError: false,
		},
		{
			name:           "Invalid configuration file",
			fileContent:    "invalid toml content",
			expectedConfig: Configuration{},
			expectedError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "config_test")
			assert.NoError(t, err)

			// Create a temporary file within the directory
			tempFile, err := os.Create(filepath.Join(tempDir, "config_test.toml"))
			assert.NoError(t, err)

			// Defer the removal of the temporary directory and its contents
			defer os.RemoveAll(tempDir)

			// Write the content to the file
			_, err = tempFile.WriteString(test.fileContent)
			assert.NoError(t, err)

			// Load configuration from the temporary file
			var config Configuration
			config, _ = LoadConfiguration("", config)

			// Check the results
			assert.Equal(t, test.expectedConfig, config)
			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
