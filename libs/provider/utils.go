package provider

import (
	"errors"
	"os"
)

// TODO: Remove this, but research exactly what is needed so we don't break anything
func Host() (string, error) {
	host := os.Getenv("WARP_HOST")
	if host != "" {
		return host, nil
	}
	host, err := os.Hostname()
	if err == nil {
		return host, nil
	}
	return "", errors.New("WARP_HOST not set")
}

// TODO: Remove this, but research exactly what is needed so we don't break anything
func RequireHost() string {
	host, err := Host()
	if err != nil {
		panic(err)
	}
	return host
}

// TODO: Remove this, but research exactly what is needed so we don't break anything
func RequireVersion() string {
	if version := os.Getenv("WARP_VERSION"); version != "" {
		return version
	}

	return LocalVersion
}

// TODO: Remove this, but research exactly what is needed so we don't break anything
func RequireConfigVersion() string {
	if version := os.Getenv("WARP_CONFIG_VERSION"); version != "" {
		return version
	}
	return LocalVersion
}
