package main

import (
	"flag"
	"github.com/djherbis/times"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path"
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
	for _, f := range files {
		i, err := f.Info()
		if err != nil {
			l.Fatalf("failed to get file info: %v", err)
		}
		t := times.Get(i)
		d := t.BirthTime().AddDate(years, months, days)
		_, err = exec.Command("touch", "-t", d.Format("200601021504.05"), path.Join(wd, folder, f.Name())).Output()
		if err != nil {
			l.Fatal(err)
		}
		_, err = exec.Command("SetFile", "-d", d.Format("01/02/2006 15:04:05"), path.Join(wd, folder, f.Name())).Output()
		if err != nil {
			l.Fatal(err)
		}
		l.Infow("timewarp done!",
			"file", i.Name(),
			"time", t.BirthTime().Format(time.RFC1123))
	}
}
