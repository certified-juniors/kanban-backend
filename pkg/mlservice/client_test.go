package mlservice

import "ofs/internal/lib/logger/handlers/slogdiscard"

func CreateTestClient() INITProClient {

	return NewClient(&Options{
		Endpoint: EndpointTest,
		Logger:   slogdiscard.NewDiscardLogger(),
	})
}
