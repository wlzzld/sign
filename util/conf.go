package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var confMap map[string]string = make(map[string]string)

func InitConfig(path string) {
	loadConfigFromFile(path)
}
func SetValue(key, value string) {
	confMap[key] = value
}

func GetValInt(key string, defaultVal ...int) (int, error) {
	env, flag := os.LookupEnv(key)
	if flag == true {
		evnVal, err := strconv.Atoi(env)
		if err != nil {
			log.Printf("UseConfParameter,key: %s,value:%s,from:environment variable", key, env)
			return evnVal, nil
		}
	}

	val := confMap[key]
	if defaultVal != nil && len(defaultVal) > 0 {
		if val == "" {
			log.Printf("UseConfParameter,key: %s,value:%d,from:defaultVal", key, defaultVal[0])
			return defaultVal[0], nil
		}
	}
	if val == "" {
		return 0, nil
	}
	v, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("UseConfParameter error,key:%s,%s It's not a number type!", key, val)
		return 0, err

	}
	return v, nil
}
func GetValStr(key string, defaultVal ...string) string {
	env, flag := os.LookupEnv(key)
	if flag == true {
		log.Printf("UseConfParameter,key: %s,value:%s,from:environment variable", key, env)
		return env
	}

	val := confMap[key]
	if defaultVal != nil && len(defaultVal) > 0 {
		if val == "" {
			log.Printf("UseConfParameter,key: %s,value:%s,from:defaultVal", key, defaultVal[0])
			return defaultVal[0]
		}
	}
	return val
}
func loadConfigFromFile(path string) {
	log.Println("Initialize parameters from the ini file ", path)
	//打开文件指定目录，返回一个文件f和错误信息
	f, err := os.Open(path)
	defer f.Close()

	//异常处理 以及确保函数结尾关闭文件流
	if err != nil {
		log.Println(path, " is not exist")
		return
	}

	//创建一个输出流向该文件的缓冲流*Reader
	r := bufio.NewReader(f)
	for {
		//读取，返回[]byte 单行切片给b
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		//去除单行属性两端的空格
		s := strings.TrimSpace(string(b))
		//fmt.Println(s)

		//判断等号=在该行的位置
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		//取得等号左边的key值，判断是否为空
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}

		//取得等号右边的value值，判断是否为空
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		//这样就成功吧配置文件里的属性key=value对，成功载入到内存中c对象里
		confMap[key] = value
	}
}
