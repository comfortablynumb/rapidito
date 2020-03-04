package helper

import (
	"io/ioutil"
	"os"

	"github.com/comfortablynumb/rapidito/errorhandler"
	"gopkg.in/yaml.v3"
)

type FileHelper struct {
	ErrorHandler *errorhandler.ErrorHandler
}

func (f *FileHelper) ParseYaml(yamlContent string, obj interface{}) {
	err := yaml.Unmarshal([]byte(yamlContent), obj)

	if err != nil {
		f.ErrorHandler.Handle(err, "Could not parse YAML file contents.")
	}
}

func (f *FileHelper) GetFileContents(filePath string) string {
	file, err := os.Open(filePath)

	f.ErrorHandler.HandleIfError(err, "Could not open configuration file. Make sure the path is correct and that you have permissions to open it. File: %s", filePath)

	contents, err := ioutil.ReadAll(file)

	f.ErrorHandler.HandleIfError(err, "Could not read configuration file. Make sure you have permissions to read it. File: %s", filePath)

	return string(contents)
}

func (f *FileHelper) MkDirAll(path string, perm os.FileMode) {
	err := os.MkdirAll(path, perm)

	f.ErrorHandler.HandleIfError(err, "Could NOT create missing directories in path: %s", path)
}

func (f *FileHelper) FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func NewFileHelper(errorHandler *errorhandler.ErrorHandler) *FileHelper {
	return &FileHelper{
		ErrorHandler: errorHandler,
	}
}
