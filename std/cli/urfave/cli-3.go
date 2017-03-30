package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
	"github.com/Sirupsen/logrus"
)

func main() {

	var language string

	app := cli.NewApp()

	//GLOBAL OPTIONS
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output in logs",
		},
		cli.StringFlag{
			Name:  "root",
			Usage: "containerd state directory",
			Value: "/run/containerd",
		},
		cli.StringFlag{
			Name:  "runtime",
			Usage: "default runtime for execution",
			Value: "runc",
		},
		/*
                Alternate Names
                You can set alternate (or short) names for flags by providing a comma-delimited list for the Name
		*/
		cli.StringFlag{
			Name:  "socket, s",
			Usage: "socket path for containerd's GRPC server",
			Value: "/run/containerd/containerd.sock",
		},
		cli.StringFlag{
			Name:  "metrics-address, m",
			Usage: "tcp address to serve metrics on",
			Value: "127.0.0.1:7897",
		},
		cli.StringFlag{
			Name:  "events-address, e",
			Usage: "nats address to serve events on",
			Value: nats.DefaultURL,
		},
		/*
		Placeholder Values
                Sometimes it's useful to specify a flag's value within the usage string itself.
                Such placeholders are indicated with back quotes
                Will result in help output like:
                --config FILE, -c FILE   Load configuration from FILE
                Note that only the first placeholder is used. Subsequent back-quoted words will be left as-is
		*/
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
		/*
		Destination variable for a flag
		You can also set a destination variable for a flag, to which the content will be scanned
		*/
		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &language,
		},
                /*Values from the Environment
                  You can also have the default value set from the environment via EnvVar
                  The EnvVar may also be given as a comma-delimited "cascade",
                  where the first environment variable that resolves is used as the default
                */
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "language for the greeting",
			EnvVar: "LEGACY_COMPAT_LANG,APP_LANG,LANG",
		},
	}

	/* Flags Ordering
	  Flags for the application and commands are shown in the order they are defined.
	  However, it's possible to sort them from outside this library by using FlagsByName with sort
	*/
	sort.Sort(cli.FlagsByName(app.Flags))

	app.Action = func(context *cli.Context) error {
		//context.Args()  arguments
		fmt.Printf("Hello %q", context.Args().Get(0))
		if context.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}

		name := "someone"
		if context.NArg() > 0 {
			name = context.Args()[0]
		}
		if language == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}

		return nil
	}

	app.Run(os.Args)
}


