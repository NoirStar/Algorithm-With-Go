package dataStruct

import "fmt"

type Heap struct {
	list []int
}

// Max Heap
func (h *Heap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] > h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	if len(h.list) == 1 {
		return h.list[0]
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	// 맨마지막 있는애를 빼서 리스트를 새로 만든다.
	// 맨마지막걸 맨앞에 넣고, 비교
	h.list[0] = last
	idx := 0
	// 인덱스가 길이보다 클때까지 계속 돌림(자식노드 다돌기)
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx*2 + 1
		rightIdx := idx*2 + 2

		// left 인덱스가 길이보다 클때, 자식노드가 없다는뜻.
		if leftIdx >= len(h.list) {
			break
		}
		// 왼쪽자식 노드 값이 부모보다 크면. 바꿔야됨 바로바꾸지 않음.
		if h.list[leftIdx] > h.list[idx] {
			// 일단 바꿔야될 인덱스를 저장해둠
			swapIdx = leftIdx
		}

		// 오른쪽 인덱스가 리스트 길이보다 작을때, 즉 오른쪽인덱스 있을때
		if rightIdx < len(h.list) {
			// 오른쪽 인덱스값이 부모보다 큰경우. 바꿔야되는데,
			if h.list[rightIdx] > h.list[idx] {
				// 스왑인덱스를이용 왼쪽오른쪽 값비교
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}

		// 바꿀애가 없음
		if swapIdx < 0 {
			break
		}
		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		// 인덱스를 배꾼 인덱스로 갱신해서 포문다시돔
		idx = swapIdx
	}
	return top
}

func (h *Heap) Count() int {
	return len(h.list)
}

func (h *Heap) Print() {
	fmt.Println(h.list)
}
