# file_size

## Description  
Search for files with the specified capacity (MB). If the size of the file exceeds the specified capacity, it is output as a full path to the stdout.  

## USAGE
Build.  
```
$ go build file_size.go
```
The first argument is the file size (MB), and the second argument is the root of the directory.  
```
$ file_size.exe <FILESIZE(MB)> <DIR>
```

## REQUIRES
- Windows

## License
MIT
