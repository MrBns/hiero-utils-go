package mirrornode

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func checkAPIStatus(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return nil
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	var apiErr APIError
	if len(bodyBytes) > 0 && json.Unmarshal(bodyBytes, &apiErr) == nil {
		messages := make([]string, 0, len(apiErr.Status.Messages))
		for _, message := range apiErr.Status.Messages {
			if message.Message != "" {
				messages = append(messages, message.Message)
			}
		}
		if len(messages) > 0 {
			return fmt.Errorf("mirror node error (status=%d): %s", resp.StatusCode, strings.Join(messages, "; "))
		}
	}

	trimmedBody := strings.TrimSpace(string(bodyBytes))
	if trimmedBody == "" {
		return fmt.Errorf("mirror node error (status=%d)", resp.StatusCode)
	}

	return fmt.Errorf("mirror node error (status=%d): %s", resp.StatusCode, trimmedBody)
}

func baseURLForNetwork(network NetworkType) (string, error) {
	switch network {
	case MainnetNetwork:
		return mainnetBaseURL, nil
	case TestnetNetwork:
		return testnetBaseURL, nil
	case PreviewnetNetwork:
		return previewnetBaseURL, nil
	default:
		return "", fmt.Errorf("unsupported network type: %q", network)
	}
}
