# XS - (eXec Solenopsys) - cli tool

Point to start global infrastructure of Solenopsys

# Functions

### chart

Helm charts manipulation functions

`xs chart [command]`

**Subcommands:**

- install     Install chart
- list        List chart
- remove      Module chart

### cluster

Cluster manipulation functions

`xs cluster [command]`

**Subcommands:**

- status      Cluster status

### dev

Developer functions

`xs dev [command]`

**Subcommands:**

- init        Init monorepo
- install     Install all necessary programs (git,nx,npm,go,...)
- status      Show status of installed env programs (git,nx,npm,go,...)
- sync        Sync modules by configuration


### key

Keys manipulation functions

`xs key [command]`

**Subcommands:**

- key         Gen key
- pub         Generate public key
- seed        Generate seed

### net

Solenopsys network information

`xs net [command]`

**Subcommands:**

- list        List nodes of start network

### node

Node control functions

`xs node [command]`

**Subcommands:**

- install     Install node
- remove      Remove node
- status      Status of node

### public

Public content in ipfs

`xs public [command]`

**Subcommands:**

- dir         Public dir in ipfs
- file        Public file in ipfs

### help

Help about any command

## Get Started

## Compile

### for Windows

`GOOS=windows GOARCH=amd64 go build -o xs.exe main.go`

### for Linux

`GOOS=linux GOARCH=amd64 go build -o xs e main.go`

## Get source code for development

`xs dev init front` - get frontends monorepo

`xs dev init back` - get backend monorepo
