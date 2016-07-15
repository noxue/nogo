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

func NewConfig() config {
	var c config
	//c.files = make([]string, 5, 5)
	c.config = make(map[string]string)
	return c
}

//添加配置文件到文件列表
func (config *config) AddFile(file string) {
	config.files = append(config.files, file)
}

//从配置列表中读取所有的配置到config中
func (config *config) ReadAllConfig() error {

	for k, v := range config.files {

		fmt.Println(k, v)

		f, err := os.Open(v)

		if err != nil {
			continue
		}

		buf := bufio.NewReader(f)

		for {
			line, err := buf.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}

			line = strings.TrimSpace(line)

			fmt.Println(line)
		}

	}

	return nil
}

//获取配置值
func (config *config) GetString(key string) string {
	for k, v := range config.config {
		if key == k {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

func (config *config) GetInt(key string) int {
	v := config.GetString(key)
	if v != "" {
		iv, err := strconv.Atoi(v)
		if err != nil {
			return 0
		}
		return iv
	}

	return 0
}
