import json
import socket

request={
    "id":0,
    "params":["bobby"],
    "method":"HelloService.Hello"
}
client=socket.create_connection(("localhost",4321))
client.sendall(json.dumps(request).encode())

rsp=client.recv(1024)
rsp=json.loads(rsp.decode())
print(rsp)//{'id': 0, 'result': 'hello~bobby', 'error': None}