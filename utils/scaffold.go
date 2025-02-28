package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"scaffolder/helper"

	"gopkg.in/yaml.v3"
)

/*
Scaffold creates a directory structure based on the provided YAML file.
It reads the YAML file at the given 'yamlpath', unmarshals its contents,
and creates the specified folders and files accordingly in the current directory.

Parameters:

  - name: The name of the project folder to be created.

  - yamlpath: The path to the YAML file containing the directory structure information.

Example Usage:

	projectName := "my_project"
	yamlFilePath := "/path/to/structure.yaml"
	Scaffold(projectName, yamlFilePath)

	// Result: The directory structure defined in the "structure.yaml" file will be created in the "my_project" directory.

Note:

  - The function assumes that the 'yamlpath' points to a valid YAML file that adheres to the expected format.

  - It will create the necessary folders and subdirectories as specified in the YAML file.

  - If any error occurs during the process, the function will log a fatal error message and exit the program.
*/
func Scaffold(name string, yamlpath string) {
	// Read and get YAML data
	yamlData, err := os.ReadFile(yamlpath)
	helper.Fatal(fmt.Sprintf("Failed to read YAML file: %s", err), true, err)

	// Create map for the directory structure
	var dirs map[string]map[string]string

	// Unmarshal the YAML into our map
	err = yaml.Unmarshal(yamlData, &dirs)
	helper.Fatal(fmt.Sprintf("Error unmarshalling YAML: %s", err), true, err)

	// Create project folder
	err = os.Mkdir(name, 0755)
	helper.Fatal(fmt.Sprintf("Error creating project folder: %s", err), true, err)

	// Navigate to the project folder
	err = os.Chdir(name)
	helper.Fatal(fmt.Sprintf("Failed to navigate to project folder: %s", err), true, err)

	// Scaffold the directory structure :: iterating over the map
	for folder, files := range dirs {
		// Create the folders and subdirectories if necessary
		err = os.MkdirAll(folder, 0755)
		helper.Fatal(fmt.Sprintf("Error creating folder %s: %v", folder, err), true, err)

		// Create the files :: iterating over files from the map and getting specified content
		for fileName, content := range files {
			// Construct a file path for the file
			filePath := filepath.Join(folder, fileName)
			// Create the directories before creating the file
			err = os.MkdirAll(filepath.Dir(filePath), 0755)
			helper.Fatal(fmt.Sprintf("Error creating directories for %s: %v", filePath, err), true, err)

			// Create the files at filePath path and add specified content
			err = os.WriteFile(filePath, []byte(content), 0644)
			helper.Fatal(fmt.Sprintf("Failed to create file %s: %s", fileName, err), true, err)
		}
	}
}
