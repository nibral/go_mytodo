package config

import (
	"log"
	"os"
	"path/filepath"
	"github.com/BurntSushi/toml"
)

type AppConfig struct {
	Database Database
}

type Database struct {
	Engine string
	Source string
}

var Config AppConfig

func init() {
	// 実行環境を取得
	env := os.Getenv("ENV")
	if len(env) == 0 {
		log.Fatalln("ENV is not defined.")
	}

	// 作業ディレクトリを取得
	exec, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	workingDir := filepath.Dir(exec)

	// 設定ファイル読み込み
	file := workingDir + "/" + env + ".conf.toml"
	_, err = toml.DecodeFile(file, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
