package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][] string{
		[]string{"Alfred", "15", "10/20", "(10.32, 56.21, 30.25)"},
		[]string{"Beelzebub", "30", "30/50", "(1, 1, 1)"},
		[]string{"Hortense", "21", "80/80", "(1, 1, 1)"},
		[]string{"Pokey", "8", "30/40", "(1, 1, 1)"},
	}
	
	table := tablewriter.NewWriter(os.Stdout) // 定义 table
	table.SetHeader([]string{"NPC", "Speed", "Power", "Locatuib"}) // 设置头部
	table.AppendBulk(data) // 将数组写到 table 里
	table.Render() // table 输出
}