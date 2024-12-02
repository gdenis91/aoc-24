package aoc

import (
	"fmt"
	"io"
	"net/http"
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

func GetInput(day int) (string, error) {
	url := aocInputURL(day)
	req, err := newRequest("GET", url)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return readBody(resp.Body)
}

func readBody(r io.Reader) (string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
