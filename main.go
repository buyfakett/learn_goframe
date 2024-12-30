package main

import (
	_ "learn_golang/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"learn_golang/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
