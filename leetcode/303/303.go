package main

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		data: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.data[i]
	}
	return sum
}

func main() {

}
