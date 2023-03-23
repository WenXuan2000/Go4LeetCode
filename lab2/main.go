package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Node struct {
	x int
	y int
}

func Distance(p1 Node, p2 Node) float64 {
	// 这里只做比较，于是对于两点之间的距离没有开根号
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}
func BFFindClosestPair(nodeList []Node) (ans float64, p1 Node, p2 Node) {
	/*
		计算出二维平面内所有点对（任意两点的组合）之间的距离
		逐一比较这些距离的大小，找出其中的最短的距离，则这两点即为所求的“最近点对”
	*/
	ans = math.MaxFloat64
	for i := 0; i < len(nodeList); i++ {
		for j := i + 1; j < len(nodeList); j++ {
			dis := Distance(nodeList[i], nodeList[j])
			if ans > dis {
				ans = dis
				p1 = nodeList[i]
				p2 = nodeList[j]
			}
		}
	}
	return ans, p1, p2
}

func FindClosestPair(nodeList []Node) (float64, Node, Node) {
	var ans float64
	var p1 Node
	var p2 Node
	left := 0
	right := len(nodeList) - 1
	mid := (left + right) / 2
	if (right - left) == 0 {
		return math.MaxFloat64, nodeList[left], nodeList[right]
	}
	if (right - left) == 1 {
		return Distance(nodeList[left], nodeList[right]), nodeList[left], nodeList[right]
	}
	if (right - left) == 2 {
		d1 := Distance(nodeList[left], nodeList[right])
		d2 := Distance(nodeList[left], nodeList[left+1])
		d3 := Distance(nodeList[left+1], nodeList[right])
		if d1 < d2 && d1 < d3 {
			return d1, nodeList[left], nodeList[right]
		} else if d2 < d1 && d2 < d3 {
			return d2, nodeList[left], nodeList[left+1]
		} else if d3 < d1 && d3 < d2 {
			return d3, nodeList[left+1], nodeList[right]
		} else if d1 == d2 && d1 == d3 {
			return d1, nodeList[left], nodeList[right]
		}
	}
	leftans, leftp1, leftp2 := FindClosestPair(nodeList[:mid+1])
	rightans, rightp1, rightp2 := FindClosestPair(nodeList[mid+1:])
	if rightans < leftans {
		ans = rightans
		p1 = rightp1
		p2 = rightp2
	} else {
		ans = leftans
		p1 = leftp1
		p2 = leftp2
	}
	leftlist := []Node{}
	rightlist := []Node{}
	for i := 0; i < len(nodeList); i++ {
		if nodeList[i].x-nodeList[mid].x <= 0 && float64(nodeList[i].x-nodeList[mid].x) > -ans {
			leftlist = append(leftlist, nodeList[i])
		}
		if nodeList[i].x-nodeList[mid].x > 0 && float64(nodeList[i].x-nodeList[mid].x) < ans {
			rightlist = append(rightlist, nodeList[i])
		}
	}
	//排序
	sort.Slice(rightlist, func(i, j int) bool {
		if rightlist[i].y < rightlist[j].y {
			return true
		}
		return false
	})
	for i := 0; i < len(leftlist); i++ {
		index := 0
		for index = 0; index < len(rightlist) && leftlist[i].y < rightlist[index].y; index++ {
		}
		if index != 0 {
			index--
		}
		for j := 0; j < 7 && index+j < len(rightlist); j++ {
			dis := Distance(leftlist[i], rightlist[j+index])
			if dis < ans {
				ans = dis
				p1 = leftlist[i]
				p2 = rightlist[j+index]
			}
		}
	}
	return ans, p1, p2
}
func Printres(nodeList []Node) {
	t1 := time.Now().UnixNano()
	bftime, bfp1, bfp2 := BFFindClosestPair(nodeList)
	elapsed := time.Now().UnixNano() - t1
	fmt.Printf("最短距离：%.2f 最短距离点对：%v和%v \n", bftime, bfp1, bfp2)
	fmt.Printf("蛮力法所需时间(ns)：%v \n", elapsed)

	t1 = time.Now().UnixNano()
	sort.Slice(nodeList, func(i, j int) bool {
		if nodeList[i].x < nodeList[j].x {
			return true
		}
		return false
	})
	ctime, p1, p2 := FindClosestPair(nodeList)
	elapsed = time.Now().UnixNano() - t1
	fmt.Printf("最短距离：%.2f 最短距离点对：%v和%v \n", ctime, p1, p2)
	fmt.Printf("分治法所需时间(ns)：%v \n", elapsed)
}
func GenNode(nodeNumber int) []Node {
	nodeList := []Node{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < nodeNumber; i++ {
		myNode := Node{
			x: rand.Intn(10000000),
			y: rand.Intn(10000000),
		}
		nodeList = append(nodeList, myNode)
	}
	return nodeList
}
func main() {
	// 我们一般使用系统时间的不确定性来进行初始化
	nodeNumber := []int{1000, 2000, 3000, 4000, 5000}
	for _, v := range nodeNumber {
		Printres(GenNode(v))
	}
}
