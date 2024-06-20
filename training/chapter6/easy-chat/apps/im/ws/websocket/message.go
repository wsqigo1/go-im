package websocket

import "time"

type FrameType uint8

const (
	FrameData  FrameType = 0x0
	FramePing  FrameType = 0x1
	FrameAck   FrameType = 0x2
	FrameNoAck FrameType = 0x3
	FrameErr   FrameType = 0x9

	//FrameHeaders      FrameType = 0x1
	//FramePriority     FrameType = 0x2
	//FrameRSTStream    FrameType = 0x3
	//FrameSettings     FrameType = 0x4
	//FramePushPromise  FrameType = 0x5
	//FrameGoAway       FrameType = 0x7
	//FrameWindowUpdate FrameType = 0x8
	//FrameContinuation FrameType = 0x9
)

type Message struct {
	FrameType `json:"frameType"`
	Id        string    `json:"id"`
	AckSeq    int       `json:"ackSeq"`
	AckTime   time.Time `json:"ackTime"`
	ErrCount  int       `json:"errCount"`
	Method    string    `json:"method"` // 标识具体要处理的方法
	FormId    string    `json:"formId"` // 请求的来源
	Data      any       `json:"data"`   // map[string]interface{}，传递的参数
}

func NewMessage(formId string, data any) *Message {
	return &Message{
		FrameType: FrameData,
		FormId:    formId,
		Data:      data,
	}
}

func NewErrMessage(err error) *Message {
	return &Message{
		FrameType: FrameErr,
		Data:      err.Error(),
	}
}
