package implement

import (
	"events"
	"log"
	"time"

	"producer/errs"
	repository "producer/repository/kafka"
)

type Service interface {
	Sender(input *events.Messages) (output *Response, err error)
}

type Messages struct {
	MsgId  int    `json:"Msg_id"`
	Sender string `json:"Sender"`
	Msg    string `json:"Msg"`
}

type senderService struct {
	kafkaRepo repository.Kafka
}

func New(kafkaRepo repository.Kafka) (service Service) {
	return &senderService{kafkaRepo: kafkaRepo}
}

type Response struct {
	ReceivedTime time.Time `json:"Received_Time"`
}

func (impl *senderService) Sender(input *events.Messages) (output *Response, err error) {
	err = impl.kafkaRepo.Produce(input)
	if err != nil {
		log.Println(err)
		return nil, errs.NewUnexpectedError()
	}

	output = &Response{}
	output.ReceivedTime = time.Now()

	return output, nil
}
