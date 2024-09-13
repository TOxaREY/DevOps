package dependencies

import (
	"gitlab.praktikum-services.ru/Stasyan/momo-store/internal/store/dumplings"
	"gitlab.praktikum-services.ru/Stasyan/momo-store/internal/store/dumplings/fake"
)

// NewFakeDumplingsStore returns new fake store for app
func NewFakeDumplingsStore() (dumplings.Store, error) {
	packs := []dumplings.Product{
		{
			ID:          1,
			Name:        "Пельмени",
			Description: "С говядиной",
			Price:       5.00,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Pel%27meni&Vontony.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T130328Z&X-Amz-Expires=2592000&X-Amz-Signature=CE4BAE513B89EE16C8BAA019D783A7508350BC34C6D1B03FB605FCD0A1CEDF46&X-Amz-SignedHeaders=host",
		},
		{
			ID:          2,
			Name:        "Хинкали",
			Description: "Со свининой",
			Price:       3.50,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Hinkali&Baoczy.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T125035Z&X-Amz-Expires=2592000&X-Amz-Signature=47634DAF429B9CF1DFA236884F530DD654DC76D4E14D48CD56BCCEA63F7BFDB0&X-Amz-SignedHeaders=host",
		},
		{
			ID:          3,
			Name:        "Манты",
			Description: "С мясом молодых бычков",
			Price:       2.75,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Manty&Kundyumy.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T125913Z&X-Amz-Expires=2592000&X-Amz-Signature=E8A9224CFB7A97EDFAC495A7E130DDCA0FF25DE42BAA50912A844CAEFD5C5CD5&X-Amz-SignedHeaders=host",
		},
		{
			ID:          4,
			Name:        "Буузы",
			Description: "С телятиной и луком",
			Price:       4.00,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Buuzy&Kurze.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T123942Z&X-Amz-Expires=2592000&X-Amz-Signature=CA33CC7E69BF85FCCF10B5F4ADE13A65538A4A8D6C24BC560C33150332008E10&X-Amz-SignedHeaders=host",
		},
		{
			ID:          5,
			Name:        "Цзяоцзы",
			Description: "С говядиной и свининой",
			Price:       7.25,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Czyaoczy.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T123247Z&X-Amz-Expires=2592000&X-Amz-Signature=92B8E79A1B3029B250D2DE7F9FF003792D5F7875597D3BDDD205223858C1B826&X-Amz-SignedHeaders=host",
		},
		{
			ID:          6,
			Name:        "Гедза",
			Description: "С соевым мясом",
			Price:       3.50,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Gedza&Boraki.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T124641Z&X-Amz-Expires=2592000&X-Amz-Signature=EF117D235154E1C4FE71BF8D41D2445BEB90FB8C634CAF03755539D7BF4C6319&X-Amz-SignedHeaders=host",
		},
		{
			ID:          7,
			Name:        "Дим-самы",
			Description: "С уткой",
			Price:       2.65,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Dim-samy&Ravioli.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T123647Z&X-Amz-Expires=2592000&X-Amz-Signature=DA0B5CEA515631B2873F6CD7B323FCC8C3D4103E28E39127906466D82F02B9CA&X-Amz-SignedHeaders=host",
		},
		{
			ID:          8,
			Name:        "Момо",
			Description: "С бараниной",
			Price:       5.00,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Momo.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T130231Z&X-Amz-Expires=2592000&X-Amz-Signature=1085BB8880CE30968D35ABF5252365ACFB5F2D2FDC458124E3715086760C7089&X-Amz-SignedHeaders=host",
		},
		{
			ID:          9,
			Name:        "Вонтоны",
			Description: "С креветками",
			Price:       4.10,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Pel%27meni&Vontony.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T130328Z&X-Amz-Expires=2592000&X-Amz-Signature=CE4BAE513B89EE16C8BAA019D783A7508350BC34C6D1B03FB605FCD0A1CEDF46&X-Amz-SignedHeaders=host",
		},
		{
			ID:          10,
			Name:        "Баоцзы",
			Description: "С капустой",
			Price:       4.20,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Hinkali&Baoczy.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T125035Z&X-Amz-Expires=2592000&X-Amz-Signature=47634DAF429B9CF1DFA236884F530DD654DC76D4E14D48CD56BCCEA63F7BFDB0&X-Amz-SignedHeaders=host",
		},
		{
			ID:          11,
			Name:        "Кундюмы",
			Description: "С грибами",
			Price:       5.45,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Manty&Kundyumy.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T125913Z&X-Amz-Expires=2592000&X-Amz-Signature=E8A9224CFB7A97EDFAC495A7E130DDCA0FF25DE42BAA50912A844CAEFD5C5CD5&X-Amz-SignedHeaders=host",
		},
		{
			ID:          12,
			Name:        "Курзе",
			Description: "С крабом",
			Price:       3.25,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Buuzy&Kurze.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T123942Z&X-Amz-Expires=2592000&X-Amz-Signature=CA33CC7E69BF85FCCF10B5F4ADE13A65538A4A8D6C24BC560C33150332008E10&X-Amz-SignedHeaders=host",
		},
		{
			ID:          13,
			Name:        "Бораки",
			Description: "С говядиной и бараниной",
			Price:       4.00,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Gedza&Boraki.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T124641Z&X-Amz-Expires=2592000&X-Amz-Signature=EF117D235154E1C4FE71BF8D41D2445BEB90FB8C634CAF03755539D7BF4C6319&X-Amz-SignedHeaders=host",
		},
		{
			ID:          14,
			Name:        "Равиоли",
			Description: "С рикоттой",
			Price:       2.90,
			Image:       "https://storage.yandexcloud.net/toxarey.bucket/Dim-samy&Ravioli.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEQ4PHQ2teqFmsVOSA0eFU%2F20240826%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240826T123647Z&X-Amz-Expires=2592000&X-Amz-Signature=DA0B5CEA515631B2873F6CD7B323FCC8C3D4103E28E39127906466D82F02B9CA&X-Amz-SignedHeaders=host",
		},
	}

	store := fake.NewStore()
	store.SetAvailablePacks(packs...)

	return store, nil
}
