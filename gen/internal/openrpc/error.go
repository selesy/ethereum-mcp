package openrpc

import "errors"

// ErrUnmarshalingMethod is returned when an attempt to unmarshal one or
// more OpenRPC methods in a file from the Ethereum execution apis fails.
var ErrUnmarshalingMethod = errors.New("error unmarshaling method")

// ErrUnmarshalingParams is returned when an attempt to unmarshal one or
// more OpenRPC method parameters in a file from the Ethereum execution
// apis fails.
var ErrUnmarshalingParams = errors.New("error unmarshaling params")
