package utils

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
)

func PrintBanner(addr string, env string) {
	title := color.New(color.FgCyan, color.Bold)
	info := color.New(color.FgWhite)
	muted := color.New(color.FgHiBlack)
	success := color.New(color.FgGreen)

	fmt.Println()

	title.Println("████████╗ ██████╗ ██████╗  ██████╗ ")
	title.Println("╚══██╔══╝██╔═══██╗██╔══██╗██╔═══██╗")
	title.Println("   ██║   ██║   ██║██║  ██║██║   ██║")
	title.Println("   ██║   ██║   ██║██║  ██║██║   ██║")
	title.Println("   ██║   ╚██████╔╝██████╔╝╚██████╔╝")
	title.Println("   ╚═╝    ╚═════╝ ╚═════╝  ╚═════╝ ")

	info.Println("           TodoList API")

	muted.Println("────────────────────────────────────────")

	fmt.Printf("\t%s %s\n", muted.Sprint("Runtime     →"), info.Sprint(runtime.Version()))
	fmt.Printf("\t%s %s\n", muted.Sprint("Database    →"), success.Sprint("PostgreSQL"))
	fmt.Printf("\t%s %s\n", muted.Sprint("Environment →"), success.Sprint(env))

	muted.Println("────────────────────────────────────────")

	fmt.Printf("%s %s\n",
		success.Sprint("➜ Server running at"),
		color.CyanString("http://%s", addr),
	)

	fmt.Println()
}
