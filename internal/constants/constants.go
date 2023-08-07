package constants

import "fmt"

func ErrMessage(filename string) string {
	return fmt.Sprintf("ls: cannot access '%s': No such file or directory", filename)
}
