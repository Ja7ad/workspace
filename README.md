# workspace
example workspace multiple modules

## init workspace

1. create modules folder and init go.mod

```shell
$ mkdir module1 && go mod init github.com/user/repo/module1
```


```shell
$ mkdir module2 && go mod init github.com/user/repo/module2
```

2. init workspace in root path

```shell
go work init
```

3. add modules to workspace

```shell
$ go work use module1 && go work use module2
```

## Use modules

```shell
$ go get -u github.com/Ja7ad/workspace/add
```

```shell
$ go get -u github.com/Ja7ad/workspace/sum
```
