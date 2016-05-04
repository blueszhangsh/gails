package gails

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

//Aes encryptor
type Aes struct {
	Cip cipher.Block
}

//Encode encoder
func (p *Aes) Encode(buf []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(p.Cip, iv)
	val := make([]byte, len(buf))
	cfb.XORKeyStream(val, buf)

	return append(val, iv...), nil
}

//Decode  decoder
func (p *Aes) Decode(buf []byte) ([]byte, error) {
	bln := len(buf)
	cln := bln - aes.BlockSize
	ct := buf[0:cln]
	iv := buf[cln:bln]

	cfb := cipher.NewCFBDecrypter(p.Cip, iv)
	val := make([]byte, cln)
	cfb.XORKeyStream(val, ct)
	return val, nil

}
