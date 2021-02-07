package main

import (
	"fmt"
	"math/rand"
	"time"
)

// probability X
const Prob = 0.5

type BoxMap struct {
	// 地图中的箱子 
	Boxes []Box
	// 最后一个箱子
	LastBox Box
	// 最后一个点
	LastPoint int
	// 访问过的箱子
	visited []string
	// 所有点
	allPoint []int
	// 访问过的点
	visitedPoint []int
}

// 检查箱子的点是否都访问过
func (b BoxMap) IsDone() bool {
	if len(b.visited) > len(b.Boxes) {
		return true
	}
	return false
}

func (b *BoxMap) Run() {
	var lastFound bool
	rand.Seed(time.Now().UnixNano())
	for {
		if b.IsDone() {
			break
		}
		currentBox := b.GetCurrentRoundBox()
		if currentBox.CheckIsDone(b.visitedPoint) {
			continue
		}
		first := true
		for {
			if currentBox.CheckIsDone(b.visitedPoint) {
				break
			}
			if !first && !lastFound {
				currentBox.GetOne(b)
				lastFound = true
			} else {
				prob := rand.Float64()
				if prob < Prob {
					currentBox.GetOne(b)
					lastFound = true
				} else {
					point := b.GetOne()
					if currentBox.IsIn(point) {
						lastFound = true
					} else {
						lastFound = false
					}
				}
			}
			first = false
		}
		if b.IsDone() {
			break
		}
	}
}

// 随机从所有未访问的点中选一个点(除了最后一个点)
func (b *BoxMap) GetOne() int {
	var points []int
	for _, p := range b.allPoint {
		found := false
		for _, item := range b.visitedPoint {
			if p == item {
				found = true
				break
			}
		}
		if found {
			continue
		}
		if p == b.LastPoint {
			continue
		}
		points = append(points, p)
	}
	if len(points) == 0 {
		points = append(points, b.LastPoint)
	}

	point := points[rand.Intn(len(points))]
	fmt.Printf("visit point, point: %d\n", point)
	b.visitedPoint = append(b.visitedPoint, point)
	return point
}

// 随机获取一个未集全所有点的箱子
func (b *BoxMap) GetCurrentRoundBox() Box {
	var result []Box
	for _, item := range b.Boxes {
		found := false
		for _, name := range b.visited {
			if item.name == name {
				found = true
				break
			}
		}
		if found {
			continue
		}
		result = append(result, item)
	}
	if len(result) == 0 {
		result = append(result, b.LastBox)
	}
	box := result[rand.Intn(len(result))]
	b.visited = append(b.visited, box.name)
	fmt.Printf("visit box, points: %v\n", box.point)
	return box
}

type Box struct {
	name  string
	point []int
}

// 检查箱子的点是否都被访问过
func (box *Box) CheckIsDone(visitedPoint []int) bool {
	for _, p := range box.point {
		found := false
		for _, item := range visitedPoint {
			if p == item {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// 随机从箱子中未访问过的点中获取一个点
func (box *Box) GetOne(boxMap *BoxMap) int {
	var points []int
	for _, p := range box.point {
		found := false
		for _, item := range boxMap.visitedPoint {
			if p == item {
				found = true
				break
			}
		}
		if found {
			continue
		}
		if p == boxMap.LastPoint {
			continue
		}
		points = append(points, p)
	}
	if len(points) == 0 {
		points = append(points, boxMap.LastPoint)
	}
	point := points[rand.Intn(len(points))]
	fmt.Printf("visit point,point: %d\n", point)
	boxMap.visitedPoint = append(boxMap.visitedPoint, point)
	return point
}

func (box *Box) IsIn(p int) bool {
	for _, item := range box.point {
		if p == item {
			return true
		}
	}
	return false
}

func main() {
	m := BoxMap{
		Boxes: []Box{
			{name: "A", point: []int{1, 2}},
			{name: "B", point: []int{2, 3, 8, 9}},
			{name: "C", point: []int{4, 5}},
			{name: "D", point: []int{7, 13}},
			{name: "E", point: []int{14, 15}},
		},
		allPoint:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
		LastPoint: 18,
		LastBox:   Box{name: "F", point: []int{11, 12, 17, 18}},
	}

	m.Run()
}
