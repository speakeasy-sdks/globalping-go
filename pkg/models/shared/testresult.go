// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/globalping-go/pkg/utils"
)

type TestResultType string

const (
	TestResultTypeInProgressTestResult         TestResultType = "InProgressTestResult"
	TestResultTypeFailedTestResult             TestResultType = "FailedTestResult"
	TestResultTypeFinishedPingTestResult       TestResultType = "FinishedPingTestResult"
	TestResultTypeFinishedTracerouteTestResult TestResultType = "FinishedTracerouteTestResult"
	TestResultTypeFinishedDNSTestResult        TestResultType = "FinishedDnsTestResult"
	TestResultTypeFinishedMtrTestResult        TestResultType = "FinishedMtrTestResult"
	TestResultTypeFinishedHTTPTestResult       TestResultType = "FinishedHttpTestResult"
)

type TestResult struct {
	InProgressTestResult         *InProgressTestResult
	FailedTestResult             *FailedTestResult
	FinishedPingTestResult       *FinishedPingTestResult
	FinishedTracerouteTestResult *FinishedTracerouteTestResult
	FinishedDNSTestResult        *FinishedDNSTestResult
	FinishedMtrTestResult        *FinishedMtrTestResult
	FinishedHTTPTestResult       *FinishedHTTPTestResult

	Type TestResultType
}

func CreateTestResultInProgressTestResult(inProgressTestResult InProgressTestResult) TestResult {
	typ := TestResultTypeInProgressTestResult

	return TestResult{
		InProgressTestResult: &inProgressTestResult,
		Type:                 typ,
	}
}

func CreateTestResultFailedTestResult(failedTestResult FailedTestResult) TestResult {
	typ := TestResultTypeFailedTestResult

	return TestResult{
		FailedTestResult: &failedTestResult,
		Type:             typ,
	}
}

func CreateTestResultFinishedPingTestResult(finishedPingTestResult FinishedPingTestResult) TestResult {
	typ := TestResultTypeFinishedPingTestResult

	return TestResult{
		FinishedPingTestResult: &finishedPingTestResult,
		Type:                   typ,
	}
}

func CreateTestResultFinishedTracerouteTestResult(finishedTracerouteTestResult FinishedTracerouteTestResult) TestResult {
	typ := TestResultTypeFinishedTracerouteTestResult

	return TestResult{
		FinishedTracerouteTestResult: &finishedTracerouteTestResult,
		Type:                         typ,
	}
}

func CreateTestResultFinishedDNSTestResult(finishedDNSTestResult FinishedDNSTestResult) TestResult {
	typ := TestResultTypeFinishedDNSTestResult

	return TestResult{
		FinishedDNSTestResult: &finishedDNSTestResult,
		Type:                  typ,
	}
}

func CreateTestResultFinishedMtrTestResult(finishedMtrTestResult FinishedMtrTestResult) TestResult {
	typ := TestResultTypeFinishedMtrTestResult

	return TestResult{
		FinishedMtrTestResult: &finishedMtrTestResult,
		Type:                  typ,
	}
}

func CreateTestResultFinishedHTTPTestResult(finishedHTTPTestResult FinishedHTTPTestResult) TestResult {
	typ := TestResultTypeFinishedHTTPTestResult

	return TestResult{
		FinishedHTTPTestResult: &finishedHTTPTestResult,
		Type:                   typ,
	}
}

func (u *TestResult) UnmarshalJSON(data []byte) error {

	inProgressTestResult := new(InProgressTestResult)
	if err := utils.UnmarshalJSON(data, &inProgressTestResult, "", true, true); err == nil {
		u.InProgressTestResult = inProgressTestResult
		u.Type = TestResultTypeInProgressTestResult
		return nil
	}

	failedTestResult := new(FailedTestResult)
	if err := utils.UnmarshalJSON(data, &failedTestResult, "", true, true); err == nil {
		u.FailedTestResult = failedTestResult
		u.Type = TestResultTypeFailedTestResult
		return nil
	}

	finishedTracerouteTestResult := new(FinishedTracerouteTestResult)
	if err := utils.UnmarshalJSON(data, &finishedTracerouteTestResult, "", true, true); err == nil {
		u.FinishedTracerouteTestResult = finishedTracerouteTestResult
		u.Type = TestResultTypeFinishedTracerouteTestResult
		return nil
	}

	finishedMtrTestResult := new(FinishedMtrTestResult)
	if err := utils.UnmarshalJSON(data, &finishedMtrTestResult, "", true, true); err == nil {
		u.FinishedMtrTestResult = finishedMtrTestResult
		u.Type = TestResultTypeFinishedMtrTestResult
		return nil
	}

	finishedPingTestResult := new(FinishedPingTestResult)
	if err := utils.UnmarshalJSON(data, &finishedPingTestResult, "", true, true); err == nil {
		u.FinishedPingTestResult = finishedPingTestResult
		u.Type = TestResultTypeFinishedPingTestResult
		return nil
	}

	finishedHTTPTestResult := new(FinishedHTTPTestResult)
	if err := utils.UnmarshalJSON(data, &finishedHTTPTestResult, "", true, true); err == nil {
		u.FinishedHTTPTestResult = finishedHTTPTestResult
		u.Type = TestResultTypeFinishedHTTPTestResult
		return nil
	}

	finishedDNSTestResult := new(FinishedDNSTestResult)
	if err := utils.UnmarshalJSON(data, &finishedDNSTestResult, "", true, true); err == nil {
		u.FinishedDNSTestResult = finishedDNSTestResult
		u.Type = TestResultTypeFinishedDNSTestResult
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TestResult) MarshalJSON() ([]byte, error) {
	if u.InProgressTestResult != nil {
		return utils.MarshalJSON(u.InProgressTestResult, "", true)
	}

	if u.FailedTestResult != nil {
		return utils.MarshalJSON(u.FailedTestResult, "", true)
	}

	if u.FinishedPingTestResult != nil {
		return utils.MarshalJSON(u.FinishedPingTestResult, "", true)
	}

	if u.FinishedTracerouteTestResult != nil {
		return utils.MarshalJSON(u.FinishedTracerouteTestResult, "", true)
	}

	if u.FinishedDNSTestResult != nil {
		return utils.MarshalJSON(u.FinishedDNSTestResult, "", true)
	}

	if u.FinishedMtrTestResult != nil {
		return utils.MarshalJSON(u.FinishedMtrTestResult, "", true)
	}

	if u.FinishedHTTPTestResult != nil {
		return utils.MarshalJSON(u.FinishedHTTPTestResult, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}