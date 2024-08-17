package service

import (
	"context"
	"errors"
	"net/http"
	"os"
)

type NotionService struct {
	apiKey string
}

func NewNotionService() *NotionService {
	return &NotionService{
		apiKey: os.Getenv("NOTION_API_KEY"),
	}
}

func (s *NotionService) Connect(ctx context.Context, url string) (*http.Response, error) {
	if url == "" {
		return nil, errors.New("url is required")
	}

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set(("Authorization"), "Bearer "+s.apiKey)
	req.Header.Set("Notion-Version", "2021-05-13")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil
}