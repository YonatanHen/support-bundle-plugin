package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-core/v2/utils/coreutils"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

func UploadCommand() components.Command {
	return components.Command{
		Name:        "upload-support-bundle",
		Description: "Upload a Support Bundle to JFrog Support SaaS instance.",
		Aliases:     []string{"usb"},
		Arguments:   getHelloArguments(),
		Flags:       getHelloFlags(),
		// EnvVars:     getHelloEnvVar(),
		Action: func(c *components.Context) error {
			return UploadCmd(c)
		},
	}
}

func getHelloArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "addressee",
			Description: "The name of the person you would like to greet.",
		},
	}
}

func getHelloFlags() []components.Flag {
	return []components.Flag{
		components.BoolFlag{
			Name:         "shout",
			Description:  "Makes output uppercase.",
			DefaultValue: false,
		},
	}
}

type SupportBundleUploadConfiguration struct {
	ticketNumber int
	files        []string
}

func UploadCmd(c *components.Context) error {
	if len(c.Arguments) == 0 {
		message := "Hello :) Now try adding an argument to the 'hi' command"
		// You log messages using the following log levels.
		log.Output(message)
		log.Debug(message)
		log.Info(message)
		log.Warn(message)
		log.Error(message)
		return nil
	}

	var conf = new(SupportBundleUploadConfiguration)
	ticketNumber, err := strconv.Atoi(c.Arguments[0])
	conf.files = c.Arguments[1:]

	// Check if ticket number is an integer
	//TODO: change the message accordingly
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	conf.ticketNumber = ticketNumber

	// Check if

	fmt.Println("Ticket Number:", conf.ticketNumber)
	fmt.Println("File Paths:", conf.files)

	if os.Getenv(coreutils.LogLevel) == "" {
		message := fmt.Sprintf("Now try setting the %s environment variable to %s and run the command again", coreutils.LogLevel, "DEBUG")
		log.Info(message)
	}
	return nil
}

// func getHelloEnvVar() []components.EnvVar {
// 	return []components.EnvVar{
// 		{
// 			Name:        "HELLO_FROG_GREET_PREFIX",
// 			Default:     "A new greet from your plugin template: ",
// 			Description: "Adds a prefix to every greet.",
// 		},
// 	}
// }

// func GetHelloCommand() components.Command {
// 	return components.Command{
// 		Name:        "hello",
// 		Description: "Says Hello.",
// 		Aliases:     []string{"hi"},
// 		Arguments:   getHelloArguments(),
// 		Flags:       getHelloFlags(),
// 		// EnvVars:     getHelloEnvVar(),
// 		Action: func(c *components.Context) error {
// 			return helloCmd(c)
// 		},
// 	}
// }
