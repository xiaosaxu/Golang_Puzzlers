package lib

import (
	"os"
	in "puzzlers/m1c1s3/q4/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
