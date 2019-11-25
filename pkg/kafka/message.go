package kafka_operator

import "encoding/json"

type (
	Message struct {
		name    string
		encoded []byte
		err     error
	}
)

func (M *Message) Length() int {
	b, e := json.Marshal(M)
	M.encoded = b
	M.err = e
	return len(string(b))
}

func (M *Message) Encode() ([]byte, error) {
	return M.encoded, M.err
}
