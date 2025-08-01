func generate(numRows int) [][]int {
    var res [][]int = [][]int{
        {1},
        {1, 1},
    }   
    
    if numRows == 1 {
        return res[:1]
    }    
    
    for i := 2; i < numRows; i++ {
        elem := res[len(res) - 1]
        size := len(elem)
        arr := make([]int, size + 1)
        arr[0] = 1
        arr[len(arr) - 1] = 1
        
        for j := 1; j < size; j++ {
            arr[j] = elem[j] + elem[j - 1]
        } 
        
        res = append(res, arr)
    }

    return res
}