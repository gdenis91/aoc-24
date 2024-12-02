package aoc

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	EnvAOCSession = "AOC_SESSION"
)

func userAgent() string {
	return "github.com/gdenis91/aoc-24 gsdenis91@gmail.com"
}

func aocInputURL(day int) string {
	return fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
}

func newRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent())
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", GetSessionCookie()))

	return req, nil
}

func GetInput(day int, part int) (string, error) {
	url := aocInputURL(day)
	resp, err := doRequest("GET", url, http.StatusOK, true)
	if err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}
	return resp, nil
}

func doRequest(
	method string,
	url string,
	statusCode int,
	cache bool,
) (string, error) {
	cacheKey := cacheKey(method + "_" + url)
	if cache {
		cached, err := readCachedResponse(cacheKey)
		if err != nil {
			return "", fmt.Errorf("read cached response: %w", err)
		}
		if cached != "" {
			return cached, nil
		}
	}

	req, err := newRequest(method, url)
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != statusCode {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("io read all: %w", err)
	}

	if cache {
		err = cacheResponse(cacheKey, string(body))
		if err != nil {
			return "", fmt.Errorf("cache response: %w", err)
		}
	}

	return string(body), nil
}

func cacheKey(key string) string {
	key = strings.ReplaceAll(key, "/", "_")
	return "aoc24_" + key
}

// cacheResponse caches the value as a tmp file named with the key and prefixed
// with "aoc24_"
func cacheResponse(key string, value string) error {
	tmpFile, err := os.Create(os.TempDir() + cacheKey(key))
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	_, err = tmpFile.WriteString(value)
	if err != nil {
		return err
	}

	return nil
}

func readCachedResponse(key string) (string, error) {
	tmpFile, err := os.Open(filepath.Join(os.TempDir(), cacheKey(key)))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", nil
		}
		return "", err
	}
	defer tmpFile.Close()

	cached, err := io.ReadAll(tmpFile)
	if err != nil {
		return "", fmt.Errorf("io read all: %w", err)
	}

	return string(cached), nil
}
