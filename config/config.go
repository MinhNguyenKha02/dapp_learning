package config

import (
	"os"

	"github.com/joho/godotenv"
)

type GoogleDriveConfig struct {
	GoogleDriveClientID     string
	GoogleDriveClientSecret string
	GoogleDriveRefreshToken string
	GoogleDriveRedirectURI  string
}

func LoadGGDriveConfig() (*GoogleDriveConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &GoogleDriveConfig{
		GoogleDriveClientID:     os.Getenv("GOOGLE_DRIVE_CLIENT_ID"),
		GoogleDriveClientSecret: os.Getenv("GOOGLE_DRIVE_CLIENT_SECRET"),
		GoogleDriveRefreshToken: os.Getenv("GOOGLE_DRIVE_REFRESH_TOKEN"),
		GoogleDriveRedirectURI:  os.Getenv("GOOGLE_DRIVE_REDIRECT_URI"),
	}, nil
}
