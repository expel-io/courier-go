package courier_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/expel-io/courier-go"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {

	expectedResponseID := "123456789"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/send", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rw.Write([]byte(fmt.Sprintf("{ \"MessageId\" : \"%s\" }", expectedResponseID)))

		}))
	defer server.Close()

	t.Run("sends request", func(t *testing.T) {

		courier := courier.CourierClient("key", server.URL)

		myData := struct {
			foo string
		}{
			foo: "bar",
		}
		myProfile := struct {
			email string
		}{
			email: "foo@bar.com",
		}
		messageID, err := courier.Send(context.Background(), "event-id", "recipient-id", myProfile, myData)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, messageID)

	})

}
