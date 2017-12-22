package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
	"github.com/mohuishou/scujwc-go"
	"strconv"
)

type Server struct {
	*socketio.Server
}

func NewServer(transportNames []string) (*Server, error) {
	ret, err := socketio.NewServer(transportNames)
	if err != nil {
		return nil, err
	}
	return &Server{ret}, nil
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Header)
	OriginList := r.Header["Origin"]
	Origin := ""
	if len(OriginList) > 0 {
		Origin = OriginList[0]
	}
	w.Header().Add("Access-Control-Allow-Origin", Origin)
	//w.Header().Add("Access-Control-Allow-Origin", "http://localhost:63342")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	s.Server.ServeHTTP(w, r)
}

func main() {
	server, err := NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		var jwc scujwc.Jwc
		so.On("login", func(params map[string]string) {
			id,ok := params["student_id"]
			uid,err := strconv.Atoi(id)
			if !ok || err!=nil{
				so.Emit("login",map[string]string{
					"status": "1",
					"msg": "请输入正确的用户名或密码",
				})
				return
			}
			password,ok := params["password"]
			if !ok{
				so.Emit("login",map[string]string{
					"status": "1",
					"msg": "请输入正确的用户名或密码",
				})
				return
			}
			jwc,err = scujwc.NewJwc(uid,password)
			if err != nil {
				so.Emit("login",map[string]string{
					"status": "1",
					"msg": err.Error(),
				})
				return
			}
			so.Emit("login",map[string]string{
				"status": "0",
				"msg": "登录成功",
			})
		})
		log.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}