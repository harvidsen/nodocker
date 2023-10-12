import uvicorn
from fastapi import FastAPI
import subprocess


app = FastAPI(title="Et web api")

@app.get("/")
async def root():
    return {"Hello": "World"}


@app.get("/cowsay")
async def cowsay():
    proc = subprocess.run(["/bin/cowsay", "Hello World"], capture_output=True)
    return proc.stdout.decode()


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
