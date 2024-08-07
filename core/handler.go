package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"github.com/aodr3w/keiji-core/logging"
)

const (
	PULL_PORT = ":8006"
	PUSH_PORT = ":8005"
)

type handlerType int

const (
	PULL handlerType = iota
	PUSH
)

func errorResp(msg string) string {
	return fmt.Sprintf("ERROR: %v", msg)
}

func HandlePush(mq *MessageQueue, conn net.Conn, logger *logging.Logger) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		var message Message
		err := json.Unmarshal(scanner.Bytes(), &message)
		if err != nil {
			logger.Error(err.Error())
			fmt.Fprint(conn, errorResp(err.Error()))
			continue
		}
		if len(mq.messages)+1 > mq.cap {
			fmt.Fprint(conn, errorResp("too many requests"))
		}
		mq.Push(message)
		logger.Info("message pushed.")
		fmt.Fprint(conn, "OK")
		break
	}
	if err := scanner.Err(); err != nil {
		logger.Error(err.Error())
		fmt.Fprint(conn, errorResp("Error reading message: "+err.Error()))
	}
}

func HandlePull(mq *MessageQueue, conn net.Conn, logger *logging.Logger) {
	defer conn.Close()
	message := mq.Pull()
	msgBytes, err := json.Marshal(message)
	if err != nil {
		logger.Error(err.Error())
		fmt.Fprint(conn, errorResp("Error marshalling message:  "+err.Error()))
		return
	}
	fmt.Fprint(conn, string(msgBytes))
	logger.Info("message pulled.")
}
