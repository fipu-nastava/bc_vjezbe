package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"io"
)

type Message struct {
	Sender	string
	Msg 	string
	Nonce 	int
}

func (m *Message) Hash() Hash {
	sum := sha256.Sum256(m.Serialize())
	return Hash(sum)
}

func (m *Message) Serialize() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)

	err := enc.Encode(m)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func Deserialize(reader io.Reader) (*Message, error) {

	var m Message

	dec := gob.NewDecoder(reader)
	err := dec.Decode(&m)

	return &m, err
}