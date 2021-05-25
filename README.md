# piihunter
piihunter is a sensitive unencrypted data detection tool built to scan source code repositories for passwords, token, weak cryptography usage, private key voilations and lot more. It can also scan your codebase for unencrypted PII such as emails, phone numbers, addresses, zip codes, credit/debit card numbers etc. 

It can be used to scan remote git repositories or local directories. 

## Installation
If you already have Go installed and $GOPATH set, then simply clone this repository and run `make install`. It will install executable to your `$GOPATH/bin` directory. 

## Usage
To launch a scan against directory: 
```
$ piihunter --dir=/path/to/directory/
```
To launch a scan against git repository: 
```
$ piihunter --git=https://github.com/[USERNAME]/[REPOSITORY]
```
Writing scan result to a file: 
```
$ piihunter --dir=/path/to/directory/ --out=/path/to/file.json
```
## Building piihunter
Assuming you have Go installed on your machine
1. Clone the repository
2. `cd piihunter`
3. `go build` 

## Running tests
In the project base directory, run `make test`. Alternatively you can go to `cmd` directory and run `go test`. 
