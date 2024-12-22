package racer

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	fastestServerTests := []struct {
		serverA *httptest.Server
		serverB *httptest.Server
		want    string
	}{
		{serverA: makeDelayedServer(0 * time.Millisecond), serverB: makeDelayedServer(10 * time.Millisecond), want: "A"},
		{serverA: makeDelayedServer(10 * time.Millisecond), serverB: makeDelayedServer(0 * time.Millisecond), want: "B"},
	}

	for i, tt := range fastestServerTests {
		t.Run("returns fastest server #"+strconv.Itoa(i), func(t *testing.T) {
			defer tt.serverA.Close()
			defer tt.serverB.Close()

			got, err := Racer(tt.serverA.URL, tt.serverB.URL, 10*time.Second)
			assertNoErr(t, err)

			switch tt.want {
			case "A":
				assertEqual(t, got, tt.serverA.URL)
			case "B":
				assertEqual(t, got, tt.serverB.URL)
			default:
				t.Errorf("Invalid server wanted: %q", tt.want)
			}
		})
	}

	t.Run("returns an error if a server doesn't respond within given timeout", func(t *testing.T) {
		server := makeDelayedServer(10 * time.Millisecond)
		defer server.Close()

		_, err := Racer(server.URL, server.URL, 5*time.Microsecond)
		assertErr(t, err, "")
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func assertEqual[V comparable](t testing.TB, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func assertErr(t testing.TB, err error, msg string) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted error but did not get one")
	}
	if len(msg) > 0 && err.Error() != msg {
		t.Errorf("got %q but wanted %q", err.Error(), msg)
	}
}

func assertNoErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("wanted no error but got %q", err.Error())
	}
}
