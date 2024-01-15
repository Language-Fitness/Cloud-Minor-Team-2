FROM python:3.10.6-slim

WORKDIR /app
COPY app/ .

RUN apt-get update && apt-get install -y libmagic1
RUN pip install --no-cache-dir -r requirements.txt


EXPOSE 4000

CMD ["python", "./main.py"]