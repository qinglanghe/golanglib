package dbutils

import (
	"database/sql"
	"reflect"
	"bytes"
	"log"
	"lang/common/utils/structutils"
	"errors"
)

// insert sql 的缓存
var insertSql = make(map[string]string)

func Save(DB *sql.DB,f interface{})  {
	s,p := getInsert(f)
	log.Println(s,p)
	stms,error:= DB.Prepare(s)
	defer stms.Close()
	checkErr(error)
	_,err := stms.Exec(p...)
	checkErr(err)
}

func Query(DB *sql.DB,sql string,args ...interface{}) []map[string]string {
	stmt, err := DB.Prepare(sql)
	defer stmt.Close()
	checkErr(err)

	rows, err := stmt.Query(args)
	checkErr(err)
	defer rows.Close()

	result := make([]map[string]string,0,50)
	for rows.Next() {
		m := make(map[string]string)

		cols,_ := rows.Columns()
		rawResult := make([][]byte, len(cols))
		dest := make([]interface{},len(cols))
		for i, _ := range rawResult {
			dest[i] = &rawResult[i]
		}

		error := rows.Scan(dest...)
		checkErr(error)

		for i := 0; i < len(rawResult) ; i++ {
			m[cols[i]] = string(rawResult[i])
			result = append(result,m)
		}
	}
	return result
}

func QueryOne(DB *sql.DB,s interface{},sql string,args ...interface{}) error {
	stmt, err := DB.Prepare(sql)
	defer stmt.Close()
	checkErr(err)

	rows, err := stmt.Query(args...)
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		m := make(map[string]string)

		cols,_ := rows.Columns()
		rawResult := make([][]byte, len(cols))
		dest := make([]interface{},len(cols))
		for i, _ := range rawResult {
			dest[i] = &rawResult[i]
		}

		error := rows.Scan(dest...)
		checkErr(error)

		for i := 0; i < len(rawResult) ; i++ {
			m[cols[i]] = string(rawResult[i])
		}
		structutils.StringMapToStruct(m,s)
		return nil
	}
	return errors.New("data not found")
}

func getInsert(f interface{}) (string,[]interface{}) {
	//v := reflect.ValueOf(f)
	e := reflect.ValueOf(f).Elem()
	tbname := e.Type().Name()

	buf := bytes.Buffer{}
	buf.WriteString("insert into ")
	buf.WriteString(tbname)
	buf.WriteString(" (")
	params := make([]interface{},0,e.NumField())
	for i := 0; i < e.NumField(); i++ {
		t := e.Type().Field(i).Tag.Get("col")
		if (t != ""){
			if (i != 0){
				buf.WriteString(",")
			}
			buf.WriteString(" `")
			buf.WriteString(t)
			buf.WriteString("`")
			params = append(params,e.Field(i).Interface())
		}
	}
	buf.WriteString(") values(")

	for i := 0; i < len(params) ; i++ {
		if (i == len(params) - 1){
			buf.WriteString("?")
		}else {
			buf.WriteString("?,")
		}
	}

	buf.WriteString(")")

	return buf.String(),params
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}