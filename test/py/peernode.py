
import ops

ospReq = ops.Ops("http://127.0.0.1:8199")

res = ospReq.peer_nodes({})

print(res.json())