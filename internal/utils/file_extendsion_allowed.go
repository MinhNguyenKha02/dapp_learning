package utils

var (
	AllowedMimeTypes = map[string]bool{
		// Video formats
		"video/mp4":  true,
		"video/avi":  true,
		"video/mpeg": true,
		"video/webm": true,

		// Audio formats
		"audio/mpeg":  true, // mp3
		"audio/wav":   true,
		"audio/ogg":   true,
		"audio/x-m4a": true,

		// Document formats
		"application/pdf":    true,
		"application/msword": true, // doc
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true, // docx
		"application/vnd.ms-excel": true, // xls
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         true, // xlsx
		"application/vnd.ms-powerpoint":                                             true, // ppt
		"application/vnd.openxmlformats-officedocument.presentationml.presentation": true, // pptx
		"application/vnd.oasis.opendocument.text":                                   true, // odt
		"application/vnd.oasis.opendocument.spreadsheet":                            true, // ods
		"application/vnd.oasis.opendocument.presentation":                           true, // odp
	}

	AllowedExtensions = map[string]bool{
		// Video
		".mp4": true, ".avi": true, ".mpeg": true, ".webm": true,
		// Audio
		".mp3": true, ".wav": true, ".ogg": true, ".m4a": true,
		// Documents
		".pdf": true, ".doc": true, ".docx": true,
		".xls": true, ".xlsx": true,
		".ppt": true, ".pptx": true,
		".odt": true, ".ods": true, ".odp": true,
	}

	// Helper function to determine folder based on extension
	DetermineFolder = func(extension string) string {
		switch extension {
		case ".mp4", ".avi", ".mov", ".mpeg", ".webm":
			return "video"
		case ".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".odt", ".ods", ".odp":
			return "document"
		case ".mp3", ".wav", ".ogg", ".m4a":
			return "audio"
		default:
			return "file"
		}
	}
)
