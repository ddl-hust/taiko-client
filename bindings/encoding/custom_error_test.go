package encoding

import (
	"errors"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type testJsonError struct{}

func (e *testJsonError) Error() string { return common.Bytes2Hex(randomBytes(10)) }

func (e *testJsonError) ErrorData() interface{} { return "0x8a1c400f" }

type emptyTestJsonError struct{}

func (e *emptyTestJsonError) Error() string { return "execution reverted" }

func (e *emptyTestJsonError) ErrorData() interface{} { return "0x" }

func TestTryParsingCustomError(t *testing.T) {
	randomErr := common.Bytes2Hex(randomBytes(10))
	require.Equal(t, randomErr, TryParsingCustomError(errors.New(randomErr)).Error())

	err := TryParsingCustomError(errors.New(
		// L1_INVALID_BLOCK_ID
		"VM Exception while processing transaction: reverted with an unrecognized custom error (return data: 0x8a1c400f)",
	))

	require.True(t, strings.HasPrefix(err.Error(), "L1_INVALID_BLOCK_ID"))

	err = TryParsingCustomError(&testJsonError{})

	require.True(t, strings.HasPrefix(err.Error(), "L1_INVALID_BLOCK_ID"))

	err = TryParsingCustomError(&emptyTestJsonError{})

	require.Equal(t, err.Error(), "execution reverted")
}
