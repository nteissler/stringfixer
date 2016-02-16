package stringfixer

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

func FirstCapital(s string) string {
	if len(s) < 1 {
		return s
	}
	t := strings.Trim(s, " ")
	t = strings.ToLower(t)
	a := []rune(t)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

func DeleteExtension(s string) (string, error) {
	// check that a non empty extension exists
	if dot := strings.LastIndex(s, "."); dot == utf8.RuneCountInString(s)-1 || dot == -1 {
		return s, errors.New("Passed in string doesn't contain non-empty extension")
	}
	return fmt.Sprint(strings.TrimSuffix(filepath.Base(s), filepath.Ext(s))), nil
}
