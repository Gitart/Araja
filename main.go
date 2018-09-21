/****************************************************************************************
 *  ESSENTIA CLIENT MONITOR
 *  Copyright ESSENTIA  (2018)
 *  12-09-2018 
 *  Ver 0.001 
 ****************************************************************************************/
package main


import (

	"net/http"
	"fmt"
	"os"
	"bytes"
        "io/ioutil"
)

var(
    ipserv="http://118.191.199.157:5898"
)

 

/****************************************************************************************
 * DATETIME         : 06-09-2018 17:44
 * NOTES            : Send log information to server
 ****************************************************************************************/
func main() {
     p:=ipserv+"/api/add/TEST*worker*test*Info*ccc34554zxcxzcaddfdf3445cvdv*acc12344*233"
     resp,err:=http.Post(p,"",nil)
    
    // defer resp.Close()
    if err != nil {
     fmt.Println("Err", err.Error())
    }

fmt.Print("What is your name? ")
    fmt.Println("Status...",                            resp.Status)
    fmt.Println("Status code...",                       resp.StatusCode)
    fmt.Println("Proto...",                             resp.Proto)
    fmt.Println("Proto Major...",                       resp.ProtoMajor)
    fmt.Println("Proto Minor...",                       resp.ProtoMinor)
    fmt.Println("Header...",                            resp.Header)
    fmt.Println("Body...",                              resp.Body)
    fmt.Println("ContentLength...",                     resp.ContentLength)
    fmt.Println("TransferEncoding []string...",         resp.TransferEncoding)
    fmt.Println("Close bool...",                        resp.Close)
    fmt.Println("Uncompressed bool...",                 resp.Uncompressed)
    fmt.Println("Trailer Header...",                    resp.Trailer)
    fmt.Println("\a")

    


	// Contains transfer encodings from outer-most to inner-most. Value is
	// nil, means that "identity" encoding is used.
	// TransferEncoding []string

	// Close records whether the header directed that the connection be
	// closed after reading Body. The value is advice for clients: neither
	// ReadResponse nor Response.Write ever closes a connection.
	// Close bool

	// Uncompressed reports whether the response was sent compressed but
	// was decompressed by the http package. When true, reading from
	// Body yields the uncompressed content instead of the compressed
	// content actually set from the server, ContentLength is set to -1,
	// and the "Content-Length" and "Content-Encoding" fields are deleted
	// from the responseHeader. To get the original response from
	// the server, set Transport.DisableCompression to true.
	// Uncompressed bool

	// Trailer maps trailer keys to values in the same
	// format as Header.
	//
	// The Trailer initially contains only nil values, one for
	// each key specified in the server's "Trailer" header
	// value. Those values are not added to Header.
	//
	// Trailer must not be accessed concurrently with Read calls
	// on the Body.
	//
	// After Body.Read has returned io.EOF, Trailer will contain
	// any trailer values sent by the server.
	// Trailer Header

	// Request is the request that was sent to obtain this Response.
	// Request's Body is nil (having already been consumed).
	// This is only populated for Client requests.
	//Request *Request

	// TLS contains information about the TLS connection on which the
	// response was received. It is nil for unencrypted responses.
	// The pointer is shared between responses and should not be
	// modified.
	//TLS *tls.ConnectionState


}

//************************************************************
//  Name    : Test  
//  Date    : 20018-09-21
//  Author  : Svachenko Arthur
//  Company : Essentia
//  Number  : 
//  Module  : 
//************************************************************
func test_Json() {
	// data := []byte(`{"foo":"bar"}`)
 //    r := bytes.NewReader(data)
 //    resp, err := http.Post("http://example.com/upload", "application/json", r)
    
 //    if err != nil {
 //       return 
 //    }
 //    defer resp.Close()
}



//************************************************************
//  Name    : Test 
//  Date    : is a snippet.
//  Author  : Svachenko Arthur
//  Company : Essentia
//  Number  : 
//  Module  : 
//************************************************************
func Send_test() {
	 // Args
     arg       :=os.Args
     Project   :=arg[1]
     Module    :=arg[2]
     Operation :=arg[3]
     Status    :=arg[4]
     BlockId   :=arg[5]
     AccountId :=arg[6]
     Send_Info(Project, Module, Operation, Status, BlockId,AccountId)
}


// *************************************************************
//  Send to log server information
//  Send_Info("EssentiaHybrid", "worker", "", "Info","ccc34554zxcxzcaddfdf3445cvdv","acc12344")
// *************************************************************
func Send_Info(Project, Module, Operation, Status, BlockId, AccountId string ){
    ipserv := "http://18.223.111.231:5898"
    url    :=  ipserv+"/api/add/"+Project+"*"+Module+"*"+Operation+"*"+Status+"*"+BlockId+"*"+AccountId
	re,errr:=  http.NewRequest("GET", url, nil)
	
	if errr!=nil{
       fmt.Println(errr.Error()) 
       return
	}

	res, erd := http.DefaultClient.Do(re)
	if erd!=nil{
       fmt.Println(erd.Error()) 
       return
	}

	defer res.Body.Close()
}



func test_003() {
	

url := "http://restapi3.apiary.io/notes"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

}
