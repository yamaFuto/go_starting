package main

//1 semaphore
//稼働しているgoroutineの数を限定する
//稼働しているgoroutine以外を中断させることができる

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"golang.org/x/sync/semaphore"
// )

// //semaphore.NewWeightedで生成⇒semaphore.Weightedtypeになる⇒ポインタなのはコピー、を防ぐため、直接加工を促すため
// var s *semaphore.Weighted = semaphore.NewWeighted(1)

// func longProcess(ctx context.Context) {
// 	//Acquireでロックしている（ロックされたらfalseを返す
// 	//一つ分確保しようとする
// 	isAcquire := s.TryAcquire(1)
// 	if !isAcquire {
// 		fmt.Println("Could not get lock")
// 		return
// 	}
// 	// //s.Aquireでロックしている(lockされていた場合にはcontextで大気処理をする)
// 	// if err := s.Acquire(ctx, 1); err != nil{
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	//ロックを解除
// 	defer s.Release(1)
// 	fmt.Println("Wait...")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Done")
// }

// func main() {
// 	ctx := context.TODO()
// 	go longProcess(ctx)
// 	go longProcess(ctx)
// 	go longProcess(ctx)
// 	time.Sleep(2 * time.Second)
// 	go longProcess(ctx)
// 	time.Sleep(5 * time.Second)
// }

//2 go-ini
//config.iniをgo操作で読み込む

// import (
// 	"fmt"
// 	"gopkg.in/ini.v1"
// )

// type ConfigList struct {
// 	Port 			int
// 	DbName 		string
// 	SQLDriver string
// }

// var Config ConfigList

// //mainよりも先に実行される
// func init() {
// 	cfg, _ := ini.Load("config.ini")
// 	Config = ConfigList{
// 		Port: cfg.Section("web").Key("port").MustInt(),
// 		//defaultがexample.sql
// 		DbName: cfg.Section("db").Key("name").MustString("example.sql"),
// 		//defaultが空
// 		SQLDriver: cfg.Section("db").Key("driver").String(),
// 	}
// }

// func main() {
// 	fmt.Printf("%T %v\n", Config.Port, Config.Port)
// 	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
// 	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
// }

//3 talib
//株価のindicatorを計算するメソッドを持っている

// import (
// 	"fmt"
// 	"github.com/markcheno/go-quote"
// 	"github.com/markcheno/go-talib"
// )

// func main() {
// 	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
// 	fmt.Print(spy.CSV())
// 	rsi2 := talib.Rsi(spy.Close, 2)
// 	fmt.Println(rsi2)
// 	mva := talib.Ema(spy.Close, 14)
// 	fmt.Println(mva)
// }

//4 bitcoin real-time get data

// import (
// 	"log"
// 	"net/url"

// 	"github.com/gorilla/websocket"
// )
 
// type JsonRPC2 struct {
// 	Version string `json:"jsonrpc"`
// 	Method string `json:"method"`
// 	Params interface{} `json:"params"`
// 	Result interface{} `json:"result,omitempty"`
// 	Id *int `json:"id,omitempty"`
// }
// type SubscribeParams struct {
// 	Channel string `json:"channel"`
// }
 
// func main() {
// 	u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path: "/json-rpc"}
// 	log.Printf("connecting to %s", u.String())
 
// 	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		log.Fatal("dial:", err)
// 	}
// 	defer c.Close()
 
// 	if err := c.WriteJSON(&JsonRPC2{Version: "2.0", Method: "subscribe", Params: &SubscribeParams{"lightning_ticker_BTC_JPY"}}); err != nil {
// 		log.Fatal("subscribe:", err)
// 		return
// 	}
 
// 	for {
// 		message := new(JsonRPC2)
// 		if err := c.ReadJSON(message); err != nil {
// 			log.Println("read:", err)
// 			return
// 		}
 
// 		if message.Method == "channelMessage" {
// 			log.Println(message.Params);
// 		}
// 	}
// }