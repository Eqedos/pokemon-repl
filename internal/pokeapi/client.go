package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/eqedos/repl/internal/cache"
)

const (
	// BaseURL is the base URL for the PokeAPI.
	BaseURL = "https://pokeapi.co/api/v2"

	// DefaultCacheTTL is the default time-to-live for cached responses.
	DefaultCacheTTL = 5 * time.Minute
)

// Client handles communication with the PokeAPI.
type Client struct {
	cache   *cache.Cache
	baseURL string
}

// NewClient creates a new PokeAPI client with caching enabled.
func NewClient() *Client {
	return &Client{
		cache:   cache.New(DefaultCacheTTL),
		baseURL: BaseURL,
	}
}

// GetLocationAreas fetches a paginated list of location areas from the given URL.
func (c *Client) GetLocationAreas(url string) (*LocationAreasResponse, error) {
	data, err := c.fetchWithCache(url)
	if err != nil {
		return nil, err
	}

	var response LocationAreasResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("failed to parse location areas: %w", err)
	}

	return &response, nil
}

// GetLocationArea fetches details for a specific location area by name.
func (c *Client) GetLocationArea(name string) (*LocationAreaResponse, error) {
	url := fmt.Sprintf("%s/location-area/%s/", c.baseURL, name)

	data, err := c.fetchWithCache(url)
	if err != nil {
		return nil, err
	}

	var response LocationAreaResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("failed to parse location area: %w", err)
	}

	return &response, nil
}

// GetFirstLocationAreasURL returns the URL for the first page of location areas.
func (c *Client) GetFirstLocationAreasURL() string {
	return fmt.Sprintf("%s/location-area/", c.baseURL)
}

// GetPokemon fetches details for a specific Pokemon by name.
func (c *Client) GetPokemon(name string) (*Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s/", c.baseURL, name)

	data, err := c.fetchWithCache(url)
	if err != nil {
		return nil, err
	}

	var response Pokemon
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("failed to parse pokemon: %w", err)
	}

	return &response, nil
}

// fetchWithCache retrieves data from the cache or fetches from the API.
// Returns whether the data was retrieved from cache.
func (c *Client) fetchWithCache(url string) ([]byte, error) {
	// Check cache first
	if data, ok := c.cache.Get(url); ok {
		fmt.Println("(using cached data)")
		return data, nil
	}

	// Fetch from API
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Store in cache
	c.cache.Add(url, data)

	return data, nil
}
