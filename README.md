# OSCAP JSON

A simple tool for parsing output from OpenSCAP and converting it to JSON

## Usage
oscap-json is meant to have the output from an oscap scan piped into it. It reads the input from `stdin`, processes it and outputs as JSON.

```
$ sudo oscap xccdf eval --profile my_profile /path/to/scap.xml | oscap-json
```