package schemaregistry

import (
	"fmt"
	"go.dfds.cloud/utils/config"
	"os"
)

const SCHEMA_FILE_LOCATION = "PROVIDER_CONFLUENT_SCHEMA_FILE_LOCATION"

var (
	SchemaFileLocation = config.GetEnvValue(SCHEMA_FILE_LOCATION, "")
)

func CreateFile(content []byte, name string, path string) (string, error) {
	filePath := path

	if len(path) == 0 {
		filePath = SchemaFileLocation
	}

	fullFilePath := fmt.Sprintf("%s/%s", filePath, name)

	err := os.WriteFile(fullFilePath, content, 0770)
	if err != nil {
		return fullFilePath, err
	}

	return fullFilePath, nil
}

func RemoveFile(path string) error {
	err := os.Remove(path)
	return err
}
