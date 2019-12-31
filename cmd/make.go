package cmd

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	myfile "github.com/crusj/file"
	"github.com/crusj/logger"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type b bool

func (it b) String() string {
	if it == true {
		return "true"
	} else {
		return "false"
	}
}

type MakeConfig struct {
	Name  string `json:"name""`
	Title string `json:"title"`
}
type Make struct {
	IsSet  bool
	Value  string
	Config []MakeConfig
}

func (it *Make) Set(value string) error {
	it.IsSet = true
	it.Value = value
	return nil
}
func (it *Make) String() string {
	return it.Value
}
func (it *Make) CheckAndRunMake() {
	if it.IsSet == true {

		//显示
		if it.Value == "list" {
			it.List()
			return
		}
		//执行某一项
		if number, err := strconv.Atoi(it.Value); err == nil {
			it.Make(number)
		}
		//执行所有项
		if it.Value == "all" {
			it.MakeAll()
		}

	}
}
func (it *Make) List() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "number"},
			{Align: simpletable.AlignLeft, Text: "module"},
			{Align: simpletable.AlignLeft, Text: "title"},
			{Align: simpletable.AlignLeft, Text: "register"},
		},
	}
	for i, item := range it.Config {
		var exists b = true
		exists = adminRouteIsRegister(item.Name)
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: strconv.Itoa(i + 1)},
			{Align: simpletable.AlignLeft, Text: item.Name},
			{Align: simpletable.AlignLeft, Text: item.Title},
			{Align: simpletable.AlignLeft, Text: exists.String()},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.SetStyle(simpletable.StyleRounded)
	fmt.Println(table.String())
}

//是否注册
func adminRouteIsRegister(serviceName string) b {
	path := "app/Admin/Controllers/" + serviceName + "Controller.php"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
func addAdminRouteStartTag(line int, content string) string {
	if content == "//路由开始" {
		return "admin_route_start";
	} else {
		return "";
	}
}
func addAdminRouteEndTag(line int, content string) string {
	if content == "//路由结束" {
		return "admin_route_end";
	} else {
		return "";
	}
}

//添加一个路由
func (it *Make) Make(number int) {
	logger.Debug("\n开始注册序号为%d的路由",number)
	for i, item := range it.Config {
		if i+1 == number {
			controller := fmt.Sprintf("%sController", item.Name)
			if adminRouteIsRegister(item.Name) == true {
				logger.Error("路由%v已经注册", item.Name)
				return
			}
			model := fmt.Sprintf("--model=App\\Http\\Models\\%s", item.Name)
			title := fmt.Sprintf("--title=%s", item.Title)
			cmds := []string{
				"artisan", "admin:make", controller, model, title,
			}
			cmd := exec.Command("php", cmds...)
			logger.Debug("执行命令:\n%s", cmd)

			output, err := cmd.Output()
			if err != nil {
				logger.Painc("执行失败命令make失败：\n,%,\n%s\n", err, output)
			} else {
				reg := `\$router->resource.*;`
				r, err := regexp.Compile(reg)
				if err != nil {
					logger.Error(err)
				}
				route := r.FindString(string(output))
				indentSpace := strings.Repeat(" ", 4)
				if route == "" {
					logger.Error("执行失败", string(output))
				} else {
					//添加路由
					f, err := myfile.NewFile("app/Admin/routes.php")
					if err != nil {
						logger.Error(err)
					} else {
						f.Scan([]myfile.AddTag{addAdminRouteStartTag, addAdminRouteEndTag}...)
						err := f.InsertBetween("admin_route_start", "admin_route_end", []string{indentSpace + route})
						if err != nil {
							logger.Error("添加路由%s,到Admin/routes.php失败,%s", route, err)
						} else {

							logger.Info("添加路由%s,到Admin/routes.php成功", route)
						}
					}
				}
			}
		}
	}
}

//添加所有路由
func (it *Make) MakeAll() {

	logger.Debug("\n开始注册所有路由")
	contents := make([]string, 0)
	indentSpace := strings.Repeat(" ", 4)
	for _, item := range it.Config {
		if adminRouteIsRegister(item.Name) == true {
			logger.Error("路由%v已经注册，将被忽略", item.Name)
			continue
		}
		controller := fmt.Sprintf("%sController", item.Name)
		model := fmt.Sprintf("--model=App\\Http\\Models\\%s", item.Name)
		title := fmt.Sprintf("--title=%s", item.Title)
		cmds := []string{
			"artisan", "admin:make", controller, model, title,
		}
		cmd := exec.Command("php", cmds...)
		logger.Debug("执行命令:\n%s\n", cmd)
		output, err := cmd.Output()
		if err != nil {
			logger.Error(err)
			continue
		}
		reg := `\$router->resource.*;`
		r, err := regexp.Compile(reg)
		if err != nil {
			logger.Error(err)
		}
		route := r.FindString(string(output))
		if route != "" {
			contents = append(contents, indentSpace+route)
		}
	}
	if len(contents) > 0 {
		f, err := myfile.NewFile("app/Admin/routes.php")
		if err != nil {
			logger.Error(err)
			return
		}
		f.Scan([]myfile.AddTag{addAdminRouteEndTag, addAdminRouteStartTag}...)
		err = f.InsertBetween("admin_route_start", "admin_route_end", contents)
		if err != nil {
			logger.Error(err)
		} else {
			for _, item := range contents {
				logger.Info("添加路由%s,到Admin/routes.php成功", item)
			}
		}
	}
}
