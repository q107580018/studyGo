/*
994. 腐烂的橘子
在给定的网格中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。

返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。



示例 1：



输入：[[2,1,1],[1,1,0],[0,1,1]]
输出：4
示例 2：

输入：[[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
示例 3：

输入：[[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。


提示：

1 <= grid.length <= 10
1 <= grid[0].length <= 10
grid[i][j] 仅为 0、1 或 2

*/
package main

import "fmt"

func main() {
	m := [][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}}
	fmt.Println(orangesRotting(m))
}
func orangesRotting(grid [][]int) int {
	bads := make(map[[2]int]bool)
	goods := make(map[[2]int]bool)
	for x, values := range grid {
		for y, value := range values {
			if value == 1 {
				goods[[2]int{x, y}] = true
				continue
			}
			if value == 2 {
				bads[[2]int{x, y}] = true
				continue
			}
		}
	}
	fmt.Println("获取好的路径为", goods)
	fmt.Println("获取坏的路径为", bads)
	minute := 0
	for {
		buf := map[[2]int]bool{}
		if len(goods) == 0 {
			return minute
		}
		for Bk, _ := range bads {
			Bx := Bk[0]
			By := Bk[1]
			if _, ok := goods[[2]int{Bx - 1, By}]; ok {
				delete(goods, [2]int{Bx - 1, By})
				buf[[2]int{Bx - 1, By}] = true
			}
			if _, ok := goods[[2]int{Bx + 1, By}]; ok {
				delete(goods, [2]int{Bx + 1, By})
				buf[[2]int{Bx + 1, By}] = true
			}
			if _, ok := goods[[2]int{Bx, By - 1}]; ok {
				delete(goods, [2]int{Bx, By - 1})
				buf[[2]int{Bx, By - 1}] = true
			}
			if _, ok := goods[[2]int{Bx, By + 1}]; ok {
				delete(goods, [2]int{Bx, By + 1})
				buf[[2]int{Bx, By + 1}] = true
			}
		}
		fmt.Println("第", minute, "次好变坏的路径为", buf)

		if len(buf) == 0 {
			return -1
		}
		for k, v := range buf {
			bads[k] = v
		}
		fmt.Println("第", minute, "次变化后坏果的坐标为", bads)
		minute++
	}
}
