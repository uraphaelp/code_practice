func maxProfit(prices []int) int {
    var i, ref, tmp int
    l:=len(prices)
    if l==0 {
        return 0
    }
    profit:=0
    ref=prices[0]
    for i=1; i<l; i++ {
        tmp=prices[i]-ref
        //更新利润值
        if tmp>profit {
            profit=tmp
        //更新对比值/序列最小值
        } else if tmp<0 {
            ref=prices[i]
        }
    }
    if profit==0 {
        return 0
    } else {
        return profit
    }
}
