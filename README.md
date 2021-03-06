# tomato-clock ð

è±æç  [README](./doc/README_en.md)

çªèå·¥ä½æ³æ¯ä¸ç§æ¶é´ç®¡çæ¹æ³ï¼é¼å±äººä»¬åäºå©ç¨æ¶é´ãä½¿ç¨è¿ç§æ¹æ³ï¼ä½ å¯ä»¥å°ä¸å¤©çå·¥ä½æ¶é´åæ 25 åéçæ¶é´æ®µï¼æ¯æ®µå·¥ä½æ¶é´ä¹åä¼æ¯ 5 åéãè¿äºé´éè¢«ç§°ä¸ºçªèãâåâäºåä¸ªçªèåï¼ä¼æ¯æ¶é´é¿ä¸äºï¼å¤§çº¦ 15 å° 20 åéã

è¿æ¯ä¸ä¸ªå­¦ä¹  golang çé¡¹ç®ãè¿ä¸ªé¡¹ç®ççµææ¥èªäºææ­£å¨ä½¿ç¨ç [ä¸æ³¨æ¸å](https://www.focustodo.cn/?lang=zh_CN) åºç¨ç¨åºã

ä½¿ç¨ä¸é¢çå½ä»¤å¯å¨ä¸ä¸ª 5 ç§ççªèæ¶éå»åä¸ä¸ªåä¸º âlearngoâ çä»»å¡ï¼å¹¶å°ä¼æ¯æ¶é´è®¾ç½®ä¸º 2 ç§ï¼è®¾ç½® 5 s ççªèåªæ¯ä¸ºäºæ´å®¹æå°å¶ä½ Gif å¾æ¼ç¤ºã

```
tomato-clock -d 5s -bt 2s -t learngo
```

![tomato-clock](./doc/images/tomato-clock.gif)

è¯¥åè½éæ¯åºäº [antonmedv/countdown](https://github.com/antonmedv/countdown) å¹¶åå° [mum4k/termdash](https://github.com/mum4k/termdash) é¡¹ç®çå¯åã

å°çªèæ¯ä¸ä¸ªç®åçåºäºç»ç«¯çåºç¨ç¨åºãæåªå¨ macOS å¹³å°ä¸æµè¯æåãå®å¯è½å¨ windows å¹³å°ä¸æ bugã


## Installation

```
go get -u github.com/wowococo/tomato-clock
```

## Usage

å¦æä½ å·²ç»å° `GOPATH/bin/` æ·»å å°ä½ çç¯å¢åéï¼å¯ä»¥ä½¿ç¨ `tomato-clock` å½ä»¤ã   

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

å¦ææ²¡æï¼å¯ä»¥å¨ pkg ç®å½ä¸ï¼æ§è¡ä¸é¢å½ä»¤

```
cd github.com/wowococo/tomato-clock
go run main.go -d 25m -bt 5m -t learngo
```



For example: 	

å¼å¯ä¸ä¸ª 25 åéççªèéã

```
tomato-clock -d 25m
```

å¼å¯ä¸ä¸ª  45 åéççªèéå»å "learngo" ä»»å¡ï¼æ¯ä¸ªçªèéä¹åä¼æ¯ 10 åéã

```
tomato-clock -d 45m -bt 10m -t learngo
```

æ è®°ä»»å¡ "learngo" å®æã

```
tomato-clock -endtask learngo
```

å±ç¤ºçªèæ¥è¡¨ãmetricsææ ç±»ååæ¬`æ»`/`æ¬å¨`/`ä»æ¥`ä¸æ³¨æ¶é´ï¼`æ»`/`æ¬å¨`/`ä»æ¥`å®æçªèæ°ä¸æ³¨æ¶é´ï¼`æ»`/`æ¬å¨`/`ä»æ¥`å®æä»»å¡ï¼æçº¿å¾åæ¬`æ¯æ`/`æ¯å¨`/`æ¯å¤©`ççªèæ²çº¿ï¼æè¿åå¹´`æ¯æ`/`æ¯å¨`/`æ¯å¤©`ä»»å¡æ²çº¿ã

	tomato-clock -chart

## Key binding

+ `p` or `P`: åé¡¿çªèéåè®¡æ¶ã
+ `c` or `C`: ç»§ç»­çªèéåè®¡æ¶ã
+ `Esc` or `Ctrl+C`: éåºåè®¡æ¶æèéåºå±ç¤ºæ¥è¡¨ã

