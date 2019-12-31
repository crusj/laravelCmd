package writer

import (
	"fmt"
	myFile "github.com/crusj/file"
	"github.com/crusj/logger"
	"math/rand"
	"time"
)

type Empty struct{}

type Between struct {
	StartFlag,
	EndFlag,
	startTag,
	endTag string
}
type Writer interface {
	FilePath() string
	Between() *Between
	Content([]Route) []string
}
func Write(parser Parser, writer Writer) error {
	file, err := myFile.NewFile(writer.FilePath())
	if err != nil {
		return err
	}
	between := writer.Between()
	err = file.Scan([]myFile.AddTag{between.addStartTag, between.addEndTag}...)
	if err != nil {
		return err
	}
	err = file.InsertBetween(between.startTag, between.endTag, writer.Content(parser.Parse()))
	if err != nil {
		return err
	}
	logger.Info("ok 写入成功")
	return nil
}
func (between *Between) addStartTag(line int, content string) string {
	if content == between.StartFlag {
		rand.Seed(time.Now().Unix())
		between.startTag = fmt.Sprintf("%d_%s", rand.Intn(1000000), between.StartFlag)
		return between.startTag
	}
	return ""
}
func (between *Between) addEndTag(line int, content string) string {
	if content == between.EndFlag {
		rand.Seed(time.Now().Unix())
		between.endTag = fmt.Sprintf("%d_%s", rand.Intn(1000000), between.EndFlag)
		return between.endTag
	}
	return ""
}
