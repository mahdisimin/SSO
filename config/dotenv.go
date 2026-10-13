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
		if strings.HasPrefix(line, "#") {
			continue
		}
		key, val, _ := strings.Cut(line, "=")
		val = strings.TrimSpace(val)
		key = strings.TrimSpace(key)
		os.Setenv(key, val)
		// TODO : Check if env var has been set before (for example at terminal) , do NOT set it here
	}
	return sc.Err()
}
