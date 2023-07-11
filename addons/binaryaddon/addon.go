package binaryaddon

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"

	"github.com/herb-go/herbplugin"
)

type Addon struct {
	Plugin herbplugin.Plugin
}

func (a *Addon) Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
func (a *Addon) Base64Decode(data string) []byte {
	result, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return result
}
func (a *Addon) Md5Sum(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
func (a *Addon) Sha1Sum(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}
func (a *Addon) Sha256Sum(data []byte) string {
	h := sha256.New()
	h.Write(data)
	hash := h.Sum(data)
	return hex.EncodeToString(hash[:])
}
func (a *Addon) Sha512Sum(data []byte) string {
	h := sha512.New()
	h.Write(data)
	hash := h.Sum(data)
	return hex.EncodeToString(hash[:])
}
func Create(p herbplugin.Plugin) *Addon {
	return &Addon{
		Plugin: p,
	}
}
