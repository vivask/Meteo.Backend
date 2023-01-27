package log

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"syscall"
	"time"
)

type RotateWriter struct {
	lock     sync.Mutex
	fp       *os.File
	filename string
}

// Make a new RotateWriter. Return nil if error occurs during setup.
func New(fp *os.File, fName string) *RotateWriter {
	return &RotateWriter{fp: fp, filename: fName}
}

// Perform the actual act of rotating and reopening file.
func (w *RotateWriter) Rotate() (err error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	// Close existing file if open
	if w.fp != nil {
		err = w.fp.Close()
		w.fp = nil
		if err != nil {
			return
		}
	}
	// Rename dest file if it already exists
	_, err = os.Stat(w.filename)
	if err == nil {
		err = os.Rename(w.filename, w.filename+"."+time.Now().Format(time.RFC3339))
		if err != nil {
			return
		}
	}

	return CreateLogger()
}

func (w *RotateWriter) Start(MaxFileSize, MaxFiles int) (err error) {

	ticker := time.NewTicker(30 * time.Second)

	go func() {
		for {
			<-ticker.C
			if isSizeExceeded(MaxFileSize) {
				w.Rotate()
				RemoveOutdated(logFileName, MaxFiles)
			}
		}
	}()

	return nil
}

func isSizeExceeded(MaxFileSize int) bool {
	fi, err := os.Stat(logFileName)
	if err != nil {
		Error(err)
	}
	return fi.Size() > int64(MaxFileSize)
}

func Rotate(fName string, MaxSize int64) error {
	const (
		chunksize int = 1024
	)

	stat, err := os.Stat(fName)
	if err != nil {
		return err
	}

	if stat.Size() > MaxSize {
		fi, err := os.OpenFile(fName, os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer fi.Close()

		fo, err := os.OpenFile(fName+"."+time.Now().Format(time.RFC3339), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer fo.Close()

		var count int
		r := bufio.NewReader(fi)
		w := bufio.NewWriter(fo)
		part := make([]byte, chunksize)
		for {
			if count, err = r.Read(part); err != nil {
				break
			}
			if _, err := w.Write(part[:count]); err != nil {
				return err
			}
		}
		err = fi.Truncate(0)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveOutdated(fName string, MaxFiles int) error {
	pattern := fmt.Sprintf("%s/%s.%s.*", filepath.Dir(fName), filepath.Base(fName), filepath.Ext(fName))
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	if len(matches) > MaxFiles {
		sort.Slice(matches, func(i, j int) bool {
			return ctime(matches[i]) < ctime(matches[j])
		})
		os.Remove(matches[0])
	}
	return nil
}

func ctime(fName string) int64 {
	fi, _ := os.Stat(fName)
	stat := fi.Sys().(*syscall.Stat_t)
	return time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec)).Unix()
}
