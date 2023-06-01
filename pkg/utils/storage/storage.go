package storage

func findHigh2Pow(num int) int {
	temp := num - 1
	temp |= temp >> 1
	temp |= temp >> 2
	temp |= temp >> 4
	temp |= temp >> 8
	temp |= temp >> 16

	if temp < 0 {
		return 1
	} else {
		return temp + 1
	}
}

func findLow2Pow(num int) int {
	return findHigh2Pow(num) >> 1
}

func CalculatePg(p_pgs, i_pgs int) Pools {
	var pools = GetDefalutPoolSize()

	// 计算data存储池pg大小
	p_highPow := findHigh2Pow(p_pgs)
	p_lowPow := findLow2Pow(p_pgs)

	if p_highPow-p_pgs > p_pgs-p_lowPow {
		pools.Sp.Data = p_lowPow
	} else {
		pools.Sp.Data = p_highPow
	}

	//计算index存储池pg大小
	i_highPow := findHigh2Pow(i_pgs)
	i_lowPow := findLow2Pow(i_pgs)

	if i_highPow-i_pgs > i_pgs-i_lowPow {
		pools.Sp.Index = i_lowPow
	} else {
		pools.Sp.Index = i_highPow
	}

	return pools
}
