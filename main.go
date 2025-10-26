package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	src := flag.String("src", "", "Source directory to backup")
	dst := flag.String("dst", "", "Destination directory for backup")
	reportOnly := flag.Bool("report", false, "Generate JSON report only")
	verbose := flag.Bool("verbose", false, "Enable verbose output")
	flag.Parse()

	if *src == "" || *dst == "" {
		fmt.Println("Usage: back-it-up -src <source> -dst <destination> [-report] [-verbose]")
		os.Exit(1)
	}

	srcPath := expandPath(*src)
	dstPath := expandPath(*dst)

	logFile, err := os.OpenFile("backup.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {

		}
	}(logFile)
	log.SetOutput(logFile)

	if *reportOnly {
		report, err := GenerateReport(dstPath)
		if err != nil {
			log.Println("Error generating report:", err)
			fmt.Println("Error generating report:", err)
			return
		}
		err = SaveReport("backupReport.json", report)
		if err != nil {
			return
		}
		fmt.Println("Report generated:", len(report.Files), "files")
		return
	}

	err = BackupFolder(srcPath, dstPath, *verbose)
	if err != nil {
		log.Println("Backup failed:", err)
		fmt.Println("Backup failed:", err)
		return
	}

	report, err := GenerateReport(dstPath)
	if err != nil {
		log.Println("Error generating report:", err)
		return
	}
	err = SaveReport("backupReport.json", report)
	if err != nil {
		return
	}
	fmt.Println("Backup completed successfully. Report saved as backup-report.json")
}
