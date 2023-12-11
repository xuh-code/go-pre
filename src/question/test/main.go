package main

/*
#cgo LDFLAGS: -lpcre2-8
#cgo CFLAGS: -DPCRE2_CODE_UNIT_WIDTH=8
#include <pcre2.h>
#include <stdio.h>
#include <string.h>
#include <errno.h>
struct Res {
	int start;
	int end;
};
struct Res  compile_regex_and_match(char *pattern,  char *subject) {
	pcre2_code *re;
	PCRE2_SIZE erroroffset;
	PCRE2_SIZE *ovector;
	int errornumber;
	int rc;

	PCRE2_SPTR p = (PCRE2_SPTR)pattern;
	PCRE2_SPTR s = (PCRE2_SPTR)subject;
	size_t subject_length = strlen(subject);

	pcre2_match_data *match_data;
   	re = pcre2_compile(
	  	p,
		PCRE2_ZERO_TERMINATED,
		0,
		&errornumber,
		&erroroffset,
		NULL);

	if (re == NULL)
	{
		printf("PCRE2 compilation failed at offset %d\n", (int)erroroffset);
		errno = EINVAL;
		struct  Res res = {0, 0};
        return res;
	}

	match_data = pcre2_match_data_create_from_pattern(re, NULL);
	rc = pcre2_match(
		re,
		s,
		subject_length,
		0,
		0,
		match_data,
		NULL);

	if (rc < 0)
	  {
	  	switch(rc)
			{
			case PCRE2_ERROR_NOMATCH: printf("No match\n"); errno=ENODATA;break;
			default: printf("Matching error %d\n", rc);errno=ENOEXEC; break;
			}
		pcre2_match_data_free(match_data);
		pcre2_code_free(re);
        //return NULL;
		struct Res res = {0, 0};
        return res;
	}
	ovector = pcre2_get_ovector_pointer(match_data);
	pcre2_match_data_free(match_data);
	pcre2_code_free(re);
	PCRE2_SPTR substring_start = s + ovector[0];
	size_t substring_length = ovector[1] - ovector[0];


	struct Res res = {(int)ovector[0], (int)substring_length};
	return res;
}
*/
import "C"

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"unsafe"
)

func main() {
	// 目标文本
	target := "a;jhgoqoghqoj0329 u0tyu10hg0h9Y0Y9827342482y(Y0y(G)_)lajf;lqjfgqhgpqjopjqa=)*(^!@#$%^&*())9999999"
	// 匹配目标文本
	subject := C.CString(target)
	defer C.free(unsafe.Pointer(subject))

	// 筛选规则
	pattern := `(?<=\d{4})[^0-9\s]{3,11}(?!\s|$)`
	// 编译正则表达式
	cPattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cPattern))

	// 获取匹配的数据
	reC, cErr := C.compile_regex_and_match(cPattern, subject)
	if cErr != nil {
		fmt.Println(cErr)
		os.Exit(1)
	}
	// 提取结果字符串
	result := target[reC.start : reC.start+reC.end]
	// 检查结果字符串是否符合规则
	validPattern := `^\D{3,11}$`
	validRegex, err := regexp.Compile(validPattern)
	if err != nil {
		fmt.Printf("Error compiling valid regex: %v\n", err)
		os.Exit(1)
	}
	if !validRegex.MatchString(result) {
		fmt.Printf("Result string '%s' does not match valid pattern '%s'\n", result, validPattern)
		os.Exit(1)
	}

	// 发送结果字符串给 bash 脚本
	conn, errs := net.Dial("udp", "127.0.0.1:12345")
	if errs != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	_, errs = conn.Write([]byte(result))
	if errs != nil {
		fmt.Fprintf(os.Stderr, "Failed to send message: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Result sent successfully")
}
