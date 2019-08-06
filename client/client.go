// Package client contains the implementation of a `dasy` client.
package client

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/vacp2p/dasy/protobuf"
	mvds "github.com/vacp2p/mvds/node"
	mvdsproto "github.com/vacp2p/mvds/protobuf"
	"github.com/vacp2p/mvds/state"
	"github.com/vacp2p/mvds/store"
)

// Chat is the ID for a specific chat.
type Chat state.GroupID

// Peer is the ID for a specific peer.
type Peer state.PeerID

// Client is the actual daisy client.
type Client struct {
	node  mvds.Node
	store store.MessageStore // @todo we probably need a different message store, not sure tho

	lastMessage state.MessageID // @todo maybe make type
}

// Invite invites a peer to a chat.
func (c *Client) Invite(chat Chat, peer Peer) {

}

// Join joins a chat.
func (c *Client) Join(chat Chat) (state.MessageID, error) {
	return c.send(chat, protobuf.Message_JOIN, c.node.ID[:])
}

// Leave leaves a chat.
func (c *Client) Leave(chat Chat) (state.MessageID, error) {
	return c.send(chat, protobuf.Message_LEAVE, c.node.ID[:])
}

// Kick kicks peer from a chat.
func (c *Client) Kick(chat Chat, peer Peer) (state.MessageID, error) {
	return c.send(chat, protobuf.Message_KICK, peer[:])
}

// Ack acknowledges `Join`, `Leave` and `Kick` messages.
func (c *Client) Ack(chat Chat, messageID state.MessageID) (state.MessageID, error) {
	// @todo We may not need this as we can rely on the acks of data sync
	return c.send(chat, protobuf.Message_ACK, messageID[:])
}

// Post sends a message to a chat.
func (c *Client) Post(chat Chat, body []byte) (state.MessageID, error) {
	return c.send(chat, protobuf.Message_POST, body)
}

func (c *Client) send(chat Chat, t protobuf.Message_MessageType, body []byte) (state.MessageID, error) {
	msg := &protobuf.Message{
		MessageType:     protobuf.Message_MessageType(t),
		Body:            body,
		PreviousMessage: c.lastMessage[:],
	}

	// @todo sign

	buf, err := proto.Marshal(msg)
	if err != nil {
		return state.MessageID{}, err
	}

	id, err := c.node.AppendMessage(state.GroupID(chat), buf)
	if err != nil {
		return state.MessageID{}, err
	}

	c.lastMessage = id

	return id, nil
}

func (c *Client) onReceive(message mvdsproto.Message) {
	var msg protobuf.Message
	err := proto.Unmarshal(message.Body, &msg)
	if err != nil {
		log.Printf("error while unmarshalling message: %s", err.Error())
		return
	}

	// @todo pump messages to subscriber channels

	if len(msg.PreviousMessage) == 0 {
		return
	}

	c.handlePreviousMessage(
		bytesToGroupID(message.GroupId),
		bytesToMessageID(msg.PreviousMessage),
	)
}

func (c *Client) handlePreviousMessage(group state.GroupID, previousMessage state.MessageID) {
	if c.store.Has(previousMessage) {
		return
	}

	err := c.node.RequestMessage(group, previousMessage)
	if err != nil {
		log.Printf("error while requesting message: %s", err.Error())
	}
}
