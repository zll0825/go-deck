package util

import (
	"fmt"
	"reflect"
)

//DiffStruct 比较相同类型结构体的不同的值
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

func GetDifference(slice1, slice2 []int) (in1NotIn2, in2NotIn1 []int) {
	m := make(map[int]int)

	inter := intersect(slice1, slice2)
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

func intersect(slice1 []int, slice2 []int) []int {
	if len(slice1) > len(slice2) {
		return intersect(slice2, slice1)
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