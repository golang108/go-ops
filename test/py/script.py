

data = \
{
"name": "脚本名称",
"content": "脚本内容",
"args": {
"property1": "",
"property2": "string"
},
"desc": "我是描述",
"type": "shell",
"creater": "lxw"
}


import ops

ospReq = ops.Ops("http://127.0.0.1:8199")

res = ospReq.script_add(data)

print(res)
print("\n")


# res = ospReq.script_query({})

# print(res)
# print("\n")

# rlist = res['list']

# item0 = res['list'][0]

# item0['name'] = "hkjhjkh"
# item0['scriptId'] = item0["scriptUid"]

# res = ospReq.script_update(item0)
# print(res)
# print("\n")


# res = ospReq.script_query({"name":"hkjhjkh"})

# print(res)
# print("\n")


# uids = []
# for item in rlist:
#     uids.append(item['scriptUid'])

# res = ospReq.script_delete({"scriptIds":uids})

# print(res)
# print("\n")
