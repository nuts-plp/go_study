package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func MapTo(filePath string, data interface{}) (err error) {
	// 判断参数是否为指针或结构体类型
	dataT := reflect.TypeOf(data)
	dataV := reflect.ValueOf(data)
	if dataT.Kind() != reflect.Ptr && dataT.Kind() != reflect.Struct {
		err = errors.Errorf("should be struct or ptr")
		return
	}
	//打开文件
	var file []byte
	file, err = ioutil.ReadFile(filePath)
	if err != nil {
		err = fmt.Errorf("open ini file failed ")
		return
	}
	//去空格和分行
	fileObj := strings.TrimSpace(string(file))
	lineSlice := strings.Split(fileObj, "\r\n")
	//定义节点和节点对应的结构体名称
	var (
		section    string
		structName string
	)
	//一行一行的遍历配置文件
	for _, line := range lineSlice {
		//如果是以[开头、]结尾并且长度不为2
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") && len(line) != 2 {
			section = line[1 : len(line)-1]
			//在一级结构体中找到对应的字段赋值给structName
			for i := 0; i < dataT.Elem().NumField(); i++ {
				field := dataT.Elem().Field(i)
				if field.Tag.Get("ini") == section {
					structName = field.Name
					break
				}
			}
			continue
		}
		if strings.Index(line, "=") != -1 && !strings.HasSuffix(line, "]") && !strings.HasPrefix(line, "[") {
			inde := strings.Index(line, "=")
			key := line[:inde]
			value := line[inde+1:]
			//根据structName找对应的一级结构体的对应域
			fValue := dataV.Elem().FieldByName(structName)

			fType := fValue.Type()
			//遍历找到的域对应结构体中的域---即二级结构体中的域 并获取域名
			var ffName string
			for i := 0; i < fType.NumField(); i++ {
				field := fType.Field(i)
				//找到二级结构体中key对应的域
				if field.Tag.Get("ini") == key {
					ffName = field.Name
					break
				}
			}
			//根据对应的域名，获取对应域的值
			ffValue := fValue.FieldByName(ffName)
			//判断该域的类型，并赋值
			switch ffValue.Kind() {
			case reflect.String:
				ffValue.SetString(value)
			case reflect.Bool:
				valueBool, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				ffValue.SetBool(valueBool)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valueInt, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				ffValue.SetInt(valueInt)
			case reflect.Float32, reflect.Float64:
				valueFloat, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				ffValue.SetFloat(valueFloat)
			}
		}

	}
	return

}

type conf struct {
	mysql `ini:"mysql"`
}
type mysql struct {
	Ip       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
}

func main() {
	var config conf
	err := MapTo("./config.ini", &config)
	if err != nil {
		fmt.Println("reflect failed  error:", err)
		return
	}

	fmt.Println(config.Ip, config.User, config.Port, config.Password)
}
