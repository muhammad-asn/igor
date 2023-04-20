## Igor (Identicon Generator)
Igor is a command-line tool for generating identicon images. It is written in Go and uses the cameron package to generate the identicon images.

Installation
1. Clone the repository: git clone https://github.com/muhammad-asn/igor.git
2. Change directory into the cloned repository: `cd igor`
3. Run make to compile the binary for your current operating system.

### Usage
To generate identicon images with Igor, run the igor binary followed by the desired options. By default, Igor generates one image with a random seed value.

#### Install
```bash
$ go install
```

### Examples
#### Generate one identicon image (by default without flags):
```bash
$ igor
```

#### Generate five identicon images:
```bash
$ igor --num-images 5
```

### Show Help
```bash
$ igor --help
NAME:
   identicon-generator - generate identicon images

USAGE:
   igor-darwin [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --num-images value, -n value  number of images to generate (default: 1)
   --help, -h                    show help
```

#### Cleaning up
To clean up the compiled binaries, run make clean.

```bash
$ make clean
```