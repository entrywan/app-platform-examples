import os
from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello_world():
    return b'hello, python!'

app.run(host='0.0.0.0')
