package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func LoadDotEnv(path string) error {
	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		key, val, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		val = strings.TrimSpace(val)
		key = strings.TrimSpace(key)
		if _, varFound := os.LookupEnv(key); !varFound {
			err := os.Setenv(key, val)
			if err != nil {
				return err
			}
		}
	}
	return sc.Err()
}
