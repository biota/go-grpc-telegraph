package config

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Parses a settings file and returns a map of loaded key-value pairs.
func LoadSettings(path string) (map[string]string, error) {
	if len(path) == 0 {
		// No settings file.
		return nil, fmt.Errorf("invalid settings file")
	}

	slog.Debug("Using settings from", "path", path)

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		slog.Error("opening settings", "path", path, "error", err)
		return nil, err
	}

	defer file.Close()

	valueRegExp := regexp.MustCompile(`^([^#"']|'[^']*'|"[^"]*")*`)

	config := make(map[string]string)

	ocr := bufio.NewScanner(file)
	for ocr.Scan() {
		line := strings.TrimSpace(ocr.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			slog.Debug("skipping blank or comment line",
				"line", line)
			continue
		}

		pieces := strings.SplitN(line, "=", 2)
		if len(pieces) != 2 {
			slog.Error("invalid config", "path", path, "line",
				line)
			return nil, fmt.Errorf("invalid config in '%v': %v", path, line)
		}

		k := pieces[0]
		value := valueRegExp.Find([]byte(pieces[1]))
		v := strings.TrimSpace(string(value))

		if s, err := strconv.Unquote(v); err == nil {
			v = s
		}

		slog.Debug("processing settings", "line", line)
		slog.Debug("parsed line", "key", k, "value", v)

		config[k] = v
	}

	return config, ocr.Err()

} // End of function  LoadSettings.
