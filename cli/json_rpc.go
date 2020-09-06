package cli

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/rpc/v2"
    "github.com/gorilla/rpc/v2/json2"
)

type JsonRPC struct {
    port Port
    url string
    services []interface{}
}

func (j *JsonRPC) SetCli(c *Cli) {
    j.port.SetCli(c)
    c.NewFlag().
        Name("url").
        Usage("url of the JSON RPC service").
        StringVar(&j.url, "/rpc")
}

func (j *JsonRPC) WithServices(services ...interface{}) (jj *JsonRPC) {
    for ii, _ := range services {
        j.Add(services[ii])
    }
    
    return j
}

func (j *JsonRPC) Add(service interface{}) {
    j.services = append(j.services, service)
}

func (j JsonRPC) Run() (err error) {
    s := rpc.NewServer()
    s.RegisterCodec(json2.NewCodec(), "application/json")
    s.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")
    
    for ii, _ := range j.services {
        err = s.RegisterService(j.services[ii], "")
        if err != nil {
            return
        }
    }

    r := mux.NewRouter()
    r.Handle(j.url, s)
    
    return http.ListenAndServe(j.port.Localhost(), r)
}
