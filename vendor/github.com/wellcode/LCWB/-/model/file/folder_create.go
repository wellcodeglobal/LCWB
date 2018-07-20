package file

import (
	//"fmt"
	"os"
	"path/filepath"
)

func CreatePreviewFolder(path, web_name string) {
	if _, err := os.Stat(path); err == nil {
		RemoveContents(path)
	}

	css_path := path + "/css"
	js_path := path + "/js"

	os.Mkdir(path, 0777)
	os.Mkdir(js_path, 0777)
	os.Mkdir(css_path, 0777)
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
