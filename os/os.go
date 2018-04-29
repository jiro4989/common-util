package os

import (
	"os"
	"runtime"
)

// Exists はパスの有無を返します。
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// GetEnvHome は環境変数HOMEの値を返す。
// Windows環境ではUSERPROFILEの値を返す。
func GetenvHome() string {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		// WindowsでHOME環境変数が定義されていない場合
		home = os.Getenv("USERPROFILE")
	}
	return home
}
