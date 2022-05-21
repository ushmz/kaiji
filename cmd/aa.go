package cmd

import (
	"github.com/spf13/cobra"
)

var aa = `
ざわ‥ ざわ‥
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
	`

func printAA(cmd *cobra.Command) {
	cmd.Println(aa)
}
