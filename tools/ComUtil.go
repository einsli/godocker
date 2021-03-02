package tools

func LimitVerify(skip int, limit int, dataLen int) (int, int) {
	// 需要验证逻辑
	if dataLen == 1{
		return 0, 1
	} else if dataLen == 0 {
		return 0, 0
	} else{
		var (
			itemRemain int
		)
		// 当翻页数小于数组总数
		if (skip + 1) * limit < dataLen {
			return skip * limit, (skip * limit) + limit
		} else {
			// 当容器数第一页数目小于 limit
			if dataLen < limit && skip == 0{
				itemRemain = dataLen
				return 0, dataLen
			}else {
				itemRemain = dataLen - (skip * limit)
				return (skip - 1) * limit , ((skip - 1) * limit) + itemRemain
			}
		}
	}
}
