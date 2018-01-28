package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

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
	home := GetEnvHome()
	// 保存先ディレクトリ
	dir := filepath.Join(home, appnm)
	err := os.MkdirAll(dir, os.ModeDir)
	return err
}

// 環境変数HOMEを取得するが、
// HOMEが空で且つ環境がWindowsならAPPDATAを使用する
func GetEnvHome() string {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		// WindowsでHOME環境変数が定義されていない場合
		home = os.Getenv("APPDATA")
	}
	return home
}

// パスのファイルの有無を確認する
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
