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
		log.Println("on connection")

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

			evaluateList,err:=jwc.GetEvaList()
			if err !=nil {
				so.Emit("getEvaluateList", map[string]string{
					"status": "1",
					"msg": err.Error(),
				})
				return
			}
			so.Emit("getEvaluateList",map[string]interface{}{
				"status": "0",
				"data": evaluateList,
			})
		})

		so.On("evaluate", func(evaluate scujwc.Evaluate) {
			err := jwc.Evaluate(&evaluate)
			if err !=nil {
				so.Emit("evaluate", map[string]string{
					"status": "1",
					"msg": err.Error(),
				})
				return
			}
			so.Emit("evaluate",map[string]interface{}{
				"status": "0",
				"data": evaluate,
				"msg": evaluate.CourseName +"-"+evaluate.TeacherName + ": 评教成功！",
			})
		})

		so.On("disconnection", func() {
			jwc.Logout()
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}