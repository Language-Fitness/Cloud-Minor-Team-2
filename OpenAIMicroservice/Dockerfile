FROM python:3.10.6-slim

WORKDIR /App
COPY App/ .

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 80

CMD ["python", "./Main.py"]