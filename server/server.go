package server

func Init() {
	r := NewRouter()
	_ = r.Run(":3000")
}
