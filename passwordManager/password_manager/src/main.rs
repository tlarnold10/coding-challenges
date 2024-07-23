use axum::{
    routing::{get},
    // http::StatusCode,
    // Json,
    Router,
};
// use serde::{Deserialize, Serialize};

#[tokio::main]
async fn main() {
    println!("Hello, world!");
    let app = Router::new().route("/", get(index));

    // run server
    let listener = tokio::net::TcpListener::bind("127.0.0.1:8675")
        .await
        .unwrap();

    println!("listing on {}", listener.local_addr().unwrap());
    axum::serve(listener, app).await.unwrap();
}

async fn index() -> &'static str {
    return "Hello world";
}
