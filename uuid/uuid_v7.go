package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"
)

type UUIDv7 struct {
	value [16]byte
	dt    time.Time
}

func NewUUIDv7(t time.Time) (u UUIDv7) {
	return UUIDv7{
		value: [16]byte{},
		dt:    t,
	}
}

// Gerar UUIDv7
// seguindo o padrão RFC 9562
// Para manter os padrões de fuso horário o UUIDv7 sempre será gerado em UTC
// https://www.ietf.org/archive/id/draft-peabody-dispatch-new-uuid-format-04.html#name-uuid-version-7
// https://datatracker.ietf.org/doc/html/rfc9562#name-uuid-version-7
func (u *UUIDv7) generate() error {
	// bytes randomicos
	var value [16]byte
	_, err := rand.Read(value[:])
	if err != nil {
		return err
	}

	// obter o timestamp em milissegundos
	timestamp := big.NewInt(u.dt.UTC().UnixMilli())

	// seguindo o padrão UUIDv7(RFC 9562)
	// preencher os primeiros 6 bytes com o timestamp
	timestamp.FillBytes(value[0:6])

	// definir a versão e variante
	value[6] = (value[6] & 0x0F) | 0x70
	value[8] = (value[8] & 0x3F) | 0x80

	u.value = value

	return nil

}

func (u *UUIDv7) Generate() (string, error) {
	err := u.generate()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (u UUIDv7) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u.value[:4], u.value[4:6], u.value[6:8], u.value[8:10], u.value[10:])
}

func GenerateUUIDv7() (string, error) {

	u := NewUUIDv7(time.Now().UTC())

	return u.Generate()

}

func GetTimeUUIDv7(uuid string) (time.Time, error) {

	dataUuid, err := parseUUID(uuid)

	if err != nil || len(dataUuid) != 16 {
		return time.Time{}, fmt.Errorf("invalid UUIDv7 format")
	}

	timestamp := int64(dataUuid[0])<<40 |
		int64(dataUuid[1])<<32 |
		int64(dataUuid[2])<<24 |
		int64(dataUuid[3])<<16 |
		int64(dataUuid[4])<<8 |
		int64(dataUuid[5])

	return time.UnixMilli(timestamp).UTC(), nil
}

func parseUUID(uuid string) ([16]byte, error) {
	var data [16]byte

	uuidStr := strings.ReplaceAll(uuid, "-", "")

	byts, err := hex.DecodeString(uuidStr)

	if err != nil {
		return [16]byte{}, fmt.Errorf("invalid UUIDv7 format: %s", err)
	}

	if len(byts) != 16 {
		return [16]byte{}, fmt.Errorf("invalid UUIDv7 length, must be 16 bytes")
	}

	copy(data[:], byts)

	return data, nil

}
