package smbauth

import (
	"context"

	"github.com/marmos91/dittofs/internal/adapter/smb/auth"
)

// NTLMPasswordValidator allows custom NTLM password validation.
// When set on the SMB adapter, it takes priority over the UserStore for every
// NTLM AUTHENTICATE message. Return (zero, false) to reject the client.
type NTLMPasswordValidator interface {
	// ValidateNTLMv2 validates a client's NTLMv2 response.
	// domain is the value the client sent in the AUTHENTICATE message (may be empty).
	// Returns the 16-byte session base key on success, or (zero, false) on failure.
	ValidateNTLMv2(ctx context.Context, username, domain string, serverChallenge [8]byte, ntChallengeResponse []byte) ([16]byte, bool)
}

// ComputeNTHash computes the NT hash (MD4 of UTF-16LE password) for a plaintext password.
func ComputeNTHash(password string) [16]byte {
	return auth.ComputeNTHash(password)
}

// ValidateNTLMv2Response validates a client's NTLMv2 response against an NT hash.
// Returns the 16-byte session base key on success.
func ValidateNTLMv2Response(ntHash [16]byte, username, domain string, serverChallenge [8]byte, ntChallengeResponse []byte) ([16]byte, error) {
	return auth.ValidateNTLMv2Response(ntHash, username, domain, serverChallenge, ntChallengeResponse)
}
