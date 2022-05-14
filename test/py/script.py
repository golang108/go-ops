

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


appid = 'adeif00jps0cjq4dq4kf4ys100bqsgp7'
apikey = 'adeif00jps0cjq4dq4kf60t2009dn99jadeif00jps0cjq4dq4kf6am30070yp17'
seckey = 'adeif00jps0cjq4dq4kf6lx400u1hfkoadeif00jps0cjq4dq4kf6v1500ciqss7adeif00jps0cjq4dq4kf73p60007did2adeif00jps0cjq4dq4kf7c4700q1uhbp'


ospReq = ops.Ops("http://127.0.0.1:8199", appid, apikey, seckey)

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
