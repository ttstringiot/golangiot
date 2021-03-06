package header

// GetProducerConnectionListRequestHeader 获得生产者连接信息请求头
// Author rongzhihong
// Since 2017/9/19
type GetProducerConnectionListRequestHeader struct {
	ProducerGroup string `json:"producerGroup"`
}

func (header *GetProducerConnectionListRequestHeader) CheckFields() error {
	return nil
}

// NewGetProducerConnectionListRequestHeader 初始化
// Author: tianyuliang
// Since: 2017/11/6
func NewGetProducerConnectionListRequestHeader(producerGroup string) *GetProducerConnectionListRequestHeader {
	producerConnectionListRequestHeader := &GetProducerConnectionListRequestHeader{
		ProducerGroup: producerGroup,
	}
	return producerConnectionListRequestHeader
}
