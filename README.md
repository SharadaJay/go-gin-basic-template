# go-gin-basic-template
A basic REST API template for Go with Gin Web Framework

* Should provide following environment variable. (The value is the name of the property file without the extension)
    * ENV
        * In Windows PowerShell
            ```bash
              $Env:ENV = 'application'
            ```
        * In Linux
            ```bash
              export ENV=application
            ```

    <br />

* Install dependencies
    ```bash
        go mod tidy
    ```


* Run application
    ```bash
        go run main.go
    ```