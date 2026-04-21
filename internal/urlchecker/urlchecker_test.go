package urlchecker

import (
	"testing"
)

type URLCheckerTest struct {
	Name               string
	URL                string
	ExpectedStatusCode int
	ExpectedError      bool
}

func TestHealthyURL(t *testing.T) {

	tests := []URLCheckerTest{
		{Name: "Testing api.github.com", URL: "https://api.github.com", ExpectedStatusCode: 200, ExpectedError: false},
		{Name: "Testing github.com", URL: "http://github.com", ExpectedStatusCode: 200, ExpectedError: false},
		{Name: "Testing cloudflare.com", URL: "https://cloudflare.com", ExpectedStatusCode: 200, ExpectedError: false},
		{Name: "Testing google.com", URL: "https://www.google.com", ExpectedStatusCode: 200, ExpectedError: false},
	}

	for i := range tests {
		t.Run(tests[i].Name, func(t *testing.T) {
			res := URLChecker(tests[i].URL)

			if tests[i].ExpectedError != (res.Err != nil) {
				t.Errorf("%s failed; %s want=%t Got=%t Error=%v", tests[i].Name, "ExpectedError", tests[i].ExpectedError, (res.Err != nil), res.Err)
			}

			if tests[i].ExpectedStatusCode != res.StatusCode {
				t.Errorf("%s failed; %s want=%d Got=%d", tests[i].Name, "StatusCode", tests[i].ExpectedStatusCode, res.StatusCode)
			}

		})
	}

}

func TestTimeoutURL(t *testing.T) {
	tests := []URLCheckerTest{
		{
			Name:               "No Timeout",
			URL:                "https://google.com",
			ExpectedStatusCode: 200,
			ExpectedError:      false,
		},
		{
			Name:               "Timeout 3s delay",
			URL:                "https://httpbin.org/delay/3",
			ExpectedStatusCode: 0,
			ExpectedError:      true,
		},
		{
			Name:               "Timeout 5s delay",
			URL:                "https://httpbin.org/delay/5",
			ExpectedStatusCode: 0,
			ExpectedError:      true,
		},
	}

	for i := range tests {
		t.Run(tests[i].Name, func(t *testing.T) {
			res := URLChecker(tests[i].URL)

			if tests[i].ExpectedError != (res.Err != nil) {
				t.Errorf("%s failed; %s want=%t Got=%t Error=%v", tests[i].Name, "ExpectedError", tests[i].ExpectedError, (res.Err != nil), res.Err)
			}

			if tests[i].ExpectedStatusCode != res.StatusCode {
				t.Errorf("%s failed; %s want=%d Got=%d", tests[i].Name, "StatusCode", tests[i].ExpectedStatusCode, res.StatusCode)
			}

		})
	}
}

func TestRedirectURL(t *testing.T) {
	tests := []URLCheckerTest{
		{
			Name:               "Redirect github.com (http → https)",
			URL:                "http://github.com",
			ExpectedStatusCode: 200,
			ExpectedError:      false,
		},
		{
			Name:               "Redirect google.com (http → https)",
			URL:                "http://google.com",
			ExpectedStatusCode: 200,
			ExpectedError:      false,
		},
		{
			Name:               "Httpbin single redirect",
			URL:                "http://httpbin.org/redirect/1",
			ExpectedStatusCode: 200,
			ExpectedError:      false,
		},
		{
			Name:               "Httpbin multiple redirect",
			URL:                "http://httpbin.org/redirect/3",
			ExpectedStatusCode: 200,
			ExpectedError:      false,
		},
	}

	for i := range tests {
		t.Run(tests[i].Name, func(t *testing.T) {
			res := URLChecker(tests[i].URL)

			if tests[i].ExpectedError != (res.Err != nil) {
				t.Errorf("%s failed; %s want=%t Got=%t Error=%v", tests[i].Name, "ExpectedError", tests[i].ExpectedError, (res.Err != nil), res.Err)
			}

			if tests[i].ExpectedStatusCode != res.StatusCode {
				t.Errorf("%s failed; %s want=%d Got=%d", tests[i].Name, "StatusCode", tests[i].ExpectedStatusCode, res.StatusCode)
			}

		})
	}
}

func TestInvalidSSLURL(t *testing.T) {
	tests := []URLCheckerTest{
		{
			Name:               "Expired SSL",
			URL:                "https://expired.badssl.com/",
			ExpectedStatusCode: 0,
			ExpectedError:      true,
		},
		{
			Name:               "Self Signed SSL",
			URL:                "https://self-signed.badssl.com/",
			ExpectedStatusCode: 0,
			ExpectedError:      true,
		},
		{
			Name:               "Wrong Host SSL",
			URL:                "https://wrong.host.badssl.com/",
			ExpectedStatusCode: 0,
			ExpectedError:      true,
		},
		{
			Name:               "Non-existent domain",
			URL:                "http://invalid.url.test",
			ExpectedStatusCode: 0,
			ExpectedError:      true,
		},
		{
			Name:               "Connection refused",
			URL:                "http://localhost:9999",
			ExpectedStatusCode: 0,
			ExpectedError:      true,
		},
	}

	for i := range tests {
		t.Run(tests[i].Name, func(t *testing.T) {
			res := URLChecker(tests[i].URL)

			if tests[i].ExpectedError != (res.Err != nil) {
				t.Errorf("%s failed; %s want=%t Got=%t Error=%v", tests[i].Name, "ExpectedError", tests[i].ExpectedError, (res.Err != nil), res.Err)
			}

			if tests[i].ExpectedStatusCode != res.StatusCode {
				t.Errorf("%s failed; %s want=%d Got=%d", tests[i].Name, "StatusCode", tests[i].ExpectedStatusCode, res.StatusCode)
			}

		})
	}
}
