# tomato-clock 🍅

英文版  [README](./doc/README_en.md)

番茄工作法是一种时间管理方法，鼓励人们善于利用时间。使用这种方法，你可以将一天的工作时间分成 25 分钟的时间段，每段工作时间之后休息 5 分钟。这些间隔被称为番茄。”吃”了四个番茄后，休息时间长一些，大约 15 到 20 分钟。

这是一个学习 golang 的项目。这个项目的灵感来自于我正在使用的 [专注清单](https://www.focustodo.cn/?lang=zh_CN) 应用程序。

使用下面的命令启动一个 5 秒的番茄时钟去做一个名为 “learngo” 的任务，并将休息时间设置为 2 秒，设置 5 s 的番茄只是为了更容易地制作 Gif 图演示。

```
tomato-clock -d 5s -bt 2s -t learngo
```

![tomato-clock](./doc/images/tomato-clock.gif)

该功能集是基于 [antonmedv/countdown](https://github.com/antonmedv/countdown) 并受到 [mum4k/termdash](https://github.com/mum4k/termdash) 项目的启发。

小番茄是一个简单的基于终端的应用程序。我只在 macOS 平台上测试成功。它可能在 windows 平台上有 bug。


## Installation

```
go get -u github.com/wowococo/tomato-clock
```

## Usage

如果你已经将 `GOPATH/bin/` 添加到你的环境变量，可以使用 `tomato-clock` 命令。   

```
$ tomato-clock -help

Usage of tomato-clock:
  -bt duration
    	break time duration (default 5m0s)
  -chart
    	show report form, metrics and linechart
  -d duration
    	tomato clock duration (default 25m0s)
  -endtask string
    	mark a task finished
  -t string
    	task name (default "Unnamed")
````

如果没有，可以在 pkg 目录下，执行下面命令

```
cd github.com/wowococo/tomato-clock
go run main.go -d 25m -bt 5m -t learngo
```



For example: 	

开启一个 25 分钟的番茄钟。

```
tomato-clock -d 25m
```

开启一个  45 分钟的番茄钟去做 "learngo" 任务，每个番茄钟之后休息 10 分钟。

```
tomato-clock -d 45m -bt 10m -t learngo
```

标记任务 "learngo" 完成。

```
tomato-clock -endtask learngo
```

展示番茄报表。metrics指标类型包括`总`/`本周`/`今日`专注时间，`总`/`本周`/`今日`完成番茄数专注时间，`总`/`本周`/`今日`完成任务；折线图包括`每月`/`每周`/`每天`的番茄曲线，最近半年`每月`/`每周`/`每天`任务曲线。

	tomato-clock -chart

## Key binding

+ `p` or `P`: 停顿番茄钟倒计时。
+ `c` or `C`: 继续番茄钟倒计时。
+ `Esc` or `Ctrl+C`: 退出倒计时或者退出展示报表。

