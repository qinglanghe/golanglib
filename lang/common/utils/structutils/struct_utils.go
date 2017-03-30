package structutils

import (
	"strings"
	"strconv"
	"reflect"
)

func StringMapArrayToStructArray(m []map[string]string,s []interface{}) {
	if len(m) != len(s){
		panic("array length must equals")
	}
	for i := 0; i < len(m); i++ {
		StringMapToStruct(m[i],s[i])
	}
}

/**
map字符串转成struct
 */
func StringMapToStruct(m map[string]string,s interface{}) {
	mm := make(map[string][]string,len(m))
	for k,v := range m{
		mm[k] = []string{v}
	}
	StringArrayMapToStruct(mm,s)
}

/**
map字符串数组转成struct
 */
func StringArrayMapToStruct(m map[string][]string,s interface{}) {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Ptr {
		panic("second arg s not a pointer")
	}

	e := reflect.ValueOf(s).Elem()

	for i := 0; i < e.NumField(); i++ {
		val, ok := m[e.Type().Field(i).Name]
		if (ok) {
			setValue(e.Field(i), val,v)
		}else {
			col := e.Type().Field(i).Tag.Get("col")
			val, ok := m[col]
			if (ok) {
				setValue(e.Field(i), val,v)
			}
		}
	}
}

func setValue(v reflect.Value,vals []string,obj reflect.Value)  {
	switch v.Type().Kind() {
	case reflect.Int,reflect.Int8,reflect.Int16, reflect.Int32 ,reflect.Int64:
		i,_ := strconv.ParseInt(vals[0],10,0)
		v.SetInt(i)
	case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
		i,_ := strconv.ParseUint(vals[0],10,0)
		v.SetUint(i)
	case reflect.Bool:
		i,_ := strconv.ParseBool(vals[0])
		v.SetBool(i)
	case reflect.Float32,reflect.Float64:
		i,_ := strconv.ParseFloat(vals[0],64)
		v.SetFloat(i)
	case reflect.String:
		v.SetString(vals[0])
	case reflect.Slice:
		setArray(v,vals)
	default:
		if(v.Type().String() == "time.Time"){
			md := obj.MethodByName("ToTime")
			vv := md.Call([]reflect.Value{reflect.ValueOf(vals[0])})[0]
			v.Set(vv)
		}
	}
}

func setArray(v reflect.Value,vals []string)  {
	tn := v.Type().String()
	if strings.Contains(tn,"uint64") {
		ret := make([]uint64,len(vals))
		for j := 0; j < len(ret); j++ {
			i,_ := strconv.ParseUint(vals[j],10,0)
			ret[j] = i
		}
		v.Set(reflect.ValueOf(ret))
	}else if strings.Contains(tn,"int64") {
		ret := make([]int64,len(vals))
		for j := 0; j < len(ret); j++ {
			i,_ := strconv.ParseInt(vals[j],10,0)
			ret[j] = i
		}
		v.Set(reflect.ValueOf(ret))
	}else if strings.Contains(tn,"float64") {
		ret := make([]float64,len(vals))
		for j := 0; j < len(ret); j++ {
			i,_ := strconv.ParseFloat(vals[j],64)
			ret[j] = i
		}
		v.Set(reflect.ValueOf(ret))
	}else if strings.Contains(tn,"string"){
		v.Set(reflect.ValueOf(vals))
	}
}
