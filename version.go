package majsoul

type majsoulVersion struct {
	Code       string `json:"code"`    // code.js 路径 v0.10.154.w/code.js
	ResVersion string `json:"version"` // 资源文件版本  0.10.154.w（注意开头没有 v）
}

func GetMajsoulVersion(apiGetVersionURL string) (version *majsoulVersion, err error) {
	version = &majsoulVersion{}
	err = get(apiGetVersionURL, version)
	return
}
