import pika
import json
import os

def listen_to_rabbitmq():

    host = os.environ.get('RABBITMQ_HOST', 'localhost')
    port = int(os.environ.get('RABBITMQ_PORT', '5672'))
    username = os.environ.get('RABBITMQ_USERNAME', 'guest')
    password = os.environ.get('RABBITMQ_PASSWORD', 'guest')

    connection = pika.BlockingConnection(
        pika.ConnectionParameters(host=host, port=port, credentials=pika.PlainCredentials(username, password)),
    )
    channel = connection.channel()
    queue = os.environ.get("QUEUE", "default")
    channel.queue_declare(queue=queue)

    def on_request(ch, method, props, body):
        rb = json.loads(body.decode('utf-8'))

        response = {{FUNCTION_CALL}}

        ch.basic_publish(
            exchange="",
            routing_key=props.reply_to,
            properties=pika.BasicProperties(correlation_id=props.correlation_id),
            body=str(response),
        )
        ch.basic_ack(delivery_tag=method.delivery_tag)


    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(queue="rpc_queue", on_message_callback=on_request)
    channel.start_consuming()

if __name__ == '__main__':
    listen_to_rabbitmq()

