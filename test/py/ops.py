
import base64
import hashlib
import hmac
import json
import uuid
import requests
import time


class Ops:
    apikey = ''
    seckey = ''
    host = ''
    appid = ''

    def __init__(self, host, appid, apikey, seckey):
        self.host = host
        self.appid = appid
        self.apikey = apikey
        self.seckey = seckey
        

    def getSign(self, nonce, timestamp, body):
        mac = hmac.new(bytes(self.seckey, encoding='utf-8'), None, hashlib.sha1)
        mac.update(bytes(self.apikey, encoding='utf-8'))
        mac.update(bytes(nonce, encoding='utf-8'))
        mac.update(bytes(str(timestamp), encoding='utf-8'))
        mac.update(bytes(body, encoding='utf-8'))
        return base64.b64encode(mac.digest()).decode()

    def request(self, method, url, data):
        method = method.upper()
        nonce = str(uuid.uuid1())
        timestamp = str(int(time.time()))
        body = json.dumps(data)
        sign = self.getSign(nonce, timestamp, body)

        header = {
            'Content-Type': 'application/json',
            'GO-OPS-X-TIMESTAMP': timestamp,
            'GO-OPS-X-NONCE': nonce,
            'GO-OPS-X-SIGNATURE': sign,
            'GO-OPS-X-APPID': self.appid
        }
        print("header:", header)
        if method == 'POST':
            return requests.post(url=url, headers=header, data=body)
        elif method == 'PUT':
            return requests.put(url=url, headers=header, data=body)
        return {}


    def post(self, address, data):
        url = self.host + address
        res = self.request('POST', url, data)
        if res.status_code == 200:
            return self.getData(res)
        print("res:", res)
        print("url:", url)
        return {}

    def put(self, address, data, isLogin=False):
        url = self.host + address
        res = self.request('PUT', url, data)
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
        return self.post("/v1/m/script/add", data)

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
