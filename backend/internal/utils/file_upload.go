package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// SaveFile menyimpan multipart file ke folder tujuan dan mengembalikan nama file baru.
// folder: contoh "uploads" (relative to working dir)
// originalFilename: nama file asli (dipakai untuk ekstensi)
// file: multipart.File yang diterima dari r.FormFile(...)
func SaveFile(folder string, originalFilename string, file multipart.File) (string, error) {
	// pastikan folder ada
	if err := os.MkdirAll(folder, 0755); err != nil {
		return "", err
	}

	ext := filepath.Ext(originalFilename)
	if ext == "" {
		ext = ".bin"
	}

	// membuat nama file unik: timestamp + nama file normalisasi
	timestamp := time.Now().Format("20060102150405")
	// normalisasi nama: hapus karakter non-alphanumeric, ganti dengan `-`
	re := regexp.MustCompile(`[^\w\d]+`)
	base := re.ReplaceAllString(originalFilename, "-")
	if len(base) > 20 {
		base = base[:20]
	}

	// 🔥 PERBAIKAN UTAMA: pastikan semua lowercase
	base = strings.ToLower(base)
	ext = strings.ToLower(ext)

	newName := timestamp + "-" + base + ext
	dstPath := filepath.Join(folder, newName)

	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// salin isi file
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return newName, nil
}