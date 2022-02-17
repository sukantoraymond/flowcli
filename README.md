# Flow CLI

## User Story

As a Flow software developer, I need a CLI tool to retrieve information regarding test accounts in Flow Emulator for development purposes. A Flow account contains address, balance, public keys and contracts and the CLI will need to query and present all these information to developer.  

Acceptance Criteria: 
- CLI should present all available options to developer when called without any command and flag
- CLI should handle error, present the error message to developer and if possible, present potential remedy to developer
- CLI should suggest the correct command if possible if developer mistyped a command
- CLI should only be able to present account information for valid account address present in the Flow Emulator

## Usage

To use flowcli CLI, this repo needs to be git cloned in $GOPATH/src/github.com/sukantoraymond

Change directory to project directory 

```
cd flowcli
```

Get all dependencies required as specified in go.mod

```
go get
```

Generate the executable binary for the CLI 

```
go install github.com/sukantoraymond/flowcli
```

Start using the CLI to get account information! (Remember to call start the Flow Emulator through docker-compose up prior to calling flowcli)
```
flowcli accounts get <address>
``` 
or simply call flowcli to present command options
```
flowcli
``` 

## Design
  
This CLI is built using [Cobra](https://github.com/spf13/cobra) library, which is the most popular library used for CLI tools. I chose to use this library because it automatically parses user commands and auto suggests the correct command without having to implement additional custom code that will have to be maintained.

The CLI command to get user command is structured as flowcli accounts get, so that in the future additional commands to interact with accounts such as delete and update can be integrated more easily. 

In addition, the CLI uses the existing [Flow Go SDK](https://github.com/onflow/flow-go-sdk/blob/master/examples/get_accounts/main.go) to query account information to minimize the number of custom code that will have to be maintained in the future.



