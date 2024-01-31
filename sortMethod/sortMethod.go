package sortMethod

import (
    "fmt"
)

//練習常見排序方法

func init(){
    BubbleSort()
    SelectionSort()
    InsertionSort()
    MergeSort()
    QuickSort()
}


// 氣泡排序
func BubbleSort(){

    nums := []int{5,3,8,1,13,4,9,2}

    for i:= 0 ; i < len(nums) ; i++ {
        for k := 0 ; k < len(nums) - 1 ; k++  {
            var tmp int

            if nums[k] > nums[k+1] {
                tmp = nums[k]
                nums[k] = nums[k+1]
                nums[k+1] = tmp
            }
        }
    }

    fmt.Println("氣泡排序 => ", nums)
}

// 選擇排序法
func SelectionSort(){

    nums := []int{5,3,8,1,13,4,9,2}

    for i:=0 ; i < len(nums) ; i++ {
        minKey := i
        for k := i ; k < len(nums) ; k++ {
            if nums[k] < nums[minKey] {
                minKey = k
            }
        }

        cn := nums[minKey]
        nums[minKey] = nums[i]
        nums[i] = cn
    }

    fmt.Println("選擇排序 => ", nums)
}


//插入排序
func InsertionSort(){

    nums := []int{5,3,8,1,13,4,9,2}

    for i :=0 ; i < len(nums) ; i++ {
        for k := i ; k > 0; k--{
            if nums[k] < nums[k-1] {
                min := nums[k-1]
                nums[k-1] = nums[k]
                nums[k] = min
            }
        }

    }

    fmt.Println("插入排序 => ", nums)
}



//合併排序 Merge Sort (遞迴)
func MergeSort() {
    var list []int = []int{5,3,8,1,13,4,9,2}
    fmt.Println("合併排序 => ", subMergeSort(list))
}


func subMergeSort(list []int) []int {
  var length = len(list)
  if length < 2 {
    return list
  }
  var mid = length / 2
  return merge(subMergeSort(list[:mid]), subMergeSort(list[mid:]))
}

func merge(x, y []int) []int {
  var r []int = make([]int, len(x)+len(y))
  for i, j := 0, 0; ; {
    if i < len(x) && (j == len(y) || x[i] < y[j]) {
      r[i+j] = x[i]
      i++
    } else if j < len(y) {
      r[i+j] = y[j]
      j++
    } else {
      break
    }
  }
  return r
}



// 快速排序 Quick Sort
func QuickSort() {
    var list []int = []int{5,3,8,1,13,4,9,2,15}
    quickly(list)
    fmt.Println("快速排序 => ", list)
}


func quickly(list []int){
  left, right := 1, len(list)-1

  if left != right && len(list) > 2 {
    for left < right {

      if list[left] <= list[0] && left < right{
        left++
      }else if list[right] > list[0] && left < right{
        right--
      }else if list[left] > list[right] && left < right && list[left] > list[0] {
        target := list[left]
        list[left] = list[right]
        list[right] = target
      }
    }

    if(left == right){
      target := list[0]
      list[0] = list[left-1]
      list[left-1] = target
    }

    quickly(list[:left])
    quickly(list[left:])
  }
}




