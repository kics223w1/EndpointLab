package api

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouterStatus() *gin.Engine {
	r := gin.Default()
	httpStatus := NewHttpStatus()

	r.Any("/status/:code", httpStatus.HandleStatus)

	return r
}

func getStatusCodes() []int {
	return []int{
		http.StatusContinue,
		http.StatusSwitchingProtocols,
		http.StatusProcessing,
		http.StatusEarlyHints,
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNonAuthoritativeInfo,
		http.StatusNoContent,
		http.StatusResetContent,
		http.StatusPartialContent,
		http.StatusMultiStatus,
		http.StatusAlreadyReported,
		http.StatusIMUsed,
		http.StatusMultipleChoices,
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusNotModified,
		http.StatusUseProxy,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect,
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusPaymentRequired,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusMethodNotAllowed,
		http.StatusNotAcceptable,
		http.StatusProxyAuthRequired,
		http.StatusRequestTimeout,
		http.StatusConflict,
		http.StatusGone,
		http.StatusLengthRequired,
		http.StatusPreconditionFailed,
		http.StatusRequestEntityTooLarge,
		http.StatusRequestURITooLong,
		http.StatusUnsupportedMediaType,
		http.StatusRequestedRangeNotSatisfiable,
		http.StatusExpectationFailed,
		http.StatusTeapot,
		http.StatusMisdirectedRequest,
		http.StatusUnprocessableEntity,
		http.StatusLocked,
		http.StatusFailedDependency,
		http.StatusTooEarly,
		http.StatusUpgradeRequired,
		http.StatusPreconditionRequired,
		http.StatusTooManyRequests,
		http.StatusRequestHeaderFieldsTooLarge,
		http.StatusUnavailableForLegalReasons,
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
		http.StatusHTTPVersionNotSupported,
		http.StatusVariantAlsoNegotiates,
		http.StatusInsufficientStorage,
		http.StatusLoopDetected,
		http.StatusNotExtended,
		http.StatusNetworkAuthenticationRequired,
	}
}

func TestStatusCodeWithMethodGET(t *testing.T) {
	router := setupRouterStatus()

	for _, code := range getStatusCodes() {
		t.Run(http.StatusText(code), func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/status/"+strconv.Itoa(code), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != code {
				t.Errorf("Expected status code %d, but got %d", code, w.Code)
			}
		})
	}
}

func TestStatusCodeWithMethodPOST(t *testing.T) {
	router := setupRouterStatus()

	for _, code := range getStatusCodes() {
		t.Run(http.StatusText(code), func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/status/"+strconv.Itoa(code), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != code {
				t.Errorf("Expected status code %d, but got %d", code, w.Code)
			}
		})
	}
}

