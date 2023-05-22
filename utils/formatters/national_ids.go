package formatters

import (
	"regexp"
	"strings"
)

func FormatRutWithoutDots(run string) string {
	rut := rutClean(run)
	result := strings.Join([]string{rut[:len(rut)-1], "-", rut[len(rut)-1:]}, "")
	return result
}

func RemoveDotsAndHyphen(run string) string {
	rut := FormatRutWithoutDots(run)
	rut = rut[:len(rut)-1]
	result := rutClean(rut)

	return result
}

func rutClean(value string) string {
	reg := regexp.MustCompile(`[^0-9kK]+`)
	value = reg.ReplaceAllString(value, "")
	return strings.ToUpper(value)
}
