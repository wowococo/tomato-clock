# tomato-clock 🍅
Pomodoro Technique, this is a project to learn Golang for me. The project is inspired by 'Focus To-Do' application that I am using. 

![tomato-clock](./doc/images/tomato-clock.gif)

The feature set is based on [antonmedv/countdown](https://github.com/antonmedv/countdown) and inspired by the [mum4k/termdash](https://github.com/mum4k/termdash) project.

tomato-clock is a simple terminal based app.I only successfully test it on macOS platform. It maybe has bug on windows platform.


## Installation

```
go get -u github.com/wowococo/tomato-clock
```

## Usage

use tomato-clock command if  you add  `GOPATH/bin/` to your PATH. 

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

	# start a 25 minutes tomato clock
	tomato-clock -d 25m
	
	# start a 45 minutes tomato clock to do the task named "learngo", and set break time 10 minutes
	# after the tomato clock.
	tomato-clock -d 45m -bt 10m -t learngo
	
	# mark the task "learngo" finished
	tomato-clock -endtask learngo
	
	# show the tomato report, include metrics and linechart
	tomato-clock -chart

## Key binding

+ `p` or `P`: To pause the tomato-clock countdown.
+ `c` or `C`: To resume the tomato-clock countdown.
+ `Esc` or `Ctrl+C`: To quit the tomato-clock when counting down or showing chart.

