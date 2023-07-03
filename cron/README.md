## Cron Parser

This module allows parsing of cron expression and prints a formatted output of possible values

### Language
Golang 1.18

### Supported values
It supports following parts in expression
* Minute
* Hour
* Day of Month
* Month
* Day of week

Allowed characters for each are "/", ",", "*" and "-"

Note: It doesn't support values like Sunday or @Yearly and can only support numbers

## Running the program
The binary file for mac os has been put in bin/ directory. To run use below command
```shell
./bin/cron_parser "*/15 0 1,15 * 1-5 /usr/bin/find"
```
For Linux OS and target platform amd64, a similar file exists with name `cron_parser_linux_amd64`.

If need be, to try out generating the files for other platforms. Please refer to - https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
However, you would have to install Go in your system to build the code.


