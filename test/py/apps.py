
# data = \
# {
#   "name": "name1",
#   "owner": "owner1"
# }

import ops

ospReq = ops.Ops("http://82.157.165.187:30004")

for i in range(100,200):
  data = {}
  data['name'] = "name" + str(i)
  data['owner'] = "owner" + str(i)
  res = ospReq.app_add(data)

# print(res)
# print("\n")

print('finish')