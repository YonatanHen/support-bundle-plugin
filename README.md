# support-bundle upload

## About this plugin
This JFrog CLI plugin let JFrog customers upload [Support Bundles](https://jfrog.com/help/r/jfrog-platform-administration-documentation/support-zone) and files to the JFrog Support SaaS instance.

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
  $ ./upload-support-bundle u 1 ./test2.txt ./test3.txt ./test4.txt 
  
    15:06:05 [🔵Info] Uploaded file: ./test2.txt.
    15:06:06 [🔵Info] Uploaded file: ./test3.txt.
    15:06:06 [🔵Info] Uploaded file: ./test4.txt.
  ```

    <img width="1137" alt="image" src="https://github.com/YonatanHen/upload-support-bundle-plugin/assets/57364867/9a74f3b8-6d12-4c68-bb89-8ffde0270749">

## Release Notes
The release notes are available [here](RELEASE.md).
