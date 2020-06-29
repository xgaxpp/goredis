package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func init() {

}
// Execute 启动进程
func Execute(host string, port uint) {
	for {
		fmt.Printf("%s:%d>", host, port)

		bio := bufio.NewReader(os.Stdin)
		inputBytes, _, err := bio.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		input := string(inputBytes)
		commod := strings.Replace(input, " ", "", -1)
		if commod == "exit" || commod == "quit" {
			return
		}
		InitCli(analysisCmd(input))
	}
}
// 解析命令行
func analysisCmd(input string) []string {
	// 去掉收尾空格
	input = strings.TrimSpace(input)
	// 去掉无用空格
	input = deleteExtraSpace(input)
    // 根据空格分组
	cmds := strings.Split(input, " ")
	// 返回结果
	return cmds
}

/*
函数名：delete_extra_space(s string) string
功  能:删除字符串中多余的空格(含tab)，有多个空格时，仅保留一个空格，同时将字符串中的tab换为空格
参  数:s string:原始字符串
返回值:string:删除多余空格后的字符串
创建时间:2018年12月3日
修订信息:
*/
func deleteExtraSpace(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "	", " ", -1)    //替换tab为空格
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}