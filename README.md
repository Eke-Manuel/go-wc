# go-wc
Implementation of the Unix wc command in Go


### If you have Go installed and set up:

Run: 
```bash
go install
```

Usage:
```
gwc [global options] [FILE]
```
Sample Output:
```
3625 test.txt
```

For a list of available options, run:
```bash
gwc help
```

### If you don't have Go installed:

Usage:
```
gwc.exe [global options] [FILE]
```

For a list of available options, run:
```bash
gwc.exe help
```

Support for default action to be added soon, use flags to indicate option. Run 'gwc help' for a list of available flags. 
Compound flags are not supported yet.