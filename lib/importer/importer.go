package importer

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/quiteawful/Gjallarhorn/lib/config"
)

// Importer struct
type Importer struct {
	InputDir    string      // pdf input directory
	Processor   chan string // process chan
	Wait        int         // how long should the loop sleep until the next input read starts
	UseImporter bool
}

// New returns a new instance of the automatic import bot
// dir is the input directory where the importer searches for new pdf files
func New(cfg config.ImportConfig) Importer {
	return Importer{
		InputDir:    cfg.ScanDir,
		UseImporter: cfg.UseImporter,
		Processor:   make(chan string),
		Wait:        10,
	}
}

// Name returns the name, to satisfy the service interface
func (i Importer) Name() string {
	return "importer"
}

// Run scans for new files and sends them on a channel to the processor
func (i Importer) Run() {
	if !i.UseImporter {
		log.Println("Did not start importer. UseImporter is false")
		return
	}

	log.Printf("[Importer] Scan directory %s\n", i.InputDir)
	for {
		run(i.InputDir, i.Processor)
		time.Sleep(time.Second * time.Duration(i.Wait))
	}
}

func run(inputDir string, ch chan string) {
	// TODO: tests need to be improved
	files, err := readFiles(inputDir)
	if err != nil {
		log.Printf("[Importer] error while reading files from dir")
		return
	}

	// process files and wait
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		ch <- inputDir + f.Name()
	}
}

func readFiles(dir string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}
