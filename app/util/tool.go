package util

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// 比较相同类型结构体的不同的值
func DiffStruct(before, after interface{}) (bf, af map[string]interface{}, err error) {
	bf = make(map[string]interface{}, 0)
	af = make(map[string]interface{}, 0)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var (
		t0 = reflect.TypeOf(before)
		t  = reflect.TypeOf(after)
	)
	if t0.PkgPath()+t0.Name() != t.PkgPath()+t.Name() {
		//err = er.New("比较的类型必须相等:" + t0.PkgPath() + t0.Name() + t.PkgPath() + t.Name())
		return
	}
	var (
		v0 = reflect.ValueOf(before)
		v  = reflect.ValueOf(after)
		k  = 0
	)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		f0 := v0.Field(i)
		switch f.Kind() {
		case reflect.String:
			if f.String() != "" && f.String() != f0.String() {
				bf[t0.Field(i).Tag.Get("json")] = f0.String()
				af[t0.Field(i).Tag.Get("json")] = f.String()
				k++
			}
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
			if f.Int() != 0 && f.Int() != f0.Int() {
				bf[t0.Field(i).Tag.Get("json")] = f0.Int()
				af[t0.Field(i).Tag.Get("json")] = f.Int()
				k++
			}
		default:
			s := fmt.Sprintf("%v", f0)
			s0 := fmt.Sprintf("%v", f)
			if s != s0 {
				bf[t0.Field(i).Tag.Get("json")] = s
				af[t0.Field(i).Tag.Get("json")] = s0
				k++
			}
		}
	}
	if k == 0 {
		//err = er.New("没有匹配到不同的项")
	}
	return bf, af, err
}

// 分页，计算开始结束
func Pagination(page int, size int) (offset int, limit int) {
	if page == 0 {
		page = 1
	}

	if size == 0 {
		size = 10
	}

	offset = (page - 1) * size
	limit = size
	return
}

// 比较字符串类型切片
func GetStringDifference(slice1, slice2 []string) (in1NotIn2, in2NotIn1 []string) {
	m := make(map[string]int)

	inter := intersectString(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, v := range slice1 {
		if m[v] == 0 {
			in1NotIn2 = append(in1NotIn2, v)
		}
	}

	for _, v := range slice2 {
		if m[v] == 0 {
			in2NotIn1 = append(in2NotIn1, v)
		}
	}
	return
}

// 比较int类型切片
func GetIntDifference(slice1, slice2 []int) (in1NotIn2, in2NotIn1 []int) {
	m := make(map[int]int)

	inter := intersectInt(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, v := range slice1 {
		if m[v] == 0 {
			in1NotIn2 = append(in1NotIn2, v)
		}
	}

	for _, v := range slice2 {
		if m[v] == 0 {
			in2NotIn1 = append(in2NotIn1, v)
		}
	}
	return
}

func intersectInt(slice1 []int, slice2 []int) []int {
	if len(slice1) > len(slice2) {
		return intersectInt(slice2, slice1)
	}
	m := map[int]int{}
	for _, num := range slice1 {
		m[num]++
	}

	var intersection []int
	for _, num := range slice2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}

func intersectString(slice1 []string, slice2 []string) []string {
	if len(slice1) > len(slice2) {
		return intersectString(slice2, slice1)
	}
	m := map[string]int{}
	for _, num := range slice1 {
		m[num]++
	}

	var intersection []string
	for _, num := range slice2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}

// 驼峰转蛇形 XxYy to xx_yy , XxYY to xx_y_y
func bCamel2Snake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// 蛇形转大驼峰 xx_yy to XxYx  xx_y_y to XxYY
func Snake2BCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// 蛇形转小驼峰 xx_yy to xxYx  xx_y_y to xxYY
func Snake2SCamel(s string) string {
	return LcFirst(Snake2BCamel(s))
}

// 首字母大写
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 首字母小写
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}