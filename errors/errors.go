package errors

import serrors "errors"

var (
	// ErrClientNotFound 客户端不存在
	ErrClientNotFound = serrors.New("客户端不存在")
	// ErrClientExist 客户端已存在
	ErrClientExist = serrors.New("客户端已存在")
	// ErrReadDataLenght 读取数据错误
	ErrReadDataLenght = serrors.New("数据错误，长度小于5")
)
