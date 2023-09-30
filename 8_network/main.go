package main

//1 http

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// )

// func main() {
// 	// //urlに接続
// 	// resp, _ := http.Get("http://example.com")
// 	// defer resp.Body.Close()
// 	// body, _ := ioutil.ReadAll(resp.Body)
// 	// fmt.Println(string(body))

// 	//urlが正しいか判定する
// 	base, err := url.Parse("http://example.com")
// 	//baseの後につけるURL
// 	reference, _ := url.Parse("/test?a=1&b=2")
// 	endpoint := base.ResolveReference(reference).String()

// 	//header
// 	fmt.Println(endpoint)
// 	fmt.Println(base, err)

// 	//get,リクエスト作成
// 	req, _ := http.NewRequest("GET", endpoint, nil)

// 	// //post
// 	// //postはqueryをurlで見られないように第三引数(body)に置く
// 	// //NewBufferは抽象的化した読み書き用のメソッドをそろえたメモリ領域
// 	// req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byete("password")))

// 	req.Header.Add("IF-None-Match", `W/"wizzy"`)

// 	//query
// 	q := req.URL.Query()
// 	q.Add("c", "3&%")
// 	fmt.Println(q)
// 	//encodeすることによって&%などをビット化する
// 	fmt.Println(q.Encode())
// 	req.URL.RawQuery = q.Encode()

// 	//送信
// 	//ポインタ→意向で初期構造体が関数上で書き換えられやすくするため
// 	var client *http.Client = &http.Client{}
// 	resp, _ := client.Do(req)
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(body))
// }

//2 json

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type T struct{}

// type Person struct {
// 	//json化(Marshal)するときの規則を定義する
// 	//json:"age,string"⇒string化
// 	//json: "_"⇒非表示
// 	//json:"age,omitempty⇒０、空の時は見せない
// 	//json:`"T, omitempty"`⇒ポインタ型でなくてはいけない(json検査ではpointを初期化する前にメモリを探しに行くためnilが返ってくる)
// 	//Tとした場合はTを補完的に{}で埋めてしまうからnilと整合性が取れない
// 	Name 			string   `json:"name"`
// 	Age 			int			 `json:"age"`
// 	Nicknames []string `json:"nicknames"`
// 	T 				*T			 `json:"T,omitempty"`
// }

// //独自のmarshalを生成
// //MarshalJSONにしなくてはいけない
// //json.Marshalが呼ばれると自動的に呼ばれる(引数のstructと紐づけることによって、引数のstructによってそれぞれ違うmarshalを作ることができる)
// func (p Person) MarshalJSON() ([]byte, error) {
// 		// ストラクトを初期化している
// 	v, err := json.Marshal(&struct{
// 		Name string
// 	}{
// 		Name: "Mr." + p.Name,
// 	})
// 	return v, err
// }

// //marshal(unmarshal)されたら、まず、その関数がmarshal(unmarshal)JSONを引数のstructが保持していないか探し、あったらそれを実行する
// //第二引数のstructと連動させることによって、それぞれのsrructが第二引数で呼ばれたときに呼ばれる
// //書き換えが必要なためポインタを渡している
// func (p *Person) UnmarshalJSON(b []byte) error {
// 	type Person2 struct{
// 		Name string
// 	}
// 	var p2 Person2
// 	err := json.Unmarshal(b, &p2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	p.Name = p2.Name + "!"
// 	return err
// }

// func main() {
// 	b := []byte(`{"name":"mike", "age":20, "nicknames":["a", "b", "c"]}`)
// 	var p Person

// 	//ネットワークから帰ってきたバイト配列をストラクトにjsonとbyteを変換して第二引数のストラクトに合うように値を格納している
// 	if err := json.Unmarshal(b, &p); err != nil{
// 		fmt.Println(err)
// 	}
// 	fmt.Println(p.Name,p.Age, p.Nicknames)

// 	//byte[]jsonに変換⇒byte型で返ってくる
// 	v, _ := json.Marshal(&p)
// 	fmt.Println(string(v))
// }

//3 hmac
//APIでサーバにアクセスする際のauthenticationとしてheaderに含める

// import (
// 	"crypto/hmac"
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"fmt"
// )

// var DB = map[string]string{
// 	"User1Key": "User1Secret",
// 	"User2Key": "User2Secret",
// }

// // Server側で照合
// func Server(apiKey, sign string, data []byte) {
// 	apiSecret := DB[apiKey]
// 	h := hmac.New(sha256.New, []byte(apiSecret))
// 	h.Write(data)
// 	expectedHMAC := hex.EncodeToString(h.Sum(nil))
// 	fmt.Println(sign == expectedHMAC)
// }

// func main() {
// 	const apiKey = "User1Key"
// 	const apiSecret = "User1Secret"

// 	//sha256暗号アルゴリズムをもとにapiKey1をハッシュキー(バイトの配列のちにstring型にしてhashに直す)に変換している
// 	data := []byte("data")
// 	h := hmac.New(sha256.New, []byte(apiSecret))
// 	//dataを書き込む
// 	h.Write(data)
// 	//hexでbyte⇒stringにしてから末尾のnilを付け加える
// 	sign := hex.EncodeToString(h.Sum(nil))
// 	fmt.Println(sign)

// 	Server(apiKey, sign, data)
// }
