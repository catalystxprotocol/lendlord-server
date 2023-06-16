package tools

import "os"

// CheckFileExist 检查文件是否存在
func CheckFileExist(filePath string) (bool, error) {
	exists := true

	_, err := os.Stat(filePath)
	if err == nil {
		return exists, nil
	}

	if os.IsNotExist(err) {
		exists = false
	} else {
		// unknown err
		return false, err
	}

	return exists, nil
}

// CreateFolder 创建文件夹
func CreateFolder(folderPath string) error {
	folderExist, err := CheckFileExist(folderPath)
	if err != nil {
		return err
	}

	if !folderExist {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
