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
		Arguments:   getUploadArguments(),
		Flags:       getHelloFlags(),
		// EnvVars:     getHelloEnvVar(),
		Action: func(c *components.Context) error {
			return UploadCmd(c)
		},
	}
}

func getUploadArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "ticket number",
			Description: "The support ticket number in JFrog portal.",
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
	// TODO: change the message accordingly
	if err != nil {
		log.Error(err)
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
