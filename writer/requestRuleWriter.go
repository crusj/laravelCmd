package writer

import (
	"fmt"
	"github.com/crusj/logger"
	"github.com/stoewer/go-strcase"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type requestRuleWriter struct {
	Dir,
	StartFlag,
	EndFlag string
}

func NewRequestRuleWriter(dir, startFlag, EndFlag string) *requestRuleWriter {
	return &requestRuleWriter{
		Dir:       dir,
		StartFlag: startFlag,
		EndFlag:   EndFlag,
	}
}
func (requestWriter *requestRuleWriter) FilesPath(routes []Route) []string {
	filesPath := make([]string, 0)

	hash := make(map[string]Empty)
	for _, route := range routes {
		switch strings.ToUpper(route.Method.String()) {
		case "POST", "PUT":
			act := strings.ToUpper(route.Method.String())
			if act == "POST" {
				act = "store_"
			} else {
				act = "update_"
			}
			fileName := strcase.UpperCamelCase(act+route.ModuleName+"_"+route.ActionName) + ".php"
			if route.ActionName == "" {
				fileName = strcase.UpperCamelCase(act+route.ModuleName) + ".php"
			}
			ruleName := strings.Split(fileName, ".")[0]
			if _, exists := hash[fileName]; exists {
				continue
			} else {
				filesPath = append(filesPath, requestWriter.Dir+fileName)
				hash[fileName] = Empty{}
			}
			//文件不存在则创建
			if _, err := os.Stat(requestWriter.Dir + fileName); os.IsNotExist(err) {
				cmd := exec.Command("php", "artisan", "make:request", ruleName)
				rsl, err := cmd.Output()
				if strings.Contains(string(rsl), "created successfully") {
					logger.Info("创建request,%s", ruleName)
				} else {
					logger.Warn("创建request失败,%s,%s", ruleName, err)
				}
			}
		default:
			continue
		}
	}
	return filesPath
}
func (requestWriter *requestRuleWriter) Between() *Between {
	return &Between{
		StartFlag: requestWriter.StartFlag,
		EndFlag:   requestWriter.EndFlag,
	}
}
func (requestWriter *requestRuleWriter) Contents(routes []Route) [][]string {

	contents := make([][]string, 0)
	for _, route := range routes {
		//function name
		methodName := strings.ToUpper(route.Method.String())

		switch methodName {
		case "POST", "PUT":
			str := make([]string, 0)
			for _, rp := range route.RequestParam {
				r := &requestRule{
					Name:     rp.Name,
					Type:     rp.Type,
					Required: rp.Required,
					Max:      0,
					Min:      0,
				}
				desS := strings.Split(rp.Desc, "|")
				if len(desS) > 1 {
					lengthLimit := strings.Split(desS[0], ",")
					if len(lengthLimit) == 2 {
						max, err := strconv.Atoi(lengthLimit[1])
						if err != nil {
							max = 0
						}
						min, err := strconv.Atoi(lengthLimit[0])
						if err != nil {
							min = 0
						}
						r.Max = max
						r.Min = min
					}
				}
				str = append(str, r.String())
			}
			contents = append(contents, str)
		default:
			continue
		}
	}
	return contents
}

type requestRule struct {
	Name     string
	Type     string
	Required bool
	Max      int
	Min      int
}

func (request *requestRule) String() string {
	space12 := strings.Repeat(" ", 12)
	request.Type = strings.ToLower(request.Type)
	str := ""
	switch request.Type {
	case "string":
		str = space12 + "'%s' => [" + request.RequiredBody() + " 'string', " + request.lengthBody() + request.telBody() + "],"
	case "int":
		str = space12 + "'%s' => [" + request.RequiredBody() + " 'integer', " + request.lengthBody() + "],"
	case "array":
		str = space12 + "'%s' => [" + request.RequiredBody() + " 'array', " + request.lengthBody() + "],"
	default:
		str = ""

	}
	if str == "" {
		return str
	} else {
		return fmt.Sprintf(str, request.Name)
	}
}
func (request *requestRule) RequiredBody() string {
	if request.Required == true {
		return "'required', "
	} else {
		return ""
	}
}

//长度
func (request *requestRule) lengthBody() string {
	if request.Max > request.Min {
		if request.Min > 0 {
			return fmt.Sprintf("'max:%d', 'min:%d',", request.Max, request.Min)
		} else {
			return fmt.Sprintf("'max:%d', ", request.Max)
		}
	} else {
		return ""
	}
}

//手机
func (request *requestRule) telBody() string {
	str := ""
	if request.Name == "tel" {
		str = `function ($attribute, $value, $fail) {
                if (preg_match("/^1[3456789]\d{9}$/", $value)) {
                    return true;
                } else {
                    $fail('手机号格式错误');
                }
            },
`
	}
	return str
}
