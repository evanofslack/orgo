# orgo

simple executable written in Go to help organize your files

### Installation

clone this repo
```bash
git clone https://github.com/evanofslack/orgo
```

navigate into the directory
```bash
cd orgo
```

build the executable

```bash
go build 
```

### Usage

move all files on your desktop ending with extension `.png` and `.jpeg` into a new directory called screenshots:
```bash
orgo $HOME/Desktop screenshots .png .jpg
```

run `orgo --help` for full instructions


### Testing

run all tests with `go test`