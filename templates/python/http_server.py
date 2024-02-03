from flask import Flask, request, jsonify

from function import myFunc

app = Flask(__name__)

def run_server():
    app.run(port=8080)

@app.route('/')
def handler1():

    input=request.json
    myFunc()
    response = {{FUNCTION_NAME}}
    return jsonify(response)

if __name__ == '__main__':
    run_server()