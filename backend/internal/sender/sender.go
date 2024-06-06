package sender

import (
	"backend/internal/auth"
	"backend/internal/rabbitmq"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"net/http"
)

type sender struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}
type Sender interface {
	AuthLoginHandler(c *gin.Context)
	AuthRegisterHandler(c *gin.Context)
}

func NewSender(conn *amqp.Connection, ch *amqp.Channel) Sender {
	return &sender{conn, ch}
}

func (s *sender) AuthLoginHandler(c *gin.Context) {
	var req auth.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reqString, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = rabbitmq.PublishMessages(s.ch, "auth_queue", string(reqString))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (s *sender) AuthRegisterHandler(c *gin.Context) {}
