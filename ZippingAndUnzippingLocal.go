package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)


func ZippingFolderWindowsLocal() {
	// Folder to be archived
	sourceFolder := "D:/tomcat"

	// Tar file to be created
	tarFilePath := "D:/GoLangTask/Demo.tar.gz"

	// Create a new tar.gz file
	tarFile, err := os.Create(tarFilePath)
	if err != nil {
		log.Fatalf("Failed to create tar file: %v", err)
	}
	defer tarFile.Close()

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(tarFile)
	defer gzipWriter.Close()

	// Create a new tar writer
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// Walk through the source folder and add files to the tar archive
	err = filepath.Walk(sourceFolder, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the relative path
		relPath, err := filepath.Rel(sourceFolder, filePath)
		if err != nil {
			return err
		}

		// Create a tar header
		header, err := tar.FileInfoHeader(fileInfo, relPath)
		if err != nil {
			return err
		}

		// Write the header to the tar archive
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// If the file is not a directory, write its contents to the tar archive
		if !fileInfo.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			// Copy file contents to the tar archive
			if _, err := io.Copy(tarWriter, file); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to add files to tar archive: %v", err)
	}

	fmt.Printf("Folder '%s' zipped successfully to '%s'.\n", sourceFolder, tarFilePath)
}

func UnzippingFolderWindowsLocal() {
	// Tar.gz file to be extracted
	tarGzFilePath := "D:/GoLangTask/Demo.tar.gz"

	// Destination folder for extraction
	extractFolder := "D:/GoLangTask"

	// Open the tar.gz file for reading
	tarGzFile, err := os.Open(tarGzFilePath)
	if err != nil {
		log.Fatalf("Failed to open tar.gz file: %v", err)
	}
	defer tarGzFile.Close()

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(tarGzFile)
	if err != nil {
		log.Fatalf("Failed to create gzip reader: %v", err)
	}
	defer gzipReader.Close()

	// Create a tar reader
	tarReader := tar.NewReader(gzipReader)

	// Extract files from the tar archive
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			// End of archive
			break
		}
		if err != nil {
			log.Fatalf("Failed to read tar header: %v", err)
		}

		// Calculate the target path
		targetPath := filepath.Join(extractFolder, header.Name)

		// Create directories as needed
		if header.Typeflag == tar.TypeDir {
			err := os.MkdirAll(targetPath, 0755)
			if err != nil {
				log.Fatalf("Failed to create directory: %v", err)
			}
			continue
		}

		// Create the file
		file, err := os.Create(targetPath)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
		defer file.Close()

		// Write the file contents
		if _, err := io.Copy(file, tarReader); err != nil {
			log.Fatalf("Failed to write file: %v", err)
		}

		fmt.Printf("Extracted: %s\n", targetPath)
	}

	fmt.Println("Extraction completed successfully.")
}



func ZippingFolderLinuxLocal() {
	// Folder to be archived
	sourceFolder := "D:/tomcat"

	// Tar file to be created
	tarFilePath := "D:/GoLangTask/Demo.tar.gz"

	// Create a new tar.gz file
	tarFile, err := os.Create(tarFilePath)
	if err != nil {
		log.Fatalf("Failed to create tar file: %v", err)
	}
	defer tarFile.Close()

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(tarFile)
	defer gzipWriter.Close()

	// Create a new tar writer
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// Walk through the source folder and add files to the tar archive
	err = filepath.Walk(sourceFolder, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the relative path
		relPath, err := filepath.Rel(sourceFolder, filePath)
		if err != nil {
			return err
		}

		// Create a tar header
		header, err := tar.FileInfoHeader(fileInfo, relPath)
		if err != nil {
			return err
		}

		// Write the header to the tar archive
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// If the file is not a directory, write its contents to the tar archive
		if !fileInfo.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			// Copy file contents to the tar archive
			if _, err := io.Copy(tarWriter, file); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to add files to tar archive: %v", err)
	}

	fmt.Printf("Folder '%s' zipped successfully to '%s'.\n", sourceFolder, tarFilePath)
}



func UnzippingFolderLinuxLocal() {
	// Folder to be archived
	sourceFolder := "/path/to/source/folder"

	// Tar file to be created
	tarFilePath := "/path/to/output/Demo.tar.gz"

	// Create a new tar.gz file
	tarFile, err := os.Create(tarFilePath)
	if err != nil {
		log.Fatalf("Failed to create tar file: %v", err)
	}
	defer tarFile.Close()

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(tarFile)
	defer gzipWriter.Close()

	// Create a new tar writer
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// Walk through the source folder and add files to the tar archive
	err = filepath.Walk(sourceFolder, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the relative path
		relPath, err := filepath.Rel(sourceFolder, filePath)
		if err != nil {
			return err
		}

		// Create a tar header
		header, err := tar.FileInfoHeader(fileInfo, relPath)
		if err != nil {
			return err
		}

		// Write the header to the tar archive
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// If the file is not a directory, write its contents to the tar archive
		if !fileInfo.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			// Copy file contents to the tar archive
			if _, err := io.Copy(tarWriter, file); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to add files to tar archive: %v", err)
	}

	fmt.Printf("Folder '%s' zipped successfully to '%s'.\n", sourceFolder, tarFilePath)
}










func main (){
	UnzippingFolderWindowsLocal()
}