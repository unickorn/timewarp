package main

import (
	"flag"
	"github.com/barasher/go-exiftool"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

func main() {
	z, _ := zap.NewDevelopment()
	l := z.Sugar()
	var folder string
	flag.StringVar(&folder, "folder", "images", "the folder to update the creation times of files inside")
	var years int
	flag.IntVar(&years, "years", 0, "how many years do you want to move the files?")
	var months int
	flag.IntVar(&months, "months", 0, "how many months do you want to move the files?")
	var days int
	flag.IntVar(&days, "days", 0, "how many days do you want to move the files?")
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		l.Fatalf("failed to get working directory: %v", err)
	}
	files, err := os.ReadDir(folder)
	if err != nil {
		l.Fatalf("failed to read directory: %v", err)
	}
	e, err := exiftool.NewExiftool()
	if err != nil {
		l.Fatalf("failed to create exiftool: %v", err)
	}
	for _, f := range files {
		if filepath.Ext(f.Name()) != ".JPG" {
			continue
		}
		p := filepath.Join(wd, folder, f.Name())
		defer e.Close()
		originals := e.ExtractMetadata(p)

		withTimezone := "2006:01:02 15:04:05-07:00"
		withoutTimezone := "2006:01:02 15:04:05"
		withSubSeconds := "2006:01:02 15:04:05.00"
		thingsToModify := map[string]string{
			"DateTimeOriginal":       withoutTimezone,
			"ModifyDate":             withoutTimezone,
			"CreateDate":             withoutTimezone,
			"SubSecCreateDate":       withSubSeconds,
			"SubSecModifyDate":       withSubSeconds,
			"SubSecDateTimeOriginal": withSubSeconds,
			"FileModifyDate":         withTimezone,
			"FileInodeChangeDate":    withTimezone,
			"FileAccessDate":         withTimezone,
		}
		for thing, layout := range thingsToModify {
			t, err := originals[0].GetString(thing)
			if err != nil {
				panic(err)
			}
			tm, err := time.Parse(layout, t)
			if err != nil {
				panic(err)
			}
			originals[0].SetString(thing, tm.AddDate(years, months, days).Format(layout))
		}
		e.WriteMetadata(originals)
		l.Infow("updated", "file", f.Name())
	}
	l.Infow("done!", "files", len(files))
}
