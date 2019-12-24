package usualRegularExpression

import (
	"errors"
	"fmt"
	"regexp"
)

type RegularExpression struct {
	Type       string
	Annotation string
	Range      []string
}

func (re RegularExpression) VerifyExpression(str string) (interface{}, error) {
	var result interface{}
	switch re.Type {
	//匹配任意数字
	case "AnyNumbers":
		compile, e := regexp.Compile("[0-9]+")
		if e != nil {
			return nil, e
		}
		result = compile.FindAllString(str, -1)
		return result, nil
	//	匹配n-m位数字，n<m
	case "nDigitalNumbers":
		if re.Range[1] < re.Range[0] && re.Range[1] != "" {
			return nil, errors.New("范围错误,wrong Range")
		}
		if re.Range[1] == re.Range[0] {
			compile, e := regexp.Compile(fmt.Sprintf("%s{%s}", "\\d", re.Range[0]))
			if e != nil {
				return nil, e
			}
			return compile.FindAllString(str, -1), nil
		}
		if re.Range[1] == "" || re.Range[1] > re.Range[0] {
			compile, e := regexp.Compile(fmt.Sprintf("%s{%s,%s}", "\\d", re.Range[0], re.Range[1]))
			if e != nil {
				panic(e)
				return nil, e
			}
			return compile.FindAllString(str, -1), nil
		}
	//	中文验证
	case "chinese":
		if len(re.Range) == 0 {
			compile, e := regexp.Compile(fmt.Sprintf("%s", "[\u4e00-\u9fa5]+"))
			if e != nil {
				return nil, e
			}
			return compile.FindAllString(str, -1), nil
		}
		if re.Range[1] < re.Range[0] && re.Range[1] != "" {
			return nil, errors.New("范围错误,wrong Range")
		}

		if re.Range[1] == re.Range[0] {
			compile, e := regexp.Compile(fmt.Sprintf("%s{%s}", "[\u4e00-\u9fa5]", re.Range[0]))
			if e != nil {
				return nil, e
			}
			return compile.FindAllString(str, -1), nil
		}
		if re.Range[1] == "" || re.Range[1] < re.Range[0] {
			compile, e := regexp.Compile(fmt.Sprintf("%s{%s,%s}", "[\u4e00-\u9fa5]", re.Range[0], re.Range[1]))
			if e != nil {
				return nil, e
			}
			return compile.FindAllString(str, -1), nil
		}
	//	邮箱验证
	case "Email":
		compile, e := regexp.Compile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
		if e != nil {
			return nil, e
		}
		return compile.FindAllString(str, -1), nil
	//	电话号码
	case "mobile":
		compile, e := regexp.Compile(`(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}`)
		if e != nil {
			return nil, e
		}
		return compile.FindAllString(str, -1), nil
	//	身份证号
	case "18DigitalIDCard":
		compile, e := regexp.Compile(`[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]`)
		if e != nil {
			return nil, e
		}
		return compile.FindAllString(str, -1), nil
	default:
		return "未知", nil
	}

	return result, nil
}

func Verify(types, annotation, verifyString string, ran []string) []string {
	verify := RegularExpression{types, annotation, ran}
	result, err := verify.VerifyExpression(verifyString)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("验证数据%s,验证%s,结果为%s", verifyString, verify.Annotation, result)
	return result.([]string)
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
