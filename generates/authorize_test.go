package generates

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/oauth2.v2"
	"gopkg.in/oauth2.v2/models"
)

func TestAuthorize(t *testing.T) {
	Convey("Test Authorize Generate", t, func() {
		data := &oauth2.GenerateBasic{
			Client: &models.Client{
				ID:     "123456",
				Secret: "123456",
			},
			UserID:   "000000",
			CreateAt: time.Now(),
		}
		gen := NewAuthorizeGenerate()
		code, err := gen.Token(data)
		So(err, ShouldBeNil)
		So(code, ShouldNotBeEmpty)
	})
}
