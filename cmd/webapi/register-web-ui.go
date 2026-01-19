//go:build webui
// тут фронт становится частью гоу программы 
//тут добавляем поддержку хэндлера,если путь начинается с /dashboard/, он отдаёт статические файлы из embedded dist,
//  а все остальные запросы передаёт в API router
package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/aisforaru/AruzhanRakhmzhan_wasa/webui"
)

func registerWebUI(hdl http.Handler) (http.Handler, error) {
	distDirectory, err := fs.Sub(webui.Dist, "dist")
	if err != nil {
		return nil, fmt.Errorf("error embedding WebUI dist/ directory: %w", err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.RequestURI, "/dashboard/") {
			http.StripPrefix("/dashboard/", http.FileServer(http.FS(distDirectory))).ServeHTTP(w, r)
			return
		}
		hdl.ServeHTTP(w, r)
	}), nil
}
