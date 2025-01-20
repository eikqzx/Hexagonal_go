package persistence

import (
	"fmt"
	"os"
)

// ReadSQLFromFile อ่านคำสั่ง SQL จากไฟล์ .sql และส่งคืนเป็นสตริง
func ReadSQLFromFile(filePath string) (string, error) {
	// อ่านไฟล์ SQL
	sqlBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading SQL file %s: %v", filePath, err)
	}
	return string(sqlBytes), nil
}
