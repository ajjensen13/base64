# Package base64
Implements a cross-platform base64 encoder/decoder as specified by RFC 4648

## Usage

```bash
> echo "test" | base64 -out ./output.txt | base64 -in ./output.txt -d
test
```

```bash
> base64 -h

Usage of base64:
  -d    decode (default is to encode)

  -in string
        input file (default is to read from STDIN)

  -out string
        output file (default is to write to STDOUT)
```