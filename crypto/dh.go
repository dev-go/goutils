// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"crypto/rand"
	"io"
	"math/big"
)

// DHPulbic : public value
type DHPulbic struct {
	p *big.Int // public base
	g *big.Int // public modulus
}

// NewDHPublic : create custom DHPublic
// WARNING: Do not use it if you do not know what you are doing
func NewDHPublic(base *big.Int, modulus *big.Int) (*DHPulbic, error) {
	if base == nil || modulus == nil {
		return nil, ErrParam
	}
	return &DHPulbic{
		p: base,
		g: modulus,
	}, nil
}

// GetDHPublic : get a DHPulbic by its ID as defined in either RFC 2409 or RFC 3526
func GetDHPublic(id int) (*DHPulbic, error) {
	if id <= 0 {
		id = 14
	}
	switch id {
	case 1:
		p, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A63A3620FFFFFFFFFFFFFFFF", 16)
		return &DHPulbic{
			g: new(big.Int).SetInt64(2),
			p: p,
		}, nil
	case 2:
		p, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE65381FFFFFFFFFFFFFFFF", 16)
		return &DHPulbic{
			g: new(big.Int).SetInt64(2),
			p: p,
		}, nil
	case 14:
		p, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AACAA68FFFFFFFFFFFFFFFF", 16)
		return &DHPulbic{
			g: new(big.Int).SetInt64(2),
			p: p,
		}, nil
	default:
		return nil, ErrParam
	}
}

// DHKey : key pair
type DHKey struct {
	pub *DHPulbic

	x *big.Int // private key
	y *big.Int // public key
}

// SetPrivateKey : set x
func (dhKey *DHKey) SetPrivateKey(private []byte) {
	if private == nil {
		return
	}
	dhKey.x = new(big.Int).SetBytes(private)
}

// SetPublicKey : set y
func (dhKey *DHKey) SetPublicKey(public []byte) {
	if public == nil {
		return
	}
	dhKey.y = new(big.Int).SetBytes(public)
}

// Bytes : output bytes
func (dhKey *DHKey) Bytes() []byte {
	if dhKey.y == nil {
		return make([]byte, 0)
	}
	if dhKey.pub != nil {
		var size = (dhKey.pub.p.BitLen() + 7) / 8
		var result = make([]byte, size)
		copyWithLeftPaddingZero(result, dhKey.y.Bytes())
		return result
	}
	return dhKey.y.Bytes()
}

// copyWithLeftPaddingZero : copy src to the end of dest, padding with zero bytes as needed
func copyWithLeftPaddingZero(dst []byte, src []byte) {
	var num = len(dst) - len(src)
	for i := 0; i < num; i++ {
		dst[i] = 0
	}
	copy(dst[num:], src)
}

// String : implement Stringer interface
func (dhKey *DHKey) String() string {
	if dhKey.y == nil {
		return ""
	}
	return dhKey.y.String()
}

// GenerateKey : generate a key pair
func GenerateKey(dhPulbic *DHPulbic, r io.Reader) (*DHKey, error) {
	if r == nil {
		r = rand.Reader
	}
	// x should be in (0, p).
	var x = big.NewInt(0)
	var zero = big.NewInt(0)
	for x.Cmp(zero) == 0 {
		var err error
		if x, err = rand.Int(r, dhPulbic.p); err != nil {
			return nil, err
		}
	}
	// y = g ^ x mod p
	var y = new(big.Int).Exp(dhPulbic.g, x, dhPulbic.p)
	return &DHKey{
		pub: dhPulbic,
		x:   x,
		y:   y,
	}, nil
}

// CalculateExchangeKey : calculate exchange key
func CalculateExchangeKey(comm *DHPulbic, private *DHKey, public *DHKey) (*big.Int, error) {
	if comm == nil || comm.p == nil || comm.g == nil ||
		private.x == nil || public.y == nil {
		return nil, ErrParam
	}
	var k = new(big.Int).Exp(public.y, private.x, comm.p)
	return k, nil
}
