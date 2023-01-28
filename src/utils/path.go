package utils

import (
	"net/http"
	"path/filepath"
	"strings"
)

func GetShowPathId(req *http.Request) string {
	sub := strings.TrimPrefix(req.URL.Path, "/show")
	_, id := filepath.Split(sub)
	if id != "" {
		return id
	}
	return ""
}
