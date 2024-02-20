FROM python:3.12

WORKDIR /root/

COPY requirements.txt ./requirements.txt
COPY main.py function.py ./

COPY openapi.json asyncapi.json ./

RUN pip install -r requirements.txt

EXPOSE 8080

CMD ["python", "main.py"]