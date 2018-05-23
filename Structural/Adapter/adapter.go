/*
Created on 2018/5/23 14:06

author: ChenJinLong

Content:
适配器将一个接口转换成客户希望的另一个接口，使接口不兼容的那些类可以一起工作，其别名为包装器(Wrapper)
*/
package adapter


//适配者
type QuickSort struct {

}

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}
func (q *QuickSort) quickSort(arr []int) {
	start :=0
	end := len(arr) -1
	quickSort(arr,start,end)
}

//适配者
type BinarySearch struct {

}

func (b *BinarySearch) binarySearch(arr []int,key int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left+right) / 2
		midVal := arr[mid]
		if (midVal <key){
			left  = mid + 1
		}else if (midVal >key) {
			right = mid - 1
		}else {
			return 1
		}

	}
	return -1
}


//抽象成绩操作类：目标接口
type ScoreOperation interface {
	sort(arr []int)
	search(arr []int,key int) int
}

//操作适配器：适配器
type OperationAdapter struct {
	quickSort *QuickSort
	binarySearch *BinarySearch
}

func (o *OperationAdapter) sort(arr []int) {
	o.quickSort.quickSort(arr)
}

func (o *OperationAdapter)search(arr []int,key int)int {
	return o.binarySearch.binarySearch(arr,key)
}

