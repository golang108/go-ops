package consts

const (
	CheckItemTypeScript = "script"

	OpenAPITitle       = `Go-ops Server`
	OpenAPIDescription = `这是go-ops server的swagger文档. Enjoy 💖 `

	OpenAPISign = `<h1>请求头</h1>
	<pre class="codepre">
	{
		"GO-OPS-X-APPID":"你的appid",
		"GO-OPS-X-SIGNATURE":"签名值",
		"GO-OPS-X-TIMESTAMP":"时间戳",
		"GO-OPS-X-NONCE":"请求唯一id"
	}
	</pre>
	<h1>签名</h1>
	<p>使用sha1作为签名算法,seckey作为签名使用key,依次把apikey、nonce、timestame、body 写入签名内容中,最后对签名结果进行base64转码</p>
					<code>Golang</code>
					<pre class="codepre">
func GetSign(apikey string, seckey string, nonce, timestamp string, body []byte) (r string) {
	key := []byte(seckey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(apikey))
	mac.Write([]byte(nonce))
	mac.Write([]byte(timestamp))
	mac.Write(body)
	r = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return
}
					</pre>
					`
)
