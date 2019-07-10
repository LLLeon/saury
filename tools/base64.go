package tools

// Base64StrToLengthInByte converts the specified base64 encoded string to the corresponding byte length.
// A formula to calculate: base64Len=2^6, bytesLen=2^8, base64Len^4=byteLen^3, so bytesLen=base64Len*3/4.
func Base64StrToLengthInByte(s string) int {
	bytesSize := len(s) * 3 / 4
	return bytesSize
}
