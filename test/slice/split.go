package slice

import (
	"strings"
)

// 单个字符版
// func Split(s string, symbol string) []string {
// 	tmps := ""        //临时结果元素
// 	arr := []string{} //结果集
// 	for _, v := range s {
// 		if string(v) == (symbol) {
// 			if tmps != "" {
// 				arr = append(arr, tmps)
// 				tmps = ""
// 			}
// 		} else {
// 			tmps = tmps + string(v)
// 		}
// 	}
// 	if tmps != "" {
// 		arr = append(arr, tmps)
// 	}
// 	return arr
// }

// 多个字符版[不借用其他包函数]
func Split(s string, symbol string) []string {
	tmps := ""        //临时结果元素
	tmpl := ""        //临时分割元素
	arr := []string{} //结果集
	for _, v := range s {

		if string(v) == string(symbol[len(tmpl)]) {
			// 符合分割符
			tmpl = tmpl + string(v)
			if tmpl == symbol {
				// 相同了开始分配和清空
				if tmps != "" {
					arr = append(arr, tmps)
				}
				tmps = ""
				tmpl = ""
			}
		} else {
			// 不同了追加临时分割和清空
			tmps = tmps + tmpl + string(v)
			tmpl = ""
		}
	}
	// 末尾遗留
	if tmpl != "" {
		tmps = tmps + tmpl
	}
	// 最后一个结果
	if tmps != "" {
		arr = append(arr, tmps)
	}

	return arr
}

// 借用strings包Index
func Split2(s string, symbol string) []string {

	arr := []string{}
	i := strings.Index(s, symbol)
	for i != -1 {
		if s[:i] != "" {
			arr = append(arr, s[:i])
		}
		s = s[i+len(symbol):]
		i = strings.Index(s, symbol)
	}
	arr = append(arr, s)
	return arr
}

func Split3(s string, symbol string) []string {
	arr := []string{}
	i := strings.Index(s, symbol)
	for i != -1 {
		if s[:i] != "" {
			arr = append(arr, s[:i])
		}
		s = s[i+len(symbol):]
		i = strings.Index(s, symbol)
	}
	arr = append(arr, s)
	return arr
}
