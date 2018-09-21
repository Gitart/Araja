/****************************************************************************************
 *  ETRACK CLIENT MONITOR
 *  Copyright Etrack  (2018)
 *  12-09-2018 
 *  Ver 0.001 
 ****************************************************************************************/
package main


import (
	"net/http"
	"fmt"
	"os"
	"byte"
)

var(
    ipserv="http://118.223.111.231:15898"
)



/****************************************************************************************
 * DATETIME         : 06-09-2018 17:44
 * NOTES            : Send log information to server
 ****************************************************************************************/
func main() {
     p:=ipserv+"TEST*worker*test*Info*ccc34554zxcxzcaddfdf3445cvdv*acc12344"
     resp, err := http.Post(p,"",nil)
    
    if err != nil {
       return 
    }
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
	data := []byte(`{"foo":"bar"}`)
    r := bytes.NewReader(data)
    resp, err := http.Post("http://example.com/upload", "application/json", r)
    
    if err != nil {
       return 
    }
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
    ipserv := "http://118.223.111.231:51898"
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
