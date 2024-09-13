package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.praktikum-services.ru/Stasyan/momo-store/cmd/api/dependencies"
)

func TestFakeAppIntegrational(t *testing.T) {
	store, err := dependencies.NewFakeDumplingsStore()
	assert.NoError(t, err)
	app, err := NewInstance(store)
	assert.NoError(t, err)

	t.Run("create_order", func(t *testing.T) {
		for i := 1; i <= 10; i++ {
			t.Run("id"+strconv.Itoa(i), func(t *testing.T) {
				r := httptest.NewRequest("POST", "/orders", nil)
				w := httptest.NewRecorder()
				app.CreateOrderController(w, r)

				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
				fmt.Fprintln(os.Stdout, "_____")
				fmt.Fprintln(os.Stdout, w.Body.String())
				fmt.Fprintln(os.Stdout, "_____")

				expectedJSON, err := json.Marshal(map[string]interface{}{"id": i})
				assert.NoError(t, err)
				assert.JSONEq(t, string(expectedJSON), w.Body.String())
			})
		}
	})

	t.Run("list_dumplings", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/packs", nil)
		w := httptest.NewRecorder()
		app.ListDumplingsController(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		fmt.Fprintln(os.Stdout, "_____")
		fmt.Fprintln(os.Stdout, w.Body.String())
		fmt.Fprintln(os.Stdout, "_____")

		expectedJSON := "{\"results\":[{\"id\":1,\"name\":\"Пельмени\",\"price\":5,\"description\":\"С говядиной\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Pel%27meni&Vontony.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T130328Z&X-Amz-Expires=2592000&X-Amz-Signature=CE4BAE513B89EE16C8BAA019D783A7508350BC34C6D1B03FB605FCD0A1CEDF46&X-Amz-SignedHeaders=host\"},{\"id\":2,\"name\":\"Хинкали\",\"price\":3.5,\"description\":\"Со свининой\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Hinkali&Baoczy.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T125035Z&X-Amz-Expires=2592000&X-Amz-Signature=47634DAF429B9CF1DFA236884F530DD654DC76D4E14D48CD56BCCEA63F7BFDB0&X-Amz-SignedHeaders=host\"},{\"id\":3,\"name\":\"Манты\",\"price\":2.75,\"description\":\"С мясом молодых бычков\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Manty&Kundyumy.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T125913Z&X-Amz-Expires=2592000&X-Amz-Signature=E8A9224CFB7A97EDFAC495A7E130DDCA0FF25DE42BAA50912A844CAEFD5C5CD5&X-Amz-SignedHeaders=host\"},{\"id\":4,\"name\":\"Буузы\",\"price\":4,\"description\":\"С телятиной и луком\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Buuzy&Kurze.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T123942Z&X-Amz-Expires=2592000&X-Amz-Signature=CA33CC7E69BF85FCCF10B5F4ADE13A65538A4A8D6C24BC560C33150332008E10&X-Amz-SignedHeaders=host\"},{\"id\":5,\"name\":\"Цзяоцзы\",\"price\":7.25,\"description\":\"С говядиной и свининой\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Czyaoczy.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T123247Z&X-Amz-Expires=2592000&X-Amz-Signature=92B8E79A1B3029B250D2DE7F9FF003792D5F7875597D3BDDD205223858C1B826&X-Amz-SignedHeaders=host\"},{\"id\":6,\"name\":\"Гедза\",\"price\":3.5,\"description\":\"С соевым мясом\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Gedza&Boraki.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T124641Z&X-Amz-Expires=2592000&X-Amz-Signature=EF117D235154E1C4FE71BF8D41D2445BEB90FB8C634CAF03755539D7BF4C6319&X-Amz-SignedHeaders=host\"},{\"id\":7,\"name\":\"Дим-самы\",\"price\":2.65,\"description\":\"С уткой\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Dim-samy&Ravioli.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T123647Z&X-Amz-Expires=2592000&X-Amz-Signature=DA0B5CEA515631B2873F6CD7B323FCC8C3D4103E28E39127906466D82F02B9CA&X-Amz-SignedHeaders=host\"},{\"id\":8,\"name\":\"Момо\",\"price\":5,\"description\":\"С бараниной\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Momo.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T130231Z&X-Amz-Expires=2592000&X-Amz-Signature=1085BB8880CE30968D35ABF5252365ACFB5F2D2FDC458124E3715086760C7089&X-Amz-SignedHeaders=host\"},{\"id\":9,\"name\":\"Вонтоны\",\"price\":4.1,\"description\":\"С креветками\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Pel%27meni&Vontony.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T130328Z&X-Amz-Expires=2592000&X-Amz-Signature=CE4BAE513B89EE16C8BAA019D783A7508350BC34C6D1B03FB605FCD0A1CEDF46&X-Amz-SignedHeaders=host\"},{\"id\":10,\"name\":\"Баоцзы\",\"price\":4.2,\"description\":\"С капустой\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Hinkali&Baoczy.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T125035Z&X-Amz-Expires=2592000&X-Amz-Signature=47634DAF429B9CF1DFA236884F530DD654DC76D4E14D48CD56BCCEA63F7BFDB0&X-Amz-SignedHeaders=host\"},{\"id\":11,\"name\":\"Кундюмы\",\"price\":5.45,\"description\":\"С грибами\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Manty&Kundyumy.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T125913Z&X-Amz-Expires=2592000&X-Amz-Signature=E8A9224CFB7A97EDFAC495A7E130DDCA0FF25DE42BAA50912A844CAEFD5C5CD5&X-Amz-SignedHeaders=host\"},{\"id\":12,\"name\":\"Курзе\",\"price\":3.25,\"description\":\"С крабом\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Buuzy&Kurze.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T123942Z&X-Amz-Expires=2592000&X-Amz-Signature=CA33CC7E69BF85FCCF10B5F4ADE13A65538A4A8D6C24BC560C33150332008E10&X-Amz-SignedHeaders=host\"},{\"id\":13,\"name\":\"Бораки\",\"price\":4,\"description\":\"С говядиной и бараниной\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Gedza&Boraki.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T124641Z&X-Amz-Expires=2592000&X-Amz-Signature=EF117D235154E1C4FE71BF8D41D2445BEB90FB8C634CAF03755539D7BF4C6319&X-Amz-SignedHeaders=host\"},{\"id\":14,\"name\":\"Равиоли\",\"price\":2.9,\"description\":\"С рикоттой\",\"image\":\"https://storage.yandexcloud.net/toxarey.bucket/Dim-samy&Ravioli.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T123647Z&X-Amz-Expires=2592000&X-Amz-Signature=DA0B5CEA515631B2873F6CD7B323FCC8C3D4103E28E39127906466D82F02B9CA&X-Amz-SignedHeaders=host\"}]}\n"

		assert.NoError(t, err)
		assert.JSONEq(t, string(expectedJSON), w.Body.String())
	})

	t.Run("healthcheck", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()
		app.HealthcheckController(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
