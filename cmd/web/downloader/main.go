package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"northstar/web/resources"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failure", "error", err)
		os.Exit(1)
	}
}

func run() error {
	files := map[string]string{
		"https://raw.githubusercontent.com/starfederation/datastar/develop/bundles/datastar.js":     resources.StaticDirectoryPath + "/datastar/datastar.js",
		"https://raw.githubusercontent.com/starfederation/datastar/develop/bundles/datastar.js.map": resources.StaticDirectoryPath + "/datastar/datastar.js.map",
		"https://github.com/saadeghi/daisyui/releases/latest/download/daisyui.js":                   resources.StylesDirectoryPath + "/daisyui/daisyui.js",
		"https://github.com/saadeghi/daisyui/releases/latest/download/daisyui-theme.js":             resources.StylesDirectoryPath + "/daisyui/daisyui-theme.js",
	}

	directories := []string{
		resources.StaticDirectoryPath + "/datastar",
		resources.StylesDirectoryPath + "/daisyui",
	}

	if err := removeDirectories(directories); err != nil {
		return err
	}

	if err := createDirectories(directories); err != nil {
		return err
	}

	if err := download(files); err != nil {
		return err
	}

	return nil
}

func removeDirectories(dirs []string) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(dirs))

	for _, path := range dirs {
		wg.Go(func() {
			if err := os.RemoveAll(path); err != nil {
				errCh <- fmt.Errorf("failed to remove static directory [%s]: %w", path, err)
			}
		})
	}

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func createDirectories(dirs []string) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(dirs))

	for _, path := range dirs {
		wg.Go(func() {
			if err := os.MkdirAll(path, 0755); err != nil {
				errCh <- fmt.Errorf("failed to create static directory [%s]: %w", path, err)
			}
		})
	}

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func download(files map[string]string) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(files))

	for url, filename := range files {
		wg.Go(func() {
			base := filepath.Base(filename)
			slog.Info("downloading...", "file", base, "url", url)
			if err := downloadFile(url, filename); err != nil {
				errCh <- fmt.Errorf("failed to download [%s]: %w", base, err)
			} else {
				slog.Info("finished", "file", base)
			}
		})
	}

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func downloadFile(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file [%s]: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status was not OK downloading file [%s]: %s", url, resp.Status)
	}

	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file [%s]: %w", filename, err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("failed to write file [%s]: %w", filename, err)
	}

	return nil
}
