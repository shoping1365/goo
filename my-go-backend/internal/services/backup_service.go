package services

import (
	"archive/zip"
	"io"
	"log"
	"my-go-backend/internal/models"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BackupService handles storing and retrieving backup copies of media files.
// It keeps backups under public/uploads/media/backup/<period>/... where <period> is YYYY-MM.
type BackupService struct {
	// ProjectRoot should point to repository root (parent of public directory).
	ProjectRoot string
}

// NewBackupService returns a new BackupService. If projectRoot is empty, it will attempt to use os.Getwd().
func NewBackupService(projectRoot string) *BackupService {
	if projectRoot == "" {
		if wd, err := os.Getwd(); err == nil {
			projectRoot = filepath.Dir(wd)
		}
	}
	return &BackupService{ProjectRoot: projectRoot}
}

// CurrentPeriod returns current period folder name (YYYY-MM).
func (s *BackupService) CurrentPeriod() string {
	return time.Now().Format("2006-01")
}

// DailyPeriod returns current day folder name (YYYY-MM-DD).
func (s *BackupService) DailyPeriod() string {
	return time.Now().Format("2006-01-02")
}

// backupRoot returns absolute backup root path.
func (s *BackupService) backupRoot() string {
	return filepath.Join(s.ProjectRoot, "public", "uploads", "media", "backup")
}

// CopyToBackup copies srcAbs file into backup/<YYYY-MM>/<relPath> (relPath should not start with /).
func (s *BackupService) CopyToBackup(srcAbs string, relPath string) (string, error) {
	relPath = filepath.Clean(relPath)
	if relPath[0] == os.PathSeparator {
		relPath = relPath[1:]
	}
	destDir := filepath.Join(s.backupRoot(), s.CurrentPeriod(), filepath.Dir(relPath))
	if err := os.MkdirAll(destDir, 0o750); err != nil {
		return "", err
	}
	destAbs := filepath.Join(destDir, filepath.Base(relPath))
	if err := CopyFile(srcAbs, destAbs); err != nil {
		return "", err
	}
	return destAbs, nil
}

// CopyToDailyBackup copies srcAbs file into backup/<YYYY-MM-DD>/<relPath> (relPath should not start with /).
func (s *BackupService) CopyToDailyBackup(srcAbs string, relPath string) (string, error) {
	relPath = filepath.Clean(relPath)
	if relPath != "." && relPath[0] == os.PathSeparator {
		relPath = relPath[1:]
	}
	// strip leading uploads or media/uploads if present
	relPath = filepath.ToSlash(relPath)
	if strings.HasPrefix(relPath, "media/uploads/") {
		relPath = strings.TrimPrefix(relPath, "media/uploads/")
	} else if strings.HasPrefix(relPath, "uploads/") {
		relPath = strings.TrimPrefix(relPath, "uploads/")
	}
	destDir := filepath.Join(s.backupRoot(), s.DailyPeriod(), filepath.Dir(relPath))
	if err := os.MkdirAll(destDir, 0o750); err != nil {
		return "", err
	}
	destAbs := filepath.Join(destDir, filepath.Base(relPath))
	if err := CopyFile(srcAbs, destAbs); err != nil {
		return "", err
	}
	return destAbs, nil
}

// FindLatestBackup searches backup folders in reverse chronological order and returns the first match.
func (s *BackupService) FindLatestBackup(relPath string) (string, error) {
	root := s.backupRoot()
	entries, err := os.ReadDir(root)
	if err != nil {
		return "", err
	}
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Name() > entries[j].Name()
	})
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		p := filepath.Join(root, e.Name(), relPath)
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}
	return "", os.ErrNotExist
}

// RemoveBackups deletes all copies of relPath (and its trimmed variants) in every period folder.
// It returns the number of deleted files.
func (s *BackupService) RemoveBackups(relPath string) (int, error) {
	relPath = filepath.ToSlash(strings.TrimPrefix(relPath, "/"))
	// build possible variants of path that may have been stored
	variants := []string{relPath}
	short := relPath
	if strings.HasPrefix(short, "media/uploads/") {
		short = strings.TrimPrefix(short, "media/uploads/")
	}
	if strings.HasPrefix(short, "uploads/") {
		short = strings.TrimPrefix(short, "uploads/")
	}
	if short != relPath {
		variants = append(variants, short)
	}

	root := s.backupRoot()
	entries, err := os.ReadDir(root)
	if err != nil {
		return 0, err
	}
	removed := 0
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		for _, v := range variants {
			p := filepath.Join(root, e.Name(), v)
			if _, err := os.Stat(p); err == nil {
				if err := os.Remove(p); err == nil {
					removed++
				}
			}
		}
	}
	return removed, nil
}

// CopyFile performs a buffered copy from src to dst (overwriting if exists).
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	return nil
}

// AddOriginal copies original bytes to folder of current month and inserts/updates BackupEntry.
func (s *BackupService) AddOriginal(srcAbs string, relPath string, db *gorm.DB) error {
	relPath = filepath.ToSlash(strings.TrimPrefix(relPath, "/"))
	dest, err := s.CopyToBackup(srcAbs, relPath)
	if err != nil {
		return err
	}
	// Insert or update BackupEntry
	if db != nil {
		be := models.BackupEntry{RelPath: relPath, Period: s.CurrentPeriod()}
		db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&be)
	}
	log.Printf("[backup] original saved to %s", dest)
	return nil
}

// ArchiveOldMonths compresses folders older than keepMonths into a zip and marks BackupEntry.archived=true.
func (s *BackupService) ArchiveOldMonths(keepMonths int, db *gorm.DB) error {
	root := s.backupRoot()
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}
	cutoff := time.Now().AddDate(0, -keepMonths, 0)
	for _, e := range entries {
		if !e.IsDir() || len(e.Name()) != 7 {
			continue
		}
		periodTime, _ := time.Parse("2006-01", e.Name())
		if periodTime.After(cutoff) {
			continue // keep recent months as folders
		}
		folderPath := filepath.Join(root, e.Name())
		zipPath := folderPath + ".zip"
		// Skip if already archived
		if _, err := os.Stat(zipPath); err == nil {
			continue
		}
		if err := zipFolder(folderPath, zipPath); err != nil {
			return err
		}
		// Mark entries archived
		if db != nil {
			db.Model(&models.BackupEntry{}).Where("period = ?", e.Name()).Update("archived", true)
		}
		// Remove original folder
		if err := os.RemoveAll(folderPath); err != nil {
			log.Printf("Failed to remove folder %s: %v", folderPath, err)
		}
	}
	return nil
}

// RemoveFromArchive removes relPath from any zip archives.
func (s *BackupService) RemoveFromArchive(relPath string) (int, error) {
	root := s.backupRoot()
	relPath = filepath.ToSlash(strings.TrimPrefix(relPath, "/"))
	count := 0
	filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(d.Name(), ".zip") {
			return nil
		}
		removed, _ := removeFileFromZip(path, relPath)
		count += removed
		return nil
	})
	return count, nil
}

// SaveOriginalAndBackup writes full data to absPath and stores a daily backup copy using the same relative path logic as CopyToDailyBackup.
// It is a convenience wrapper for image uploads so that handler code doesn't need to repeat backup logic.
func (s *BackupService) SaveOriginalAndBackup(absPath string, relPath string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(absPath), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(absPath, data, 0o644); err != nil {
		return err
	}
	if _, err := s.CopyToDailyBackup(absPath, relPath); err != nil {
		return err
	}
	return nil
}

// --- helper utilities ---

func zipFolder(srcDir, destZip string) error {
	zf, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zf.Close()
	zw := zip.NewWriter(zf)
	defer zw.Close()
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rel, _ := filepath.Rel(srcDir, path)
		rel = filepath.ToSlash(rel)
		fh, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		fh.Name = rel
		fh.Method = zip.Deflate
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return err
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		return err
	})
}

// removeFileFromZip recreates zip without the specific entry; returns 1 if removed.
func removeFileFromZip(zipPath, relPath string) (int, error) {
	zr, err := zip.OpenReader(zipPath)
	if err != nil {
		return 0, err
	}
	defer zr.Close()
	tmp := zipPath + ".tmp"
	zf, err := os.Create(tmp)
	if err != nil {
		return 0, err
	}
	zw := zip.NewWriter(zf)
	removed := 0
	for _, f := range zr.File {
		if f.Name == relPath {
			removed++
			continue
		}
		wf, err := zw.CreateHeader(&f.FileHeader)
		if err != nil {
			log.Printf("Failed to create header for file %s: %v", f.Name, err)
			continue
		}
		rc, err := f.Open()
		if err != nil {
			log.Printf("Failed to open file %s: %v", f.Name, err)
			continue
		}
		if _, err := io.Copy(wf, rc); err != nil {
			log.Printf("Failed to copy file %s: %v", f.Name, err)
		}
		rc.Close()
	}
	zw.Close()
	zf.Close()
	if removed > 0 {
		if err := os.Rename(tmp, zipPath); err != nil {
			log.Printf("Failed to rename temporary file %s to %s: %v", tmp, zipPath, err)
		}
	} else {
		if err := os.Remove(tmp); err != nil {
			log.Printf("Failed to remove temporary file %s: %v", tmp, err)
		}
	}
	return removed, nil
}
