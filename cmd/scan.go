/*
Copyright © 2021 Amit Kumar https://github.com/offensivedev

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan is used to initiate a scan on directory or git repository",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		var dir, _ = cmd.Flags().GetString("dir")
		var gitURL, _ = cmd.Flags().GetString("git")
		var outputFile, _ = cmd.Flags().GetString("out")
		var gitPath string
		var gitScanData []scanResult
		var dirScanData []scanResult
		var combinedOutput []scanResult
		if ( dir == "" && gitURL == "" ) {
			fmt.Println("Please supply either directory or github repository using --dir & --git flags")
		}

		if ( gitURL != "" ) {
			gitPath = cloneGitRepository(gitURL)
			if ( gitPath == "" ) {
				fmt.Println("Unable to fetch the git repository")
				os.Exit(1)
			}
			files := getFiles(gitPath)
			gitScanData = checkRegexMatchInFile(files)
			dirScanData = scanWeakFilePermissions(files)
			combinedOutput = append(combinedOutput, gitScanData...)
			combinedOutput = append(combinedOutput, dirScanData...)
		}

		if ( dir != "" ) {
			files := getFiles(dir)
			gitScanData = checkRegexMatchInFile(files)
			dirScanData = scanWeakFilePermissions(files)
			combinedOutput = append(combinedOutput, gitScanData...)
			combinedOutput = append(combinedOutput, dirScanData...)
		}
		publishScanResults(combinedOutput)
		if outputFile != "" {
			writeOutputToFile(combinedOutput, outputFile)
		}
		// clean up downloaded git repo after finishing all work. 
		defer os.RemoveAll(gitPath)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	scanCmd.PersistentFlags().String("dir", "", "Directory to scan for PII & Weak Permissions in")
	scanCmd.PersistentFlags().String("git", "", "URL for git repository (https://github.com/[USERNAME]/[REPOSITORY])")
	scanCmd.PersistentFlags().String("out", "", "Write output to file passed as argument. Currently only supports JSON.")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}