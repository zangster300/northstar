package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failure", "error", err)
		os.Exit(1)
	}
}

func run() error {
	files := map[string]string{
		"https://raw.githubusercontent.com/starfederation/datastar/develop/bundles/datastar.js":     "internal/ui/static/datastar/datastar.js",
		"https://raw.githubusercontent.com/starfederation/datastar/develop/bundles/datastar.js.map": "internal/ui/static/datastar/datastar.js.map",
	}

	if err := os.RemoveAll("internal/ui/static/datastar"); err != nil {
		return fmt.Errorf("failed to remove static datastar directory: %w", err)
	}

	if err := os.MkdirAll("internal/ui/static/datastar", 0755); err != nil {
		return fmt.Errorf("failed to make static datastar directory: %w", err)
	}

	for url, filename := range files {
		base := filepath.Base(filename)
		slog.Info("Downloading...", "file", base, "url", url)
		if err := downloadFile(url, filename); err != nil {
			return fmt.Errorf("failed to download %s: %w", base, err)
		} else {
			slog.Info("Finished", "file", base)
		}
	}

	return nil
}

func downloadFile(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status downloading %s: %s", url, resp.Status)
	}

	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}

	return nil
}
