mod templates;

use axum::{
    routing::{ get, post },
    http::StatusCode,
    Json,
    Router,
    response::{ IntoResponse, Html },
    extract::Form
};
// use serde::{ Deserialize };
use templates::{ Index, PasswordItem };
use askama::Template;
use std::env;
use rusqlite::{ Connection, Result };

#[tokio::main]
async fn main() {
    // connecting to sqlite
    let conn = Connection::open("passwords.db");

    let app = Router::new()
        .route("/", get(index))
        .route("/password", post(create_password));
    
    // run server
    let listener = tokio::net::TcpListener::bind("127.0.0.1:8675")
        .await
        .unwrap();

    println!("listing on {}", listener.local_addr().unwrap());
    axum::serve(listener, app).await.unwrap();
}

async fn index() -> impl IntoResponse {
    let mut vec: Vec<PasswordItem> = Vec::new();
    vec.push({ PasswordItem { 
        account: "test".to_string(),
        password: "thing".to_string()
    }});

    vec.push({ PasswordItem {
        account: "another one".to_string(),
        password: "wouldn't you like to know!".to_string()
    }});
    let template = Index { name: "trevor".to_string(), passwords: vec };
    match template.render() {
        Ok(html) => Html(html).into_response(),
        Err(err) => (
            StatusCode::INTERNAL_SERVER_ERROR,
            format!("Failed to render template. Error: {err}"),
        ).into_response(),
    }
}

async fn create_password(Form(payload): Form<PasswordItem>) -> StatusCode {
    println!("payload information: {} {}", payload.account, payload.password);
    let password = PasswordItem {
        account: "test".to_string(),
        password: "more testing".to_string()
    };

    StatusCode::CREATED
}
