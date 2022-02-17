# Flow CLI

## User Story

As a Flow software developer, I need a CLI tool to retrieve information regarding several test accounts in Flow Emulator for development purposes. A Flow account contains address, balance, public keys and contracts and the CLI will need to query and present all these information to developer.  

Acceptance Criteria: 
- CLI should present all available options to developer when called without any command and flag
- CLI should handle error, present the error message to developer and if possible, present potential remedy to developer
- CLI should suggest the correct command if possible if developer mistyped a command
- CLI should only be able to present account information for valid account address present in the Flow Emulator

## Usage

To use flowcli CLI, this repo needs to be git cloned in $GOPATH/github.com/sukantoraymond. 

Change directory to project directory 

cd flowcli

Get all dependencies required as specified in go.mod

go get

Start using the CLI! 

flowcli <commands> 

## Design
  
This CLI is using Cobra 


