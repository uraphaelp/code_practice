func generate(numRows int) [][]int {
    if numRows==0 {
        return [][]int{}    
    }
    var result [][]int
    var i, j, l int
    var now, last []int
    //第一行的元素是1，且不满足迭代规则
    result=append(result, []int{1})
    for i=1; i<numRows; i++ {
        //now代表当前需要新生成的一行；且第一个元素总为1
        now=[]int{1}
        //last代表上一行
        last=result[i-1]
        j=0
        l=len(last)
        //本行=上一行“前元素”+“后元素”生成
        for j!=l-1 {
            now=append(now, last[j]+last[j+1])
            j++
        }
        //每行最后一个元素为1
        now=append(now, 1)
        result=append(result, now)
    }
    return result
}
