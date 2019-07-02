# OSCAP JSON

A simple tool for parsing output from OpenSCAP and converting it to JSON

## Usage
oscap-json can either read input from a file or `stdin`. By default it will try to read from `stdin`. To override this behaviour and read from a file, use the  `-file` flag.

```
# Read from stdin
$ sudo oscap xccdf eval --profile my_profile /path/to/scap.xml | oscap-json

# Read from a file
$ sudo oscap xccdf eval --profile my_profile /path/to/scap.xml > /tmp/report.txt
oscap-json -file /tmp/report.txt
```