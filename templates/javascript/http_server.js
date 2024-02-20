const express = require('express');
const fs = require('fs');

const app = express();
app.use(express.json());

function handleHttp() {
    app.post('/', (req, res) => {
        const rb = req.body;
        // Call function
        const value = ""; // Placeholder for function call
        if (typeof value === 'object') {
            res.json(value);
        } else {
            res.type('text/plain').send(String(value));
        }
    });

    app.get('/openapi', (req, res) => {
        fs.readFile('/root/openapi.json', 'utf8', (err, data) => {
            if (err) {
                res.status(500).json({ error: 'Failed to read JSON file' });
                return;
            }
            res.setHeader('Content-Type', 'application/json');
            res.status(200).send(data);
        });
    });

    app.get('/asyncapi', (req, res) => {
        fs.readFile('/root/asyncapi.json', 'utf8', (err, data) => {
            if (err) {
                res.status(500).json({ error: 'Failed to read JSON file' });
                return;
            }
            res.setHeader('Content-Type', 'application/json');
            res.status(200).send(data);
        });
    });

    app.listen(8080, () => {
        console.log('HTTP server is running on port 8080');
    });
}

module.exports = { handleHttp };
