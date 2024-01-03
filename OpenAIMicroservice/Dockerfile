FROM python:3.10.6-slim

WORKDIR /App
COPY App/ .

RUN apt-get update && apt-get install -y libmagic1
RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 4000

CMD ["python", "./Main.py"]