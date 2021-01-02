package handler

/**
编写http拦截，实现token鉴权逻辑（类似于Java中注解功能逻辑）
*/
import "net/http"

//HttpInterceptor:http拦截器
func HttpInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.Form.Get("username")
		token := r.Form.Get("token")
		if len(username) < 3 || !IsTokenValid(token, username) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h(w, r)
	})
}
