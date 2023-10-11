package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
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

	uploadURL := "https://supportlogs.jfrog.com/logs/" + strconv.Itoa(conf.ticketNumber) + "/"

	// fmt.Println("Ticket Number:", conf.ticketNumber)
	// fmt.Println("File Paths:", conf.files)
	// fmt.Println(uploadURL)

	// Genereate files to bytes
	for i := 0; i < len(conf.files); i++ {
		cmd := exec.Command("curl", "-T", conf.files[i], uploadURL)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

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

		// wg.Add(1)

		// go func(file string) {
		// requestBody, err := requestBodyGenerator(conf.files[i])
		// if err != nil {
		// 	log.Error(err)
		// 	return nil
		// }

		// // upload files
		// resp, err := http.Post(uploadURL, "application/json", requestBody)
		// if err != nil {
		// 	log.Error("Error", err)
		// 	return nil
		// }
		// if resp.Status != "201" {
		// 	log.Error("File not uploaded: ", resp.Status)
		// 	return nil
		// }

		// // defer wg.Done()

		// // }(conf.files[i])
	}

	return nil
}
