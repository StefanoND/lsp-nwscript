package rpc_test

import (
	"lsp-nwscript/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

// Test to see if we can encode a message
// It is case sensitive, so Testing and testing will not match and fail
func TestEncode(t *testing.T) {
  // This is the message we expect, where 16  is the count from "{" to "}" after the "\n"
  // Ignoring the "\" before the ".
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"

  // This is the actual message we'll get to compare to expected
	actual := rpc.EncodeMessage(EncodingExample{Testing:true})

  // Fail if the actual message doesn't match the expected message
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}

  // You can run "go test ./..." at terminal on the root folder of this project
}

func TestDecode(t *testing.T) {
  // This is the message we expect, where 16  is the count from "{" to "}" after the "\n"
  // Ignoring the "\" before the ".
	incomingMessage := "Content-Length: 16\r\n\r\n{\"Testing\":true}"

  contentLength, err := rpc.DecodeMessage([]byte(incomingMessage))
  if err != nil {
    t.Fatal(err)
  }

  if contentLength != 16 {
    t.Fatalf("Expected: 16, Actual: %d", contentLength)
  }
}
