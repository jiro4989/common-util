package os

import (
	"log"
	"os"
	"runtime"
	"testing"
)

func TestExists(t *testing.T) {
	existFile := "os.go"
	if !Exists(existFile) {
		log.Fatal(existFile + "が存在しません。")
	}

	notExistFile := "hogepiyo.notexist"
	if Exists(notExistFile) {
		log.Fatal("存在しないはずのファイル" + notExistFile + "が存在した判定になりました。")
	}
}

func TestGetenvHome(t *testing.T) {
	const msg = "環境変数HOMEの値が一致しませんでした。"
	if runtime.GOOS == "windows" {
		if GetenvHome() != os.Getenv("USERPROFILE") {
			log.Fatal(msg)
		}
		return
	}
	if GetenvHome() != os.Getenv("HOME") {
		log.Fatal(msg)
		return
	}
}
