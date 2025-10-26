package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type FileReport struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	LastModified time.Time `json:"last_modified"`
}

type BackupReport struct {
	Files      []FileReport `json:"files"`
	TotalSize  int64        `json:"total_size"`
	TotalFiles int          `json:"total_files"`
}

func GenerateReport(folder string) (*BackupReport, error) {
	report := &BackupReport{}
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			report.Files = append(report.Files, FileReport{
				Name:         path,
				Size:         info.Size(),
				LastModified: info.ModTime(),
			})
			report.TotalSize += info.Size()
		}
		return nil
	})
	report.TotalFiles = len(report.Files)
	return report, err
}

func SaveReport(filename string, report *BackupReport) error {
	data, _ := json.MarshalIndent(report, "", "  ")
	return ioutil.WriteFile(filename, data, 0644)
}
