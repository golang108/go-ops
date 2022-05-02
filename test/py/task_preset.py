data = \
{
"type": "string",
"name": "预设任务",
"creater": "luxingwen",
"content": "我是内容hjkjjhkhk"
}


import ops

ospReq = ops.Ops("http://wwh.biggerforum.org:8199")

for num in range(1, 10):
    data["name"] = data["name"] + str(num)
    res = ospReq.task_preset_add(data)
    print(res)
    print('\n')

res = ospReq.task_preset_query({})
for item in res['list']:
    print(item)
    item["name"] = "更新+" + item["name"]
    res = ospReq.task_preset_update(item)
    print("update:", res)
    print("\n")

