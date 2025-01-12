package date

import(
	"fmt"
	"time"
	"github.com/mojosa-software/gomtool/src/mtool"
)

func Run(flagSet *mtool.Flags) {
	flagSet.Parse()
	args := flagSet.Args()

	if len(args) > 0 {
		flagSet.Usage()
	}
	
	date := time.Now()
	
	fmt.Println(date)

}
