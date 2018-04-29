// config は設定ファイルを扱います。
package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	jos "github.com/jiro4989/common-util/os"
)

const (
	// ConfigDir は設定ファイルの配置先のディレクトリ名です。
	ConfigDir = ".config"
)

// GetConfigDir は設定ファイルの配置場所ディレクトリのパスを返す。
// $HOME/.config というディレクトリを返す。
// Windows環境では$HOMEの代わりに%USERPROFILE%を使用する。
func GetConfigDir() string {
	home := jos.GetenvHome()
	home = filepath.Join(home, ConfigDir)
	return home
}

// GetAppConfigDir はアプリ設定ファイルの配置場所ディレクトリのパスを返す。
// $HOME/.config/{appname} というディレクトリを返す。
// Windows環境では$HOMEの代わりに%USERPROFILE%を使用する。
func GetAppConfigDir(appname string) string {
	return filepath.Join(GetConfigDir(), appname)
}

// MkAppConfigDir はアプリ設定ファイルの保存先ディレクトリを作成する。
// $HOME/.config/{appname} というディレクトリを作成する。
// Windows環境では$HOMEの代わりに%USERPROFILE%を使用する。
func MkAppConfigDir(appname string) error {
	dir := GetAppConfigDir(appname)
	return os.MkdirAll(dir, os.ModeDir)
}

// RemoveAppConfigDir はアプリ設定ファイルの保存先ディレクトリを削除する。
// 中に存在するファイルごとディレクトリを削除する。
// $HOME/.config/{appname} というディレクトリを削除する。
// Windows環境では$HOMEの代わりに%USERPROFILE%を使用する。
func RemoveAppConfigDir(appname string) error {
	dir := GetAppConfigDir(appname)
	return os.RemoveAll(dir)
}

// ReadConfigFile は設定ファイルからデータを取得する。
// $HOME/.config/{appname}/{filename} というファイルを読み取る。
// Windows環境では$HOMEの代わりに%USERPROFILE%を使用する。
func ReadConfigFile(appname, filename string) ([]byte, error) {
	dir := GetConfigDir()
	path := filepath.Join(dir, appname, filename)
	return ioutil.ReadFile(path)
}

// WriteConfigFile は設定ファイルを作成します。
// $HOME/.config/{appname}/{filename} というファイルを作成する。
// Windows環境では$HOMEの代わりに%USERPROFILE%を使用する。
func WriteConfigFile(appname, filename string, b []byte) error {
	dir := GetConfigDir()
	path := filepath.Join(dir, appname, filename)
	return ioutil.WriteFile(path, b, os.ModePerm)
}
