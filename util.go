package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

const (
	CONFIG_DIR_NAME = ".config"
)

// 設定ファイルからデータを取得する。
func ReadConfigFile(filename string) ([]byte, error) {
	home := GetEnvHome()
	path := filepath.Join(home, filename)
	return ioutil.ReadFile(path)
}

// .configまたはAPPDATA配下に引数のファイルを保存する
// 拡張子に合わせたファイルを生成する。
// 対応するファイル
// - json
func WriteConfigFile(filename string, b []byte) error {
	home := GetEnvHome()
	path := filepath.Join(home, filename)
	return ioutil.WriteFile(path, b, os.ModePerm)
}

// 設定ファイルの保存先ディレクトリを作成する
func MkConfigDir(appnm string) error {
	dir := filepath.Join(GetConfigDir(), appnm)
	err := os.MkdirAll(dir, os.ModeDir)
	return err
}

// 環境変数HOMEを取得するが、
// HOMEが空で且つ環境がWindowsならAPPDATAを使用する
func GetEnvHome() string {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		// WindowsでHOME環境変数が定義されていない場合
		home = os.Getenv("USERPROFILE")
	}
	return home
}

// 設定ファイルの配置場所ディレクトリのパスを返す。
func GetConfigDir() string {
	home := GetEnvHome()
	home = filepath.Join(home, CONFIG_DIR_NAME)
	return home
}

// パスのファイルの有無を確認する
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
