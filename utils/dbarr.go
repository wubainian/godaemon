package utils

import "strings"

const (
	DBARR_SPLIT_STR = "&;"
)

func DBArrAppend(oldString string, newArr []string) string {
	oldArr := DBArrFromString(oldString)
	var newAddArr []string

	for _, newItem := range newArr {
		isHas := false
		for _, item := range oldArr {
			if item == newItem {
				isHas = true
				break
			}
		}
		if !isHas {
			newAddArr = append(newAddArr, newItem)
		}
	}
	if len(newAddArr) > 0 {
		oldArr = append(oldArr, newAddArr...)
	}
	return DBArrToString(oldArr)
}

func DBArrFromString(str string) []string {
	if str == "" {
		return []string{}
	}
	arr := strings.Split(str, DBARR_SPLIT_STR)
	return arr
}

func DBArrToString(arr []string) string {
	return strings.Join(arr, DBARR_SPLIT_STR)
}
