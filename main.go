package main

import (
	"STMT/stmt"
	_ "embed"
	"os"
	"runtime"
)

//go:embed logo.txt
var logo_body []byte

func main() {
	var config_dir string
	config_file_name := "Tools.json"

	// 创建配置文件夹
	if runtime.GOOS == "windows" {
		config_dir = "C:/ProgramData/STMT/"
	} else {
		config_dir = os.Getenv("HOME") + "/.STMT/"
	}
	exist, _ := PathExists(config_dir)
	if !exist {
		os.MkdirAll(config_dir, os.ModePerm)
	}

	config_path := config_dir + config_file_name
	stmt.Cli(config_path, logo_body)
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
