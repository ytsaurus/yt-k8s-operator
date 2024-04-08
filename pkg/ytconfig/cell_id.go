package ytconfig

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func generateCellID(cellTag int16) string {
	cellID, err := uuid.NewRandomFromReader(strings.NewReader("ytsaurus-kubernetes-operator"))
	if err != nil {
		panic(err)
	}
	uuidBytes, err := cellID.MarshalBinary()
	if err != nil {
		panic(err)
	}

	uuidBytes[4] = byte(cellTag >> 8)
	uuidBytes[5] = byte(cellTag & 0xff)

	masterCellType := 601
	uuidBytes[6] = byte(masterCellType >> 8)
	uuidBytes[7] = byte(masterCellType & 0xff)

	getGUIDPart := func(data []byte) string {
		format := strings.Repeat("%02x", len(data))
		args := make([]any, 0, len(data))
		for _, value := range data {
			args = append(args, value)
		}

		part := fmt.Sprintf(format, args...)
		return strings.TrimLeft(part, "0")
	}
	return fmt.Sprintf("%s-%s-%s-%s", getGUIDPart(uuidBytes[12:]), getGUIDPart(uuidBytes[8:12]), getGUIDPart(uuidBytes[4:8]), getGUIDPart(uuidBytes[:4]))
}

func RandString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
