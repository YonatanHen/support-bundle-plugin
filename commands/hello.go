package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

func requestBodyGenerator(filePath string) (*bytes.Buffer, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Error("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Create a buffer to store the file contents
	var requestBody bytes.Buffer
	_, err = io.Copy(&requestBody, file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return &requestBody, nil
}

func UploadCommand() components.Command {
	return components.Command{
		Name:        "upload-support-bundle",
		Description: "Upload a Support Bundle to JFrog Support SaaS instance.",
		Aliases:     []string{"usb"},
		Arguments:   getUploadArguments(),
		// Flags:       getHelloFlags(),
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
		{
			Name:        "files",
			Description: "paths to the files to upload",
		},
	}
}

type SupportBundleUploadConfiguration struct {
	ticketNumber int
	files        []string
}

func UploadCmd(c *components.Context) error {
	if len(c.Arguments) == 0 {
		message := "No arguments specified, please specify a ticket number and files."
		// You log messages using the following log levels.
		log.Info(message)

		return nil
	}

	var conf = new(SupportBundleUploadConfiguration)
	ticketNumber, err := strconv.Atoi(c.Arguments[0])
	conf.files = c.Arguments[1:]

	// Check if ticket number is an integer
	if err != nil {
		log.Error(err)
		return nil
	}

	conf.ticketNumber = ticketNumber

	const url = "https://supportlogs.jfrog.com/logs/%s/"

	fmt.Println("Ticket Number:", conf.ticketNumber)
	fmt.Println("File Paths:", conf.files)

	return nil
}
