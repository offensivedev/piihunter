package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
	"time"
	"strconv"
	"path/filepath"
	"encoding/json"
	"bufio"
	"io/ioutil"
	"regexp"
)

type file struct {
	filePath 		string
	permission 	string
}

type regexRule struct {
	name 				string
	displayName string
	regex 			*regexp.Regexp
}

type scanResult struct {
	ScanName 		string	`json:"ScanName"`
	FilePath 		string	`json:"FilePath"`
	MatchedData string	`json:"MatchedData"`
}

func scanWeakFilePermissions(fileList []file) []scanResult {
	var data []scanResult
	for x := range fileList {
		if (fileList[x].permission) == "777" {
			item := scanResult {ScanName: "Weak File Permission (777)", FilePath: fileList[x].filePath, MatchedData: ""}
			data = append(data, item)
		}
	}
	return data
}

/* 
	cloneGitRepository clones the repository locally on the base path specified by variable along with unix timestamp as random directory. 
	Expects a URL in gitURL parameter and returns the path of cloned repository if it was cloned. Returns nil otherwise. 
*/
func cloneGitRepository(gitURL string) string {
	fmt.Println("Fetching git repository " + gitURL)
	baseDirectory := "/tmp/"
	directory := "piihunter-" + strconv.FormatInt(time.Now().Unix(), 10)
	path := baseDirectory + directory
	_, err := git.PlainClone(path, false, &git.CloneOptions {
		URL: gitURL,
		Progress: os.Stdout,
	})
	
	if err != nil {
		fmt.Println(err)
		return "" 
	}
	return path
}

/* 
	getFiles is used to get list of files and their respective permissions from a base directory. 
	Returns a slice of file struct with their path and permissions. 
*/
func getFiles(baseDirectory string) []file {
	var fileList []file
	_, folderErr := os.Stat(baseDirectory)
	if os.IsNotExist(folderErr) {
		fmt.Println("Directory does not exist.")
		return fileList
	}
	err := filepath.Walk(baseDirectory, 
		func (path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// get file permission and convert it to octal value
			perm := info.Mode().Perm()
			octalPerm := strconv.FormatUint(uint64(perm), 8)

			item := file{filePath: path, permission: octalPerm}
			fileList = append(fileList, item)
			return nil
		})
	if err != nil {
		return fileList
	}
	return fileList
}

/* 
	checkRegexMatchInFile is used to match regex defined in const file against list of files. 
	Returns a slice of scanResult struct with files matched against regex. 
*/
func checkRegexMatchInFile(files []file) []scanResult {
	var data []scanResult
	for i := range files {
		fileToCheck, err := os.Open(files[i].filePath)
		if err != nil { 
			fmt.Println(err)
		}
		defer fileToCheck.Close()

		// check regex match line by line in fileToCheck file
		scanner := bufio.NewScanner(fileToCheck)
		for scanner.Scan() {
			for r := range regexRules {
				if regexRules[r].regex.MatchString(scanner.Text()) {
					item := scanResult {ScanName: regexRules[r].displayName,FilePath: files[i].filePath, MatchedData: scanner.Text()}
					data = append(data, item)
				}
			}
		}
	}
	return data
}

/* 
	publishScanResults is used output all the matches found. 
*/
func publishScanResults(data []scanResult) {
	for i := range data {
		if data[i].MatchedData != "" {
			fmt.Printf("%s identified in file %s\nMatched: %s\n\n", data[i].ScanName, data[i].FilePath, data[i].MatchedData)
		}	else {
			fmt.Printf("%s detected in file %s\n", data[i].ScanName, data[i].FilePath)
		}
	}
}

/* 
	writeOutputToFile is used write all results to supplied file in json format. 
*/
func writeOutputToFile(data []scanResult, fileName string) {
	output, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(fileName, output, 0644)
}