package sockets

import (
	"github.com/theflyingcodr/sockets"

	"context"

	"github.com/theflyingcodr/sockets/server"
)

type payment struct {
}

func NewPayment() *payment {
	return &payment{}
}

// Register will register new handler/s with the socket server.
func (p *payment) Register(s *server.SocketServer) {
	s.RegisterChannelHandler("payment", p.payment)
	s.RegisterChannelHandler("payment.ack", p.paymentAck)
}

// buildPaymentRequest will forward a payment message to all connected clients.
func (p *payment) payment(ctx context.Context, msg *sockets.Message) (*sockets.Message, error) {
	return msg, nil
}

// buildPaymentRequest will forward a payment.ack message to all connected clients.
func (p *payment) paymentAck(ctx context.Context, msg *sockets.Message) (*sockets.Message, error) {
	return msg, nil
}
