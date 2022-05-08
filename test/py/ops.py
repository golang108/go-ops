
import json
import requests


class Ops:
    host = ''

    def __init__(self, host, username='test', password='test'):
        self.host = host
        token = self.login({'username': username, 'passwd': password})
        self.token = "Bearer " + token

    def post(self, address, data, isLogin=False):
        url = self.host + address
        dataStr = json.dumps(data)

        header = {
            'Content-Type': 'application/json',
        }
        if not isLogin:
            header['GO-OPS-X-TOKEN'] = self.token
        res = requests.post(url=url, headers=header, data=dataStr)
        return self.getData(res)

    def put(self, address, data, isLogin=False):
        url = self.host + address
        dataStr = json.dumps(data)

        header = {
            'Content-Type': 'application/json'
        }
        if not isLogin:
            header['GO-OPS-X-TOKEN'] = self.token
        res = requests.put(url=url, headers=header, data=dataStr)
        return self.getData(res)

    def login(self, data) -> str:
        res = self.post('/user/login', data, True)
        token = res['token']
        return token

    def getData(self, data):
        res = data.json()
        if res['code'] == 0:
            return res['data']
        return res

    def peer_nodes(self, data):
        return self.post("/peer/nodes", data)

    def script_add(self, data):
        return self.post("/v1/m/script", data)

    def script_update(self, data):
        return self.put("/v1/m/script", data)

    def script_query(self, data):
        return self.post("/v1/m/script/query", data)

    def script_delete(self, data):
        return self.post("/v1/m/script/delete", data)

    def task_preset_add(self, data):
        return self.post("/v1/m/task/preset/create", data)

    def task_preset_update(self, data):
        return self.post("/v1/m/task/preset/update", data)

    def task_preset_query(self, data):
        return self.post("/v1/m/task/preset/query", data)

    def task_preset_delete(self, data):
        return self.post("/v1/m/task/preset/delete", data)

    def app_add(self, data):
        return self.post("/v1/m/app/create", data)
