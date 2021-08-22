# tomato-clock 🍅
The Pomodoro Technique is a time management system that encourages people to work *with* the time they have—rather than against it. Using this method, you break your workday into 25-minute chunks separated by five-minute breaks. These intervals are referred to as pomodoros. After about four pomodoros, you take a longer break of about 15 to 20 minutes.

This is a project to learn Golang for me. The project is inspired by 'Focus To-Do' application that I am using.

Use the following command to start a 5 seconds tomato clock to do the task named "learngo", and set break time 2 seconds after the tomato clock, which is just to make Gif easier. 

```
tomato-clock -d 5s -bt 2s -t learngo
```

![tomato-clock](./images/tomato-clock.gif)

The feature set is based on [antonmedv/countdown](https://github.com/antonmedv/countdown) and inspired by the [mum4k/termdash](https://github.com/mum4k/termdash) project.

tomato-clock is a simple terminal based app.I only successfully test it on macOS platform. It maybe has bug on windows platform.


## Installation

```
go get -u github.com/wowococo/tomato-clock
```

## Usage

Use `tomato-clock` command if  you add  `GOPATH/bin/` to your PATH. 

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
If not, you can go to your pkg directory, then

```
cd github.com/wowococo/tomato-clock
go run main.go -d 25m -bt 5m -t learngo
```



For example: 	

Starts a 25 minutes tomato clock.

```
tomato-clock -d 25m
```

Starts a 45 minutes tomato clock to do the task named "learngo", and set break time 10 minutes after the tomato clock.

```
tomato-clock -d 45m -bt 10m -t learngo
```

Marks the task "learngo" finished.

```
tomato-clock -endtask learngo
```

Shows the tomato report, include metrics and linechart.

	tomato-clock -chart

## Key binding

+ `p` or `P`: To pause the tomato-clock countdown.
+ `c` or `C`: To resume the tomato-clock countdown.
+ `Esc` or `Ctrl+C`: To quit the tomato-clock when counting down or showing chart.

