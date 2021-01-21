package test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/wingfeng/idx/utils"
)

func TestHashString(t *testing.T) {
	s := utils.HashString("vue_secret")

	s1 := utils.HashString(`eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ2dWVfY2xpZW50IiwiZXhwIjoxNjExMTU2NzExLCJzdWIiOiI3YTQ1Y2I1NC1iMGZmLTRlY2QtOTViOS0wNzRkMzNhYWFjMWUifQ.Cf29KL_pfoGyAx-CezXOIbjF-BDcGV-qucIlX1kXokywpt9WYciLKVv4Khxp4WbGzkbj1fXY4t8ksHDZ7-xauVnDIXkH65ZrY2Pl1v21O1tB_E5MsHik-gZHOjgEF7QuPxN0egKPhzHiLWQmP_f-5tPaZFCCzQ4LJ9UDBdymgurL-q_kzZ3CPrFdLYSFEoHMWFxVuB2UiyPO2uMYpvqYHSoHmSLFhD3wrHFfsk14x0O91NQJF3lDXKWv3YMX2o5wfElBHLo8dGmIsc6braPu5UBsMXmmzXESlMHO168sytuZ4-8y1m2WK_OUU2lLYwLVNx-aRJFmrreyuJXBi-NaYQ`)
	t.Logf("result:%s", s)
	t.Logf("result for long test:%s", s1)
	assert.Equal(t, len(s), 64)
	assert.Equal(t, len(s1), 64)
}

func TestHashToken(t *testing.T) {
	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjMyRTg4NDNBN0ZDNEVFRjZEREZENzM2OUJDNkE0MTc3NDk0NTI1NDUiLCJ0eXAiOiJhdCtqd3QiLCJ4NXQiOiJNdWlFT25fRTd2YmRfWE5wdkdwQmQwbEZKVVUifQ.eyJuYmYiOjE2MTExNTM3NTIsImV4cCI6MTYxMTE1NzM1MiwiaXNzIjoiaHR0cHM6Ly9zdHMucWlhbnFpdXNvZnQuY29tIiwiY2xpZW50X2lkIjoianNjbGllbnQiLCJzdWIiOiJlYWVmZjMwYi0yZDU3LTQ3MzItYjVhMi1kM2Y2ZjBhNWE3MTAiLCJhdXRoX3RpbWUiOjE2MTExNTM3NTIsImlkcCI6ImxvY2FsIiwic2NvcGUiOlsib3BlbmlkIiwiZW1haWwiXSwiYW1yIjpbInB3ZCJdfQ.Vhxa5W82eQSntUj3R8eKQTaOn0Uz3FouIkf6pIZJl9VtYOT8Por4u68WYr6WktcemYH0DLEHX3D5UuwW2zgO91bIMHes5QodTmnZU7ErAQmmco7FCkK9_qO0s7UhWTekqUK_iHOiaF8gNfiYPWifK5Eq-TdkgNWe4Sx7-5zKGXK-yISAqxFJcGCtGaF8f1Rmtb_IMjZjpRN91knGTyG1Gx6kVOT5tdPL4KJZ2qpuX6gBhGF-nsb8smlMuJbEOpyaNe10-YcYzesLGkQDP-7iKvqidjxQdRIbczGUiusMM-Unj7aIXN_ziB_j642u0_izzAntTS9RwlLZwJliSuvjTw"
	s := utils.HashAccessToken(token)
	t.Logf("at_hash:%s", s)
	assert.Equal(t, s, "-dAlvgCDkzyBPWIxMLZFmQ")
}
