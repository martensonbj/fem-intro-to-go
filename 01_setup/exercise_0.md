# Exercise 0: Installing Go

## Goals: 
- Install Go on your machine

## Setup 
- Visit [golang.org/dl](https://golang.org/dl)

- If you're installing Go for the first time, start at the section `Installing Go From Scratch`.

- If you already have Go installed, feel free to skip installation or find the section below called `Upgrading To Go 1.13`. 

## Directions

### Installing Go From Scratch

#### Download Go
1. Visit [golang.org/dl](https://golang.org/dl) 
2. Select the download package for your operating system
3. Once the installation package is downloaded, double click on that package to kick off installation. 

#### Verify Installation
1. When the installation wizard is complete, open up a terminal window and verify Go is installed with the following commands: 

```bash
which go
=> /usr/local/go

go version
=> go version go1.13.1 darwi/amd64
```

#### Modify Your PATH 

1. Go checks for a certain environment variable, `GOPATH` when executing a Go program. We need to update our bash_profile to help Go find where our files will live.

2. In a code editor, open up your `.bash_profile` which should be located in your home directory. 

3. Add the following lines to the bottom of the file (or wherever you export your PATH environment variables)

```bash
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

#### Verify These Updates

1. Restart your terminal

2. Run the following commands and verify you see similar output

```bash
echo $PATH
=> ..../usr/local/go/bin: 

echo $GOPATH
=> /Users/brennamartenson/go
```

#### Create a go workspace

We just told Go that whenever we run a Go command in our terminal (ie: `go run main.go`), it should look for those files in the GOPATH, which typically defaults to `/Users/yourusername/go`. 

1. Navigate to your home directory (for me that is `/Users/brennamartenson`), and create a `go` directory with `mkdir go`.

2. Change into the go directory with `cd go`, then `mkdir src`, and change into the directory with `cd src`.

3. Clone the repo for this course into the `go` directory.

`git clone https://github.com/martensonbj/fem-intro-to-go`

#### Verify Everything Is Good To GO 😉

1. Change into the `go/fem-intro-to-go` directory

2. Run `go run main.go`

3. If all went well, you should see `"Hello Front End Masters!"`


### Upgrading to Go 1.13

1. Verify you do have Go installed in the first place:

```bash
go version 
=> go version go1.12.1 darwin/amd64
```

2. Uninstall Go

```bash
sudo rm -rf /usr/local/go
```

3. Verify Go was uninstalled

```bash
go version
=> -bash: go: command not found
```

4. Visit [golang.org/dl](https://golang.org/dl) to install the latest Go.


