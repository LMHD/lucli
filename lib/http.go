package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
	pb "gopkg.in/cheggaaa/pb.v2"
)

// GetJson does a HTTP request to a specified URL, then applies data to your interface
func GetJson(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// DownloadFile downloads a file from a URL to a specified directory
func DownloadFile(url, dir string, size int, replace bool) error {

	filename := dir + "/" + path.Base(url)

	// If we don't want to replace an existing file...
	if !replace {
		if _, err := os.Stat(filename); !os.IsNotExist(err) {
			log.Infof("%v already exists", filename)
			return nil
		}
	}

	log.Debugf("Downloading %v to %v", url, filename)

	// Create dir if it doesn't exist
	os.MkdirAll(dir, os.ModePerm)

	f, err := os.Create(filename)
	defer f.Close()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Could not download file: %s", err)
	}
	defer resp.Body.Close()

	// start new bar
	bar := pb.New(size)
	bar.Start()

	// create proxy reader
	body, err := ioutil.ReadAll(bar.NewProxyReader(resp.Body))
	if err != nil {
		return fmt.Errorf("Could not download file: %s", err)
	}

	// Write to file
	_, err = f.Write(body)
	if err != nil {
		return fmt.Errorf("Could not download file: %s", err)
	}

	bar.Finish()
	log.Debugf("Downloaded %v")

	return nil
}
