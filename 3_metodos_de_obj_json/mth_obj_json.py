# 1- object.keys

User = {
  "name": "User Name",
  "age": "User Age",
  "date": "27/01/26"
}

keys = User.keys()

print(f"Object.keys no caso keys do py", keys)

# json.loads e json loaf
# json.loads() ≈ decode / parse

# json.dumps() ≈ encode / stringify

import json

json_string = '{ "name": "Cássio", "working": "Dev", "year": 2026}'

parsed = json.loads(json_string)

print(parsed)