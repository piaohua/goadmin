package libs

import (
	"sort"
	"strings"
)

// SortAndConcat sort the map by key in ASCII order,
// and concat it in form of "k1=v1&k2=2"
func SortAndConcat(param map[string]string) string {
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	var sortedParam []string
	sort.Strings(keys)
	for _, k := range keys {
		// fmt.Println(k, "=", param[k])
		sortedParam = append(sortedParam, k+"="+param[k])
	}
	return strings.Join(sortedParam, "&")
}

// Sign the parameter in form of map[string]string with app key.
// Empty string and "sign" key is excluded before sign.
// Please refer to http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=4_3
func Sign(param map[string]string, key string) string {
	newMap := make(map[string]string)
	// fmt.Printf("%#v\n", param)
	for k, v := range param {
		if k == "sign" {
			continue
		}
		if v == "" {
			continue
		}
		newMap[k] = v
	}
	// fmt.Printf("%#v\n\n", newMap)
	preSignStr := SortAndConcat(newMap)
	preSignWithKey := preSignStr + "&key=" + key
	//return fmt.Sprintf("%X", md5.Sum([]byte(preSignWithKey)))
	return strings.ToUpper(MD5(preSignWithKey))
}
