package envelope

import (
	"errors"
	"fmt"

	"github.com/notaryproject/notation-core-go/signature/cose"
	"github.com/notaryproject/notation-core-go/signature/jws"
	gcose "github.com/veraison/go-cose"
)

// Supported envelope format.
const (
	COSE = "cose"
	JWS  = "jws"
)

// SpeculateSignatureEnvelopeFormat speculates envelope format by looping all builtin envelope format.
//
// TODO: abandon this fature in RC1.
func SpeculateSignatureEnvelopeFormat(raw []byte) (string, error) {
	var msg gcose.Sign1Message
	if err := msg.UnmarshalCBOR(raw); err == nil {
		return cose.MediaTypeEnvelope, nil
	}
	if len(raw) == 0 || raw[0] != '{' {
		// very certain
		return "", errors.New("unsupported signature format")
	}
	return jws.MediaTypeEnvelope, nil
}

// GetEnvelopeMediaType converts the envelope type to mediaType name.
func GetEnvelopeMediaType(sigFormat string) (string, error) {
	switch sigFormat {
	case JWS:
		return jws.MediaTypeEnvelope, nil
	case COSE:
		return cose.MediaTypeEnvelope, nil
	}
	return "", fmt.Errorf("signature format %q not supported", sigFormat)
}
