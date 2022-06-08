# kaiji

Utility command for cheating.

ざわ…ざわ……

```
　　＿＿　　＿
　 ＞　＼／ /"￣＼
　∠　　　　　 ＜￣
　/　 ／/ 人＼　＞
 /　 /メ＜　ﾚ ヘヽ
｜　ｲ∠二ｿ <∠_| ヽ|
｜/||＼・/〈･／|
｜L|| u￣　 ＼ |
｜ﾋ||　　 ^ _ >|
/　 |、(三三ヲ ﾊ
／| | ＼u ―　/｜
　| |＼ ＼　 /|｜
　レ|　ヽ ＼/ |｜
　  |　 |ニﾆ| |/＼
```

## Usage

```
Write a RANDOM selection of the input lines to standard output.

Usage:
  kaiji [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about kaiji
  keno        Choose one of given arguments
  roulette    Choose one of given arguments

Flags:
  -h, --help   help for kaiji
```

### completion

```
Generate the autocompletion script for kaiji for the specified shell.
See each sub-command's help for details on how to use the generated script.

Usage:
  kaiji completion [command]

Available Commands:
  bash        Generate the autocompletion script for bash
  fish        Generate the autocompletion script for fish
  powershell  Generate the autocompletion script for powershell
  zsh         Generate the autocompletion script for zsh
```

### keno

```
Choose `n` counts of the input lines EXCLUDING `n`th input line to standard output.

Usage:
  kaiji keno [flags]

Flags:
  -n, --count int   Number to select (default 1)
  -h, --help        help for keno
```

### roulette

```
Choose one of the input lines to standard output.
Probability increases as the index grows.

Usage:
  kaiji roulette [flags]
```
