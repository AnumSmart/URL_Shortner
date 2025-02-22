package stat

import (
	"log"
	"server/pkg/event"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatSrvice struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(deps *StatServiceDeps) *StatSrvice {
	return &StatSrvice{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatSrvice) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Bad EventLinkVisited Data", msg.Data)
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}
}
