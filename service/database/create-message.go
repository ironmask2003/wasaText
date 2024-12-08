package database

import (
	"database/sql"
	"errors"
	"strconv"
)

// Query for create a message in the message table
var queryAddMessage = "INSERT INTO message (MessageId, Message, SendUserId, ConversationId) VALUES (?, ?, ?, ?) "

// Query for get the last ID used in the conversation table
var queryLastMessageId = "SELECT MAX(MessageId) FROM message WHERE ConversationId = ? AND SendUserId = ?"

func GetLastElemMessage(db *appdbimpl, convId int, senderUserId int) (int, error) {
	var _maxID = sql.NullInt64{Int64: 0, Valid: false}
	row, err := db.c.Query(queryLastMessageId, convId, senderUserId)
	if err != nil {
		return 0, err
	}
	var maxID int
	for row.Next() {
		if row.Err() != nil {
			return 0, err
		}
		err = row.Scan(&_maxID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return 0, err
		}
		if !_maxID.Valid {
			maxID = 0
		} else {
			maxID = int(_maxID.Int64)
		}
	}
	return maxID, nil
}

func (db *appdbimpl) CreateMessage(m Message) (Message, error) {
	var msg Message

	// Set the user id, conversation id and the text
	msg.ConversationId = m.ConversationId
	msg.SenderUserId = m.SenderUserId
	msg.Text = m.Text

	// Get conv of rcv
	conv, err := db.GetConversationById(msg.ConversationId, msg.SenderUserId)
	if err != nil {
		return msg, err
	}

	convRcv, err := db.GetConversationsBySender(msg.SenderUserId, conv.SenderUserId)
	if err != nil {
		// Getting the max id in the conversation table
		maxID, err := GetLastElemMessage(db, msg.ConversationId, msg.SenderUserId)
		if err != nil {
			return msg, err
		}

		msg.MessageId = maxID + 1
		// -- INSERT THE MESSAGE IN THE DATABSE -- //
		_, err = db.c.Exec(queryAddMessage, msg.MessageId, msg.Text, msg.SenderUserId, msg.ConversationId)
		if err != nil {
			return msg, err
		}
		return msg, nil
	}
	maxRcv, err := GetLastElemMessage(db, convRcv, conv.SenderUserId)
	if err != nil {
		return msg, err
	}

	// Getting the max id in the conversation table
	maxID, err := GetLastElemMessage(db, conv.ConversationId, msg.SenderUserId)
	if err != nil {
		return msg, err
	}

	if maxID > maxRcv {
		msg.MessageId = maxID + 1
	} else {
		msg.MessageId = maxRcv + 1
	}

	// -- INSERT THE MESSAGE IN THE DATABSE -- //
	_, err = db.c.Exec(queryAddMessage, msg.MessageId, msg.Text, msg.SenderUserId, msg.ConversationId)
	if err != nil {
		return msg, errors.New("Error creating the message " + strconv.Itoa(msg.MessageId))
	}

	return msg, nil
}
