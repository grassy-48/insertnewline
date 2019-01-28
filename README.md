# What's INSERTNEWLINE?
This tool is to add insert new line free position.

# Usage
* run command
```
# windows
> ./insernewline_win.exe -f {YOUR_TARGETFILE} -n {INSERT_POSITION_NUM} -d {HANGING_FLAG}

# Mac
> ./insernewline -f {YOUR_TARGETFILE} -n {INSERT_POSITION_NUM} -d {HANGING_FLAG}
```

* sample
```
# windows
> ./insernewline_win.exe -f ./body.txt -n 10 -d

# Mac
> ./insernewline -f ./body.txt -n 10 -d
```

* Get Output file!
`./dist/body.txt`

## OPTION

```
NAME:
   insertnewline - insert new line free position your text

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, -f value  source test file.
   --num value, -n value   insert new line position num(1>n) (default: 0)
   --down, -d              use string down
   --help, -h              show help
   --version, -v           print the version
```
