package packet

// Head 包头
type Head byte

const (
	// HeadStart 开始包
	HeadStart Head = 0
	// HeadBody 中间包
	HeadBody Head = 1
	// HeadEnd 结束包
	HeadEnd Head = 2
	// HeadClosure 数据量少的包，直接闭包
	HeadClosure Head = 3
)
