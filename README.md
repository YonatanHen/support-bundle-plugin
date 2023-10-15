# support-bundle upload

## About this plugin
This JFrog CLI plugin allows JFrog customers to easily upload [Support Bundles](https://jfrog.com/help/r/jfrog-platform-administration-documentation/support-zone) and files to the [JFrog Support SaaS instance](https://supportlogs.jfrog.com) using CLI command instead of cURL. 
This plugin also enables the uploading of multiple files using just one command.

## Installation with JFrog CLI
Installing the latest version:

`$ jf plugin install support-bundle`

<!---Installing a specific version:`--->

<!--- j`$ jf plugin install hello-frog@version` --->

Uninstalling a plugin

`$ jf plugin uninstall support-bundle`

## Usage
### Commands
* upload
    - Arguments:
        - ticket number - The support ticket number in JFrog Portal.
        - files - paths to the files to upload.
    - Example:
      ```
      $ jf support-bundle u {ticket-number} {files}
      ```
  
        ```
      $ jf support-bundle u 1 ./test2.txt ./test3.txt ./test4.txt 
      
        15:06:05 [🔵Info] Uploaded file: ./test2.txt.
        15:06:06 [🔵Info] Uploaded file: ./test3.txt.
        15:06:06 [🔵Info] Uploaded file: ./test4.txt.
      ```
        
## Release Notes
The release notes are available [here](RELEASE.md).
