FROM rust:alpine

WORKDIR /app

COPY src ./src
COPY Cargo.toml .
COPY Cargo.lock .

CMD [ "cargo", "run" ]
