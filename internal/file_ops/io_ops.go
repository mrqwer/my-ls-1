package file_ops

import "os"

// func FileOrDirExists(name string) (bool, int) {
// 	_, err := os.Stat(name)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return false, 404
// 		} else if os.IsPermission(err) {
// 			return false, 403
// 		}
// 		return false, 0

// 	}
// 	return true, 1
// }
func FileOrDirExists(name string) bool {
	_, err := os.Stat(name)
	if err != nil && os.IsNotExist(err) {
			return false
	}
	return true
}