package mysql_db

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"time"
)

func GetMd5(data string) string {
	md5Byte := md5.Sum([]byte(data))
	return hex.EncodeToString(md5Byte[:])
}

func GetFnv128(data string) string {
	hasher := fnv.New128()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetUrlShortKey(url string) string {
	return GetMd5(url) + GetFnv128(url)
}

func GetByteLenByPartitionCount(count int) int {
	count -= 1
	if count <= 0 {
		return 0
	}
	res := 0
	for ; count > 0; count /= 16 {
		res += 1
	}
	return res
}

func GetTableNumFromRawLinkKey(partitionCount int, key string) int {
	if partitionCount <= 1 {
		return 0
	}
	keyNum := HexStrToInt64(key)
	res := keyNum % uint64(partitionCount)
	return int(res)
}

func GetTableNumFromShortLinkPath(partitionCount int, shortLinkPath string) int {
	if partitionCount <= 1 {
		return 0
	}
	partitionLen := GetByteLenByPartitionCount(partitionCount)
	hexStr := shortLinkPath[:partitionLen]
	return int(HexStrToInt64(hexStr))
}

func HexStrToInt64(hexStr string) uint64 {
	const HEX_LEN = 16
	if len(hexStr) > HEX_LEN {
		hexStr = hexStr[:HEX_LEN]
	}
	if len(hexStr) < HEX_LEN {
		hexStr = fmt.Sprintf("%016s", hexStr)
	}
	inputBytes, _ := hex.DecodeString(hexStr)
	bytesBuffer := bytes.NewBuffer(inputBytes)
	res := uint64(0)
	binary.Read(bytesBuffer, binary.BigEndian, &res)
	return res
}

func Int64ToHexStr(n uint64) string {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	resBytes := bytesBuffer.Bytes()
	resHex := hex.EncodeToString(resBytes)
	location := 0
	for ; location < len(resHex) && resHex[location] == '0'; location++ {
	}
	if location == len(resHex) {
		return "0"
	}
	return resHex[location:]
}

func GetShortLinkPrefix(partitionCount int, tableNum int) string {
	hexStr := Int64ToHexStr(uint64(tableNum))
	partitionLen := GetByteLenByPartitionCount(partitionCount)
	return fmt.Sprintf("%0"+fmt.Sprintf("%vs", partitionLen), hexStr)
}

func MsTimeStampToTime(timeStamp int64) time.Time {
	return time.Unix(timeStamp/1e3, (timeStamp%1e3)*1e6)
}
