package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func WindowsRemoteUnzipping() {
	// Remote server details
	remoteServer := "T50523@INCHN0117DC"
	remoteTarFilePath := "D:/folder.tar"

	// Build the SSH command to extract the folder remotely using tar
	sshCmd := exec.Command("ssh", remoteServer, fmt.Sprintf("tar -xzvf %s", remoteTarFilePath))
	var stderr bytes.Buffer
	sshCmd.Stderr = &stderr

	// Execute the SSH command
	if err := sshCmd.Run(); err != nil {
		log.Fatalf("Failed to extract folder remotely: %v. Error output: %s", err, stderr.String())
	}

	fmt.Println("Folder extracted successfully on the remote server.")
}


func WindowsRemoteZipping() {
	remoteServer := "T50523@INCHN0117DC"
	remoteFolderPath := "D:/tomcat"
	remoteZipFilePath := "D:/GoLangTask/Demo.zip"

	// SSH command to create a tar.gz file remotely
	sshCmd := exec.Command("ssh", remoteServer, fmt.Sprintf("tar -czvf %s %s", remoteZipFilePath, remoteFolderPath))
	var stderr bytes.Buffer
	sshCmd.Stderr = &stderr

	// Execute the SSH command
	if err := sshCmd.Run(); err != nil {
		log.Fatalf("Failed to create tar.gz file remotely: %v. Error output: %s", err, stderr.String())
	}

	fmt.Println("Folder zipped successfully on the remote server.")
}




func LinuxRemoteUnzipping() {
	// Remote server details
	remoteServer := "T50523@INCHN0117DC"
	remoteTarFilePath := "/path/to/folder.tar"

	// Build the SSH command to extract the folder remotely using tar
	sshCmd := exec.Command("ssh", remoteServer, fmt.Sprintf("tar -xzvf %s -C %s", remoteTarFilePath, "/path/to/destination"))
	var stderr bytes.Buffer
	sshCmd.Stderr = &stderr

	// Execute the SSH command
	if err := sshCmd.Run(); err != nil {
		log.Fatalf("Failed to extract folder remotely: %v. Error output: %s", err, stderr.String())
	}

	fmt.Println("Folder extracted successfully on the remote server.")
}

func LinuxRemoteZipping() {
	remoteServer := "T50523@INCHN0117DC"
	remoteFolderPath := "/path/to/remote/tomcat"
	remoteZipFilePath := "/path/to/remote/Demo.tar.gz"

	// SSH command to create a tar.gz file remotely on a Linux system
	sshCmd := exec.Command("ssh", remoteServer, fmt.Sprintf("tar -czvf %s %s", remoteZipFilePath, remoteFolderPath))
	var stderr bytes.Buffer
	sshCmd.Stderr = &stderr

	// Execute the SSH command
	if err := sshCmd.Run(); err != nil {
		log.Fatalf("Failed to create tar.gz file remotely: %v. Error output: %s", err, stderr.String())
	}

	fmt.Println("Folder zipped successfully on the remote Linux server.")
}