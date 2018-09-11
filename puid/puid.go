package puid

import (
	"bytes"
	"encoding/base32"
	"math/rand"
	"os"
	"strings"
	"time"
	"utils"
)

/*
PUID is a 128 bit implementation comparable to UUIDs it combines
*/
type PUID struct {
	number [16]byte
}

func (id *PUID) String() string {
	encoded := new(bytes.Buffer)
	encoder := base32.NewEncoder(base32.HexEncoding, encoded)
	_, err := encoder.Write(id.number[:])
	utils.Log(err, "Error in conversion from PUID to base32 string")
	defer encoder.Close()
	return strings.TrimRight(encoded.String(), "=")
}

func generatePUID(unixTime, unixNanoTime int64, processID uint32) *PUID {
	var myUUID [16]byte
	//Take the low side of unix time
	_, unixSecondsLowSide := splitInt64Time(unixTime)
	//Take Both sides of nano unix time
	_, unixNanoLowSide := splitInt64Time(unixNanoTime)
	for x := 0; x < 4; x++ {
		rand.Seed(unixTime)
		random := rand.Uint32()
		myUUID[x] = byte(unixSecondsLowSide >> uint8(x*8))
		myUUID[x+3*4] = byte(random >> uint8(x*8))
		myUUID[x+2*4] = byte(unixNanoLowSide >> uint8(x*8))
	}
	//Low side of the processID
	myUUID[4+2], myUUID[4+3] = byte(processID>>8), byte(processID)
	//ASCII signature :3
	myUUID[4+0], myUUID[4+1] = byte(0x4A), byte(0x52)

	return &PUID{myUUID}
}

/*
NewPUID creates a UUID using the unix clock and the process's ID
*/
func NewPUID() *PUID {
	return generatePUID(time.Now().Unix(), time.Now().UnixNano(), uint32(os.Getpid()))
}

func splitInt64Time(time int64) (uint32, uint32) {
	return uint32(time >> 32), uint32(time)
}
