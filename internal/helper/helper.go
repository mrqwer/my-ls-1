package helper

import "my-ls-1/internal/file_ops"

func SplitOptionsAndFilenames(parameters []string) map[string][]string {
	options := []string{}
	names := []string{}
	collaborator := make(map[string][]string)

	for _, parameter := range parameters {
		if parameter[0] == '-' {
			options = append(options, parameter)
		} else {
			names = append(names, parameter)
		}
	}
	collaborator["options"] = options
	collaborator["names"] = names

	return collaborator
}

func ExtractFakeNames(parameters []string) []string {
	fakeNames := []string{}

	for _, parameter := range parameters {
		if !file_ops.FileOrDirExists(parameter) {
			fakeNames = append(fakeNames, parameter)
		}
	}
	return fakeNames
}
