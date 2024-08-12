package utils

import (
	"io"
	"net/http"
	"net/http/httputil"
)

func DumpRequest(writer io.Writer, header string, r *http.Request) error {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	writer.Write([]byte("\n" + header + " Request: \n"))
	writer.Write(data)
	writer.Write([]byte("\n----------------------------------- \n"))
	return nil
}
func DumResponse(writer io.Writer, header string, r *http.Response) error {
	data, err := httputil.DumpResponse(r, true)
	if err != nil {
		return err
	}
	writer.Write([]byte("\n" + header + " Response: \n"))
	writer.Write(data)
	writer.Write([]byte("\n----------------------------------- \n"))
	return nil
}
