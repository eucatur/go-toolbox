package signature

import (
	"crypto/tls"

	"github.com/amdonov/xmlsig"
)

func SingStruct(certFile, keyFile string, i interface{}) (sig *xmlsig.Signature, err error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return
	}

	s, err := xmlsig.NewSigner(cert)
	if err != nil {
		return
	}

	sig, err = s.CreateSignature(i)

	return
}
