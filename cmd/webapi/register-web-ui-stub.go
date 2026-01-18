//go:build !webui

//заглушка, чтобы просто возвращать хэндлер без изменений
//чтобы сервер работал только как апи
package main

import (
	"net/http"
)

func registerWebUI(hdl http.Handler) (http.Handler, error) {
	return hdl, nil
}
