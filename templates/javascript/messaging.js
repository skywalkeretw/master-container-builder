const amqp = require('amqplib/callback_api');
const { getEnvBool } = require('./utils');

function handleMessaging() {
    amqp.connect('amqp://localhost', function(error0, connection) {
        if (error0) {
            throw error0;
        }
        connection.createChannel(function(error1, channel) {
            if (error1) {
                throw error1;
            }
            const queue = 'rpc_queue';
            channel.assertQueue(queue, {
                durable: false
            });
            channel.prefetch(1);
            console.log(" [x] Awaiting RPC requests");

            channel.consume(queue, function reply(msg) {
                const requestData = JSON.parse(msg.content.toString());
                console.log("Received request:", requestData);

                const responseData = {{FUNCTION_CALL}} // Placeholder for response data

                channel.sendToQueue(msg.properties.replyTo,
                    Buffer.from(JSON.stringify(responseData)), {
                        correlationId: msg.properties.correlationId
                    });

                channel.ack(msg);
            });
        });
    });
}

module.exports = { handleMessaging };
