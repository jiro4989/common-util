package util

import (
	"os"
	"path/filepath"
	"runtime"
)

// 設定ファイルの保存先ディレクトリを作成する
func MkdirConfigDir(appnm string) error {
	home := getHomeEnv()
	// 保存先ディレクトリ
	dir := filepath.Join(home, dcfg, appnm)
	err := os.MkdirAll(dir, os.ModeDir)
	return err
}

// 環境変数HOMEを取得するが、
// HOMEが空で且つ環境がWindowsならAPPDATAを使用する
func getHomeEnv() string {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		// WindowsでHOME環境変数が定義されていない場合
		home = os.Getenv("APPDATA")
	}
	return home
}
