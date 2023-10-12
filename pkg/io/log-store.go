package io

import (
	"io"
	"strings"
	"sync"
)

type LogMessage struct {
	Message string
	Key     string
}

type LogStore struct {
	MessagesStream chan LogMessage
	store          map[string]io.StringWriter
	statistic      map[string]int
}

func (s *LogStore) Processing() {
	for s.MessagesStream != nil {
		select {
		case message := <-s.MessagesStream:
			r := strings.Replace(message.Message, "\n", "\r\n", -1)

			if s.store[message.Key] == nil {
				s.store[message.Key] = &strings.Builder{}
				s.statistic[message.Key] = 0
			}

			count, err := s.store[message.Key].WriteString(r)

			s.statistic[message.Key] = s.statistic[message.Key] + count
			if err != nil {
				Panic(err)
			}
		}
	}
}

var logStoreOnce sync.Once
var logStoreInstance *LogStore

func GetLogStore() *LogStore {
	logStoreOnce.Do(func() {
		logStoreInstance = &LogStore{}
		go logStoreInstance.Processing()
	})
	return logStoreInstance
}
