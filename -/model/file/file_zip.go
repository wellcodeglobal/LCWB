package file

import (
	zip_impl "archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

func ZipFile(path, output string) {
	progress := func(archivePath string) {
		fmt.Sprintln(archivePath)
	}

	err := ArchiveFile(path, output, progress)
	if err != nil {
		fmt.Println("Error - ZipFile : ", err)
	}
}

func Archive(inFilePath string, writer io.Writer, progress ProgressFunc) error {
	zipWriter := zip_impl.NewWriter(writer)

	basePath := filepath.Dir(inFilePath)

	err := filepath.Walk(inFilePath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil || fileInfo.IsDir() {
			return err
		}

		relativeFilePath, err := filepath.Rel(basePath, filePath)
		if err != nil {
			return err
		}

		archivePath := path.Join(filepath.SplitList(relativeFilePath)...)

		if progress != nil {
			progress(archivePath)
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer func() {
			_ = file.Close()
		}()

		zipFileWriter, err := zipWriter.Create(archivePath)
		if err != nil {
			return err
		}

		_, err = io.Copy(zipFileWriter, file)
		return err
	})
	if err != nil {
		return err
	}

	return zipWriter.Close()
}

func ArchiveFile(inFilePath string, outFilePath string, progress ProgressFunc) error {
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = outFile.Close()
	}()

	return Archive(inFilePath, outFile, progress)
}

type ProgressFunc func(archivePath string)
