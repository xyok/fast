# fast cli


> 从模版生成Go 代码
>


### 项目框架

- [x] gin sample
- [ ] fiber

可以是任意项目模版结构, 放在执行目录的tpl下即可


### 参数说明

```
NAME:
   fast cli - Create proj scaffold from current framework

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name value, -n value    project name
   --tpl value               project tamplate tpl/[*] (default: "gin")
   --output value, -o value  generate out
   --remote value            git remote
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)
```

### 创建工程

```bash
go run main.go gen --name apisvr --frame gin --output /home/user/workspace
```
