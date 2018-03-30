package game

var gp *Processor

type Proxy struct {
}

/**
proxy类主要负责接收对应逻辑的信息，然后处理逻辑

ProcessMessage ：
	proxy 类中唯一暴露出来的方法。接收信息，处理信息，返回信息

 */
func (p *Proxy) ProcessMessage(code int, msg []uint8) {

}

func (p *Proxy) SendMessage() {

}

func NewProxy(processor *Processor) *Proxy {
	proxy := &Proxy{}
	gp = processor
	return proxy
}
