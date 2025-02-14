package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"dapp_learning/config"
	"dapp_learning/internal/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type DriveService struct {
	service *drive.Service
}

func NewGGDriveService(config *config.GoogleDriveConfig) (*DriveService, error) {
	ctx := context.Background()

	oauthConfig := &oauth2.Config{
		ClientID:     config.GoogleDriveClientID,
		ClientSecret: config.GoogleDriveClientSecret,
		RedirectURL:  config.GoogleDriveRedirectURI,
		Scopes: []string{
			drive.DriveFileScope,
		},
		Endpoint: google.Endpoint,
	}

	token := &oauth2.Token{
		RefreshToken: config.GoogleDriveRefreshToken,
	}

	client := oauthConfig.Client(ctx, token)

	service, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	return &DriveService{service: service}, nil
}

func (d *DriveService) UploadFile(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Determine folder based on file type
	folderName := utils.DetermineFolder(filepath.Ext(file.Filename))
	folderId, err := d.getOrCreateFolder("dapp_learning/media-courses/" + folderName)
	if err != nil {
		return "", err
	}

	driveFile := &drive.File{
		Name:     file.Filename,
		Parents:  []string{folderId},
		MimeType: file.Header.Get("Content-Type"),
	}

	// Create the file with public permissions
	_file, err := d.service.Files.Create(driveFile).Media(f).Do()
	if err != nil {
		return "", err
	}

	// Update permissions to make it publicly accessible
	permission := &drive.Permission{
		Role:               "reader",
		Type:               "anyone",
		AllowFileDiscovery: false,
	}
	_, err = d.service.Permissions.Create(_file.Id, permission).Do()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://drive.google.com/file/d/%s/view?usp=sharing", _file.Id), nil
}

func (d *DriveService) getOrCreateFolder(folderPath string) (string, error) {
	// Split path into parts
	parts := strings.Split(folderPath, "/")
	var currentFolderId string

	// Create each folder in path if it doesn't exist
	for _, part := range parts {
		query := fmt.Sprintf("name='%s' and mimeType='application/vnd.google-apps.folder'", part)
		if currentFolderId != "" {
			query += fmt.Sprintf(" and '%s' in parents", currentFolderId)
		}

		folder, err := d.service.Files.List().Q(query).Do()
		if err != nil {
			return "", err
		}

		// If folder exists, use it
		if len(folder.Files) > 0 {
			currentFolderId = folder.Files[0].Id
			continue
		}

		// Create folder if it doesn't exist
		folderMetadata := &drive.File{
			Name:     part,
			MimeType: "application/vnd.google-apps.folder",
		}
		if currentFolderId != "" {
			folderMetadata.Parents = []string{currentFolderId}
		}

		createdFolder, err := d.service.Files.Create(folderMetadata).Do()
		if err != nil {
			return "", err
		}
		currentFolderId = createdFolder.Id
	}

	return currentFolderId, nil
}
