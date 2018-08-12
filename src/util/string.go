// TODO: 今のままだと全てのutility系メソッドがutil直下に生えてしまう
// 煩雑になってきたらディレクトリ分ける？
package util

import "unicode/utf8"

func BytesToString(b []byte) string {
	return string(b)
}

func StringToBytes(s string) []byte {
	return []byte(s)
}

func LengthOf(s string) int {
	return utf8.RuneCountInString(s)
}
