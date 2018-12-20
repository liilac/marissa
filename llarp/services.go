package llarp

import (
	"github.com/tv42/zbase32"
	"github.com/zeebo/bencode"
	"golang.org/x/crypto/blake2b"
)

const tldSuffix = ".loki"

// ServiceAddress is the 256-bit blake2b hash of the bencoded form of the
// ServiceInfo struct
type ServiceAddress [32]byte

func (sa ServiceAddress) String() string {
	return zbase32.EncodeToString(sa[:]) + tldSuffix
}

// ServiceInfo describes the "public information blob" for a hidden service
type ServiceInfo struct {
	// long term public key used for encryption
	EncryptionPublicKey [32]byte `json:"e"`
	// long term public key used for signatures
	SigningPublicKey [32]byte `json:"s"`
	// protocol version
	Version uint `json:"v"`
	// nonce used to enable vanity addresses
	Nonce [16]byte `json:"x"`
}

// Bencoded returns a bencoded form of the ServiceInfo struct
func (si ServiceInfo) Bencoded() (be []byte, err error) {
	return bencode.EncodeBytes(si)
}

// ServiceAddress returns the service address for the ServiceInfo struct,
// defined as Blake2b-256(Bencode(ServiceInfo))
func (si ServiceInfo) ServiceAddress() (sa ServiceAddress, err error) {
	var be []byte
	be, err = si.Bencoded()
	if err != nil {
		return
	}

	return ServiceAddress(blake2b.Sum256(be)), nil
}
