package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	log "github.com/cihub/seelog"

	"golang.org/x/crypto/pbkdf2"
)

/*
为了兼容Asp.net core 3.1的密码
代码参考:https://github.com/dotnet/AspNetCore/blob/master/src/Identity/Extensions.Core/src/PasswordHasher.cs
*/

const (
	PWD_SHA1 = iota
	PWD_SHA256
)

func GenHashedPWD(pwd string) string {
	/*
			HashPasswordV3(password, rng,
		                prf: KeyDerivationPrf.HMACSHA256,
		                iterCount: _iterCount,
		                saltSize: 128 / 8,
		                numBytesRequested: 256 / 8);
			Produce a version 3 (see comment above) text hash.
			byte[] salt = new byte[saltSize];
			rng.GetBytes(salt);
			byte[] subkey = KeyDerivation.Pbkdf2(password, salt, prf, iterCount, numBytesRequested);

			var outputBytes = new byte[13 + salt.Length + subkey.Length];
			outputBytes[0] = 0x01; // format marker
			WriteNetworkByteOrder(outputBytes, 1, (uint)prf);
			WriteNetworkByteOrder(outputBytes, 5, (uint)iterCount);
			WriteNetworkByteOrder(outputBytes, 9, (uint)saltSize);
			Buffer.BlockCopy(salt, 0, outputBytes, 13, salt.Length);
			Buffer.BlockCopy(subkey, 0, outputBytes, 13 + saltSize, subkey.Length);
			return outputBytes;
	*/
	var iterCount, saltSize int
	iterCount = 10000
	saltSize = 128 / 8
	salt := make([]byte, saltSize)
	fmt.Println(salt)
	rand.Read(salt)
	fmt.Println(salt) //生成密文
	subkey := pbkdf2.Key([]byte(pwd), salt, iterCount, 256/8, sha256.New)
	var outputBytes = make([]byte, 13+len(salt)+len(subkey))
	outputBytes[0] = 0x01                             // format marker
	writeNetworkByteOrder(outputBytes, 1, PWD_SHA256) //使用SHA256
	writeNetworkByteOrder(outputBytes, 5, uint(iterCount))
	writeNetworkByteOrder(outputBytes, 9, uint(saltSize))
	blockCopy(salt, 0, outputBytes, 13, len(salt))
	blockCopy(subkey, 0, outputBytes, 13+saltSize, len(subkey))
	return base64.StdEncoding.EncodeToString(outputBytes)
}
func VerifyPassword(hashedPassword string, password string) (bool, error) {
	pwdBytes, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false, err
	}
	result, _ := VerifyHashedPasswordV3(pwdBytes, password)
	return result, nil
}
func VerifyHashedPasswordV3(hashedPassword []byte, password string) (bool, int) {
	var iterCount = 0

	// Read header information
	prf := readNetworkByteOrder(hashedPassword, 1) //用来决定是选SHA1还是SHA256.
	log.Debugf("Prf:%v", prf)

	iterCount = int(readNetworkByteOrder(hashedPassword, 5))
	log.Debugf("inter Count:%d", iterCount)
	saltLength := int(readNetworkByteOrder(hashedPassword, 9))
	log.Debugf("salt Length:%d", saltLength)
	// Read the salt: must be >= 128 bits
	if saltLength < 128/8 {
		return false, iterCount
	}

	var salt = make([]byte, saltLength)
	blockCopy(hashedPassword, 13, salt, 0, saltLength)

	// Read the subkey (the rest of the payload): must be >= 128 bits
	subkeyLength := len(hashedPassword) - 13 - saltLength
	if subkeyLength < 128/8 {
		return false, iterCount
	}
	expectedSubkey := make([]byte, subkeyLength)
	blockCopy(hashedPassword, 13+len(salt), expectedSubkey, 0, len(expectedSubkey))

	// Hash the incoming password and verify it

	actualSubkey := pbkdf2.Key([]byte(password), salt, iterCount, subkeyLength, sha256.New)
	return byteArraysEqual(actualSubkey, expectedSubkey), iterCount

}
func readNetworkByteOrder(buffer []byte, offset int) uint {
	return ((uint)(buffer[offset+0]) << 24) | ((uint)(buffer[offset+1]) << 16) | ((uint)(buffer[offset+2]) << 8) | ((uint)(buffer[offset+3]))
}

func writeNetworkByteOrder(buffer []byte, offset int, value uint) {
	buffer[offset+0] = (byte)(value >> 24)
	buffer[offset+1] = (byte)(value >> 16)
	buffer[offset+2] = (byte)(value >> 8)
	buffer[offset+3] = (byte)(value >> 0)
}
func blockCopy(src []byte, srcOffset int, dst []byte, dstOffset, count int) (bool, error) {
	srcLen := len(src)
	if srcOffset > srcLen || count > srcLen || srcOffset+count > srcLen {
		return false, errors.New("源缓冲区 索引超出范围")
	}
	dstLen := len(dst)
	if dstOffset > dstLen || count > dstLen || dstOffset+count > dstLen {
		return false, errors.New("目标缓冲区 索引超出范围")
	}
	index := 0
	for i := srcOffset; i < srcOffset+count; i++ {
		dst[dstOffset+index] = src[srcOffset+index]
		index++
	}
	return true, nil
}

func byteArraysEqual(a []byte, b []byte) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || len(a) != len(b) {
		return false
	}
	areSame := true
	for i := 0; i < len(a); i++ {
		areSame = areSame && (a[i] == b[i])
	}
	return areSame
}
