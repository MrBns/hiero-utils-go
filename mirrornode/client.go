package mirrornode

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	network    NetworkType
	baseURL    string
	httpClient *http.Client
}

func NewClient(network NetworkType) (*Client, error) {
	baseURL, err := baseURLForNetwork(network)
	if err != nil {
		return nil, err
	}

	return &Client{
		network:    network,
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
	}, nil
}

func (c *Client) doGet(ctx context.Context, path string) (*http.Response, error) {
	reqURL, err := c.buildURL(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request mirror node: %w", err)
	}

	if err := checkAPIStatus(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) buildURL(path string) (string, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path, nil
	}

	if strings.HasPrefix(path, "/") {
		return c.baseURL + path, nil
	}

	parsedBase, err := url.Parse(c.baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base url: %w", err)
	}

	if parsedBase.Path == "" || strings.HasSuffix(parsedBase.Path, "/") {
		parsedBase.Path += path
	} else {
		parsedBase.Path += "/" + path
	}

	return parsedBase.String(), nil
}
