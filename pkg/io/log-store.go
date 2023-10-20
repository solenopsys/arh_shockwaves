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
			//	print(r)

			s.statistic[message.Key] = s.statistic[message.Key] + count
			if err != nil {
				Panic(err)
			}
		}
	}
}

func (s *LogStore) GetStatistic(key string) int {
	if s.store[key] == nil {
		return 0
	}
	return s.statistic[key]
}

func (s *LogStore) GetLog(key string) string {
	if s.store[key] == nil {
		return ""
	}
	return s.store[key].(*strings.Builder).String()
}

var logStoreOnce sync.Once
var logStoreInstance *LogStore

func GetLogStore() *LogStore {
	logStoreOnce.Do(func() {
		logStoreInstance = &LogStore{
			MessagesStream: make(chan LogMessage),
			store:          make(map[string]io.StringWriter),
			statistic:      make(map[string]int),
		}
		go logStoreInstance.Processing()
	})
	return logStoreInstance
}
