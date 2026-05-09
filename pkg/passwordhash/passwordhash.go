package passwordhash

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"math/big"
)

// alphabet must stay byte-for-byte identical to the legacy implementation.
const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Hash returns md5(password+salt) hex-encoded. Compatible with previously stored
// hashes produced by user-data-service.
func Hash(password, salt string) string {
	h := md5.New()
	io.WriteString(h, password+salt)
	return hex.EncodeToString(h.Sum(nil))
}

// Salt returns a cryptographically random alphanumeric string of length n.
func Salt(n int) (string, error) {
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", err
		}
		out[i] = alphabet[idx.Int64()]
	}
	return string(out), nil
}
