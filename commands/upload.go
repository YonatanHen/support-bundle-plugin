package commands

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

func UploadCommand() components.Command {
	return components.Command{
		Name:        "upload",
		Description: "Upload files to JFrog Support SaaS instance.",
		Aliases:     []string{"u"},
		Arguments:   getUploadArguments(),
		Action: func(c *components.Context) error {
			return UploadCmd(c)
		},
	}
}

func getUploadArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "ticket number",
			Description: "The support ticket number in JFrog Portal.",
		},
		{
			Name:        "files",
			Description: "paths to the files to upload.",
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

	uploadURL := "https://supportlogs.jfrog.com/logs/" + strconv.Itoa(conf.ticketNumber) + "/"

	for i := 0; i < len(conf.files); i++ {

		cmd := CreateCmdCommand(conf.files[i], uploadURL)

		if err := cmd.Run(); err != nil {
			fmt.Println("Error running curl:", err)
			return nil
		}

		// Check the exit status
		if cmd.ProcessState.ExitCode() == 0 {
			log.Debug("curl command completed successfully.")
		} else {
			log.Info("curl command failed with exit code", cmd.ProcessState.ExitCode())
			return nil
		}

		log.Info("Uploaded file: " + conf.files[i] + ".")
	}

	return nil
}

func CreateCmdCommand(filePath string, uploadURL string) *exec.Cmd {
	return exec.Command("curl", "-T", filePath, uploadURL)
}
