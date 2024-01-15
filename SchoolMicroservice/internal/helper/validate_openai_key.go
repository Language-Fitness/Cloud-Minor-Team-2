package helper

import (
	"fmt"
	"io"
	"net/http"
)

func ValidateOpenAiKey(apiKey string) error {
	url := "https://api.openai.com/v1/engines"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	} else {
		return fmt.Errorf("API key validation failded. Status code %d, Response %s\n", resp.StatusCode, string(body))
	}
}
