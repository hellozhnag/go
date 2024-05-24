import requests

rsp=requests.post("http://localhost:4321/jsonrpc",json={
    "id":0,
    "params":["bobby"],
    "method":"HelloService.Hello"
})

print(rsp.text) # {"id":0,"result":"hello~bobby","error":null}