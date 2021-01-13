package ex

func Clamp (num int,low int,up int) interface{}{
	if low > up{
		return nil
	}
	if num < low {
		return low
	} else if num > up{
		return up
	} else {
		return num
	}

}