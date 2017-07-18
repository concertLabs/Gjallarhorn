package importer

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Importer struct
type Importer struct {
	InputDir  string      // pdf input directory
	Processor chan string // process chan
	Wait      int         // how long should the loop sleep until the next input read starts
}

// NewImporter returns a new instance of the automatic import bot
// dir is the input directory where the importer searches for new pdf files
func NewImporter(dir string) Importer {
	return Importer{InputDir: dir, Processor: make(chan string), Wait: 10}
}

// Run scans for new files and sends them on a channel to the processor
func (i *Importer) Run() {
	for {
		files, err := readFiles(i.InputDir)
		if err != nil {
			continue
		}
		// process files and wait
		for _, f := range files {
			if !f.IsDir() {
				continue
			}

			i.Processor <- i.InputDir + f.Name()
		}

		time.Sleep(time.Second * time.Duration(i.Wait))
	}
}

func readFiles(dir string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("[Importer] readFiles: " + err.Error())
		return nil, err
	}
	return files, err
}
