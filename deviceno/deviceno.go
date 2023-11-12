package deviceno

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/golangrustnode/comutils/comfile"
	"github.com/google/uuid"
)

/*
1.生成uuid
2.hash(uuid)
3.hash(prefix+hash(uuid))
*/
func GenerateDeviceNoWithPrefix(prefix, path string) {
	deviceNo := GenerateUniStringWithPrefix(prefix)
	comfile.Write2file(deviceNo, path)
}
func GenerateUniStringWithPrefix(prefix string) string {
	//1.生成uuid
	id := uuid.New()
	//2.生成uuid的hash
	md5Hasher1 := md5.New()
	md5Hasher1.Write([]byte(id.String()))
	hash1 := md5Hasher1.Sum(nil)
	hash1String := hex.EncodeToString(hash1)
	//3.prefix+uuid的再hash
	md5Hasher2 := md5.New()
	md5Hasher2.Write([]byte(hash1String + prefix))
	hash2 := md5Hasher2.Sum(nil)
	hash2String := hex.EncodeToString(hash2)
	return prefix + hash1String + hash2String[:3]
}
