from flask import Flask, request, jsonify

from function import myFunc

app = Flask(__name__)

def run_server():
    app.run(port=8080)

@app.route('/', methods=['POST'])
def handle_function():
    if request.is_json:
        
    rb=request.json
    response = {{FUNCTION_CALL}}
    return jsonify(response)

if __name__ == '__main__':
    run_server()