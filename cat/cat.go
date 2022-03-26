package cat
/* Concatenate files in "stdout". */
import(
	"os"
	"io"
	"flag"
	"fmt"
)

var(
	blockSize int
)

func Cat(p string) error {
	f, e := os.Open(p)
	if e != nil {
		return e
	}
	defer f.Close()
	fcat(f)
	return nil
}

func fcat(f *os.File) error {
	b := make([]byte, blockSize)
	for {
		n, e := f.Read(b)
		if n>0 {
			fmt.Print(string(b[:n]))
		}
		if e == io.EOF {
			break
		}else if e != nil {
			return e
		}
	}
	return nil
}

func Run(args []string) int {	
	status := 0
	arg0 := args[0]
	args = args[1:]
	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.IntVar(&blockSize, "bs", 512, "block size")
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options] [files]\n", arg0)
		flagSet.PrintDefaults()
	}
	flagSet.Parse(args)
	args = flagSet.Args()
	if len(args)>0 {
		for _, p := range args {
			e := Cat(p)
			if e != nil {
				fmt.Fprintf(os.Stderr, "%s: %s.\n", arg0, e)
				status = 1
			}
		}
	} else {
		fcat(os.Stdin)
	}
	return status
}
