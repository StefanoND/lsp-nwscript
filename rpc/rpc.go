package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// Have a message come in, serialize it and create this message into a json
func EncodeMessage(msg any) string {
	// We must first turn whatever message we got into a json blob
	content, err := json.Marshal(msg)
	// If can't handle it, panic with err
	if err != nil {
		panic(err)
	}

	// Sprintf, String print formatted, print formatted string
	// Figure how many bytes the "content" is and place it at %d
	// Place content at %s
	// return the above string formatted
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)

	// The output shold be like this:
	/*
		Content-Length: ...\r\n
		\r\n
		{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "textDocument/completion",
			"params": {
			...
			}
		}
	*/
}

// Take something in json and decode it to an actual format
func DecodeMessage(msg []byte) (int, error) {
  // E need to find the header and know where it ends by knowing where content starts
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})

  // If didn't find, return error
	if !found {
		return 0, errors.New("Did not find separator")
	}

  // With the above we'll get "Content-Length: <number>"
  // Now we need to know how many bytes to read after <number>.
  contentLengthBytes := header[len("Content-Length: "):]
  // We're going convert a string (strconv) from ASCI to int (Atoi)
  contentLength, err := strconv.Atoi(string(contentLengthBytes))
  if err != nil {
    return 0, err
  }

  // TODO: Implement content
  _ = content

  // If found, return out contentLength
	return contentLength, nil
}
