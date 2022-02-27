package service

import (
	"context"
	"encoding/json"
	"events"
	"log"

	"consumer/config"
	"consumer/domain"
	repository "consumer/repository/database"
)

type EventHandleService interface {
	Handle(topic string, eventBytes []byte)
	SaveData(ctx context.Context, eventBytes []byte) (err error)
}

type dataEventHandler struct {
	dbRepo repository.DatabaseRepo
	config *config.Config
}

func NewHandleData(dbRepo repository.DatabaseRepo, appConfig *config.Config) (service EventHandleService) {
	return &dataEventHandler{dbRepo: dbRepo, config: appConfig}
}


func (e *dataEventHandler) Handle(topic string, eventBytes []byte){
	ctx := context.Background()
	switch topic {
	case e.config.MessageTopic:
		err := e.SaveData(ctx, eventBytes)
		if err != nil {
			return
		}
	default:
		log.Println("no event handler")
	}
}

func (e *dataEventHandler) SaveData(ctx context.Context, eventBytes []byte) (err error){
	event := &events.Messages{}
	err = json.Unmarshal(eventBytes,event)
	if err != nil {
		log.Println(err)
		return err
	}

	ent := &domain.Message{
		MsgId: event.MsgId,
		Sender: event.Sender+"++",
		Msg: event.Msg,
	}

	err = e.dbRepo.Save(ctx, ent)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(ent)
	return  nil
}