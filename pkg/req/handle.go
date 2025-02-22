package req

import (
	"net/http"
	"server/pkg/resp"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		resp.ResponceJson(*w, 402, err.Error())
		return nil, err
	}

	//валидация данных с помощью валидатора
	//---------------------------------------------
	err = IsValid(body)
	if err != nil {
		resp.ResponceJson(*w, 402, err.Error())
		return nil, err
	}
	//---------------------------------------------
	return &body, nil
}
