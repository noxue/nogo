package nogo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type config struct {
	files  []string          //配置文件列表
	config map[string]string //配置内容map
}

func newConfig() *config {
	c := new(config)
	c.files = make([]string, 0)
	c.config = make(map[string]string)
	return c
}

//添加配置文件到文件列表
func (config *config) AddFile(file string) {
	file = strings.TrimSpace(file)
	config.files = append(config.files, file)
}

//从配置列表中读取所有的配置到config中
func (config *config) ReadAllConfig() error {

	for k, v := range config.files {

		fmt.Println(k, v)

		f, err := os.Open(v)

		//打开失败的话，继续打开下一个配置文件
		if err != nil {

			continue
		}

		buf := bufio.NewReader(f)

		for {
			line, err := buf.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					//如果读取玩了一个配置文件，退出内层循环
					break
				}
				//如果是其他错误，返回错误信息给调用者
				return err
			}

			line = strings.TrimSpace(line)

			//如果第一个字符是# 当作注释，不处理
			if len(line) > 0 && line[0] == '#' {
				continue
			}

			strs := strings.SplitN(line, "=", 2)
			if len(strs) == 2 {
				config.setString(strings.TrimSpace(strs[0]), strings.TrimSpace(strs[1]))
			} else {
				LogPrintln("配置格式出错:" + line)
			}

		}

	}

	return nil
}

//设置值
func (config *config) setString(k, v string) {
	config.config[k] = strings.TrimSpace(v)
}

//获取配置值
func (config *config) getString(key string) string {
	for k, v := range config.config {
		if key == k {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

func (config *config) getInt(key string) int {
	v := config.getString(key)
	if v != "" {
		iv, err := strconv.Atoi(v)
		if err != nil {
			return 0
		}
		return iv
	}

	return 0
}
