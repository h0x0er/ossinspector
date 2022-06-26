package main

import (
	"fmt"

	"github.com/h0x0er/ossinspector"
)

func main() {
	conf, _ := ossinspector.NewConfig("config.yml")
	fmt.Println("", conf.PackageTrustRule)
}
