import pika
import json
import os

from function import {{FUNCTION_NAME}}

def listen_to_rabbitmq():

    host = os.environ.get('RABBITMQ_HOST', 'localhost')
    port = int(os.environ.get('RABBITMQ_PORT', '5672'))
    username = os.environ.get('RABBITMQ_USERNAME', 'guest')
    password = os.environ.get('RABBITMQ_PASSWORD', 'guest')

    connection = pika.BlockingConnection(
        pika.ConnectionParameters(host=host, port=port, credentials=pika.PlainCredentials(username, password)),
    )
    channel = connection.channel()
    queue = os.environ.get("QUEUE", "rpc_queue")  # Use 'rpc_queue' consistently here
    channel.queue_declare(queue=queue, durable=True, auto_delete=False)


    def on_request(ch, method, props, body):
        rb = json.loads(body.decode('utf-8'))
        print(rb)
        response = {{FUNCTION_CALL}}

        ch.basic_publish(
            exchange="",
            routing_key=props.reply_to,
            properties=pika.BasicProperties(correlation_id=props.correlation_id),
            body=str(response),
        )
        ch.basic_ack(delivery_tag=method.delivery_tag)


    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(queue=queue, on_message_callback=on_request)  # Use 'rpc_queue' consistently here
    channel.start_consuming()

if __name__ == '__main__':
    listen_to_rabbitmq()

