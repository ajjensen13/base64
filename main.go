package main

import (
	"encoding/base64"
	"flag"
	"io"
	"os"
)

func main() {
	d := flag.Bool("d", false, "decode (default is to encode)")
	i := flag.String("in", "", "input file (default is to read from STDIN)")
	o := flag.String("out", "", "output file (default is to write to STDOUT)")
	flag.Parse()

	rc, err := in(*i)
	if err != nil {
		exitOnError(err)
	}
	defer rc.Close()

	wc, err := out(*o)
	if err != nil {
		exitOnError(err)
	}
	defer wc.Close()

	switch *d {
	case true:
		err = decode(wc, rc)
		if err != nil {
			exitOnError(err)
		}
	default:
		err = encode(wc, rc)
		if err != nil {
			exitOnError(err)
		}
	}
}

func exitOnError(err error) {
	_, _ = os.Stderr.WriteString(err.Error())
	os.Exit(2)
}

func in(fpath string) (io.ReadCloser, error) {
	if fpath == "" {
		return os.Stdin, nil
	}

	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func out(fpath string) (io.WriteCloser, error) {
	if fpath == "" {
		return os.Stdout, nil
	}

	f, err := os.Create(fpath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func encode(des io.Writer, src io.Reader) error {
	ec := base64.NewEncoder(base64.StdEncoding, des)
	defer ec.Close()
	_, err := io.Copy(ec, src)
	return err
}

func decode(des io.Writer, src io.Reader) error {
	dc := base64.NewDecoder(base64.StdEncoding, src)
	_, err := io.Copy(des, dc)
	return err
}
