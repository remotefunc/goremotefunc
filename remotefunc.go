package remotefunc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type RemoteFunc struct {
	*http.ServeMux
	Port string
}

func New() RemoteFunc {
	mux := http.NewServeMux()
	return RemoteFunc{
		mux,
		":3000",
	}
}

func (rf *RemoteFunc) Start() {
	log.Println("Starting RemoteFunc on port", rf.Port)
	log.Fatal(http.ListenAndServe(rf.Port, rf))
}

func (rf *RemoteFunc) AddFunc(name string, fun interface{}) {
	endpoint := fmt.Sprintf("/%s", name)
	rf.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		params := string(b)

		result := rf.callfunc(params, fun)
		w.Write([]byte(result))
	})
}

/*
callfunc calls a function with params from json string and returns a json string with the result
*/
func (rf *RemoteFunc) callfunc(jsonparams string, fun interface{}) string {
	funv := reflect.ValueOf(fun) //Function value
	funt := funv.Type()          //Function type
	funplen := funt.NumIn()

	//Reflect over the function type to get input values
	inputinterfaces := make([]interface{}, funplen)
	for i := 0; i < funplen; i++ {
		inputinterfaces[i] = reflect.New(funt.In(i)).Interface()
	}

	if len(inputinterfaces) > 0 {
		fromjson(jsonparams, &inputinterfaces)
	}

	//Make input from into values
	inputvalues := make([]reflect.Value, funplen)
	for i := 0; i < funplen; i++ {
		inputvalues[i] = reflect.ValueOf(inputinterfaces[i]).Elem()
	}

	//Call function
	resultvalues := funv.Call(inputvalues)

	//Return result
	if len(resultvalues) == 1 {
		return tojson(resultvalues[0].Interface())
	} else if len(resultvalues) > 1 {
		//More than one result, return array
		resultinterfaces := make([]interface{}, len(resultvalues))
		for i := 0; i < len(resultvalues); i++ {
			resultinterfaces[i] = resultvalues[i].Interface()
		}
		return tojson(resultinterfaces)
	} else {
		return ""
	}
}

func fromjson(j string, v interface{}) {
	err := json.Unmarshal([]byte(j), v)

	if err != nil {
		panic(err)
	}
}

func tojson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}
