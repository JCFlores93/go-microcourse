package polo
import (
	"github.com/JCFlores93/go-microcourse/src/api/utils/test_utils"
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestPolo(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/marco", nil)
	c := test_utils.GetMockedContext(request, response)

	Marco(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}