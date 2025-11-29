package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
)

func (h *handler) handleUpload(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB limit
		helper.BadRequest(h.logger, w, "failed to parse multipart form", err)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		helper.BadRequest(h.logger, w, "failed to retrieve file from form", err)
		return
	}
	defer file.Close()

	// Create uploads directory if it doesn't exist
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	// Generate unique filename
	ext := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, filename)

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, file); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	// Construct file URL
	// Assuming the server serves static files from /uploads
	fileURL := fmt.Sprintf("http://localhost:8080/uploads/%s", filename)

	helper.Created(h.logger, w, "file uploaded successfully", map[string]string{"url": fileURL})
}
