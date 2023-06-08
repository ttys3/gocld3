# cld3

[![GoDoc](https://godoc.org/github.com/ttys3/gocld3/cld3?status.svg)](https://godoc.org/github.com/ttys3/gocld3)

Package cld3 implements language detection using the Compact Language Detector v3.

This package includes the relevant sources from the cld3 project, so it doesn't
require any external dependencies. For more information on CLD3, see [https://github.com/google/cld3/](https://github.com/google/cld3/).

Install with `go get github.com/ttys3/gocld3/cld3`.

Documentation is available on [GoDoc](https://godoc.org/github.com/ttys3/gocld3/cld3).

### Example

```go
	cld := cld3.NewDefault()
	defer cld.Free()

	res := cld.FindLanguage("Hey, this is an english sentence")
	if res.IsReliable {
		fmt.Println("pretty sure we've got text written in", res.Language)
	}
	res = cld.FindLanguage("Muy bien, gracias.")
	if res.IsReliable {
		fmt.Println("ah, and this one is", res.Language)
	}
  ```

### Build

you need `protobuf-devel` which provides

libprotobuf, libprotoc, /usr/lib64/pkgconfig/protobuf.pc, /usr/include/google/protobuf header files


```shell
sudo dnf install -y protobuf-devel
```

you also need `protobuf-compiler` if you need to regenerate the pb files

for Fedora 38, the repo package is too old, you need a new one:

```shell
cd /tmp/

curl -LZO https://github.com/protocolbuffers/protobuf/releases/download/v23.2/protoc-23.2-linux-x86_64.zip
unzip protoc-23.2-linux-x86_64.zip
sudo cp -v -R include/google/protobuf/* /usr/include/google/protobuf 
sudo cp -b bin/protoc /usr/bin/protoc
```

### Credits

this project is based on https://github.com/jmhodges/gocld3

thanks for jmhodges's work.

