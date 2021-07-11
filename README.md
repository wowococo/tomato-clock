# tomato-clock 🍅
Pomodoro Technique, this is a project to learn Golang for me. The project is inspired by 'Focus To-Do' application that I am using. The feature set is based on [antonmedv/countdown](https://github.com/antonmedv/countdown) and inspired by the [mum4k/termdash](https://github.com/mum4k/termdash) project.

tomato-clock is a cross-platform terminal based app.

## Installation

```
go get github.com/wowococo/tomato-clock
```

## Usage
the second param use the time.Duration format, for example: 

	# start a 25 seconds tomato clock
	tomato 25s  

	# start a 25 minutes tomato clock
	tomato 25m

	# start a 1 hour 20 minutes 20 seconds tomato clock
	tomato 1h20m20s

OR:
```
go run main.go 25m
```

