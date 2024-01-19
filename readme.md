#  cli tool for start shockwaves

Point to start global infrastructure of Solenopsys network.

### About Solenopsys project:
https://solenopsys.org


### You may use it for: 
- start own Converged cluster  
- get source code for development
- public your projects in Solenopsys network
- create own application for Converged

## Get Started

### Compile for Windows

`GOOS=windows GOARCH=amd64 go build -o xs.exe ./cmd`

### Compile for Linux

`GOOS=linux GOARCH=amd64 go build -o xs ./cmd`

### Get source code for development

`xs dev init front` - get frontends monorepo

`xs dev init back` - get backend monorepo

## Functions

### chart

Helm charts manipulation functions

`xs chart [command]`

- **install** - Install chart
- **list** - List chart
- **remove** - Module chart

### cluster

Cluster manipulation functions

`xs cluster [command]`

- **status** - Cluster status

### dev

Developer functions

`xs dev [command]`

- **init**       - Init monorepo
- **install**    - Install all necessary programs (git,nx,npm,go,...)
- **status**     - Show status of installed env programs (git,nx,npm,go,...)
- **sync**       - Sync modules by configuration


### key

Keys manipulation functions

`xs key [command]`

- **key**        - Gen key
- **pub**        - Generate public key
- **seed**       - Generate seed

### net

Solenopsys network information

`xs net [command]`

- **list**       - List nodes of start network

### node

Node control functions

`xs node [command]`

- **install**  - Install node
- **remove**     - Remove node
- **status**    - Status of node

### public

Public content in ipfs

`xs public [command]`

- **dir**       - Public dir in ipfs
- **file**     - Public file in ipfs

### help
Help about any command


