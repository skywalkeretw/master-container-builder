import pika


def listen_to_rabbitmq():


    connection = pika.BlockingConnection(
        pika.ConnectionParameters(host="localhost"),
    )
    channel = connection.channel()

    channel.queue_declare(queue="rpc_queue")

    def on_request(ch, method, props, body):
        n = int(body)

        print(f" [.] fib({n})")
        response = {{FUNCTION_NAME}}

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
