// TODO: 今のままだと全てのutility系メソッドがutil直下に生えてしまう
// 煩雑になってきたらディレクトリ分ける？
package util

import "unsafe"

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
