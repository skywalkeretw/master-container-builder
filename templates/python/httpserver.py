from flask import Flask, request, jsonify

from function import {{FUNCTION_NAME}}

app = Flask(__name__)

def run_server():
    app.run(port=8080)

@app.route('/', methods=['POST'])
def handle_function():
    if request.is_json:        
        rb=request.json
        response = {{FUNCTION_CALL}}
        return jsonify(response)

@app.route('/openapi', methods=['GET'])
def get_openapi():
    try:
        with open('/root/openapi.json', 'r') as file:
            json_data = file.read()
            return jsonify(json_data)
    except Exception as e:
        return jsonify({"error": "Failed to read JSON file"}), 500

@app.route('/asyncapi', methods=['GET'])
def get_asyncapi():
    try:
        with open('/root/asyncapi.json', 'r') as file:
            json_data = file.read()
            return jsonify(json_data)
    except Exception as e:
        return jsonify({"error": "Failed to read JSON file"}), 500

if __name__ == '__main__':
    run_server()