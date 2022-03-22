package errors

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// Default error type for internal telemetry
type Errors struct {
	errorList []errorData
}

type errorData struct {
	err      error
	callSite string
}

func NewError(err ...error) Errors {
	res := make([]errorData, 0)

	for _, e := range err {
		res = append(res, errorData{
			err: e,
			// 2 so it skips the getMetadata() and this function
			callSite: getMetadata(2),
		})
	}

	return Errors{errorList: res}
}

func (self *Errors) Error() string {
	sb := strings.Builder{}

	for i, data := range self.errorList {
		sb.WriteString(fmt.Sprintf("{%03d} %s:\n\t%s\n", i, data.callSite, data.err.Error()))
	}

	return sb.String()
}

// Returns the last added error, if any
func (self *Errors) Last() error {
	l := len(self.errorList)

	if l == 0 {
		return nil
	}

	return self.errorList[l-1].err
}

func (self *Errors) AddError(err error) {
	self.errorList = append(self.errorList, errorData{
		err: err,
		// 2 so it skips the getMetadata() and this function
		callSite: getMetadata(2),
	})
}

func getMetadata(skip int) string {
	pc, filename, line, _ := runtime.Caller(skip)

	return fmt.Sprintf("[%s] at %s [%s:%d]",
		time.Now().Format("2006-01-02 15:04:05 -0700 MST 2006"),
		filename,
		runtime.FuncForPC(pc).Name(),
		line)
}
