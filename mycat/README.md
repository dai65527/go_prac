# mycat
A simple cat command by Go.

# compile

```
go build mycat.go
```

# usage
With argument.

```
$ ./mycat file1 file2 file3
content of file1
file1 file1 file1
content of file2
file1 file1 file1
content of file3
file3 file3 file3
```

Without arguments (read from stdin).

```
$ ./mycat
hoge (input and enter)
hoge
fuga (input and enter)
fuga
(input and enter)

(End with Ctrl-D)
```

With `-n` option (display line number)

```
     1 content of file1
     2 file1 file1 file1
     3 content of file2
     4 file1 file1 file1
     5 content of file3
     6 file3 file3 file3
```
