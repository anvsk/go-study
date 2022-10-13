package ticket

import (
	"go-study/pkg/util"
	"os"

	"github.com/akamensky/argparse"
)

var Config util.TicketConfig

func initFlagArgs() {
	Config = util.Config.TicketConfig
	parser := argparse.NewParser("ticket", "get ticket of shengsi")
	from := parser.String("f", "from", &argparse.Options{Default: Config.From})
	to := parser.String("t", "to", &argparse.Options{Default: Config.To})
	date := parser.String("d", "date", &argparse.Options{Default: Config.Date})
	mst := parser.String("m", "mst", &argparse.Options{Default: Config.Customization.MinShipTime})
	lst := parser.String("e", "lst", &argparse.Options{Default: Config.Customization.LatestShipTime})
	line := parser.String("l", "line", &argparse.Options{Default: Config.Customization.LineNum})
	class := parser.String("c", "class_name", &argparse.Options{Default: Config.Customization.Class})
	err := parser.Parse(os.Args)
	if err != nil {
		panic(err)
	}
	Config.Customization = util.Customization{
		From:           *from,
		To:             *to,
		Date:           *date,
		MinShipTime:    *mst,
		LatestShipTime: *lst,
		LineNum:        *line,
		Class:          *class,
	}
	// log.DebugDump(Config, "config")
}
