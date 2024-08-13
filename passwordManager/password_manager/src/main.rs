mod templates;

use axum::{
    routing::{ get, post },
    http::StatusCode,
    Json,
    Router,
    response::{ IntoResponse, Html },
    extract::{ Form, Extension }
};
// use serde::{ Deserialize };
use templates::{ Index, PasswordItem, PasswordItemSnippet };
use askama::Template;
use std::env;
use rusqlite::{ Connection, Result, Statement };
use lazy_static::lazy_static;
use tokio::sync::Mutex;

#[tokio::main]
async fn main() {
    // connecting to sqlite
    // let conn = Connection::open("passwords.db");

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

// lazy static is a macro for static variables that are defined at runtime and
// persist throughout the life of the program
lazy_static! {
    static ref DATA: Mutex<Vec<PasswordItem>> = Mutex::new(Vec::new());
    static ref CONN: Mutex<Connection> = Mutex::new(Connection::open("passwords.db").unwrap());
}

fn print_vec(items: &[PasswordItem]) -> () {
    let v_iter = items.iter();
    for item in v_iter {
        println!("password: {}, account: {}, username: {}", item.password, item.account, item.username);
    }
}

async fn index() -> impl IntoResponse {
    let mut vec: Vec<PasswordItem> = Vec::new();
    let mut passwords = DATA.lock().await;
    let query = "SELECT account, username, password FROM passwords";
    let conn = CONN.lock().await;
    let mut stmt = conn.prepare(&query).unwrap();
    let passwords_iter = stmt.query_map([], |row| {
        Ok( PasswordItem {
            account: row.get(0)?,
            username: row.get(1)?,
            password: row.get(2)?,
        })
    });

    for password in passwords_iter.unwrap() {
        let temp_password = password.unwrap();
        passwords.push({ PasswordItem {
            account: temp_password.account,
            username: temp_password.username,
            password: temp_password.password
        }})
    } 

    let template = Index { name: "trevor".to_string(), passwords: passwords.to_vec() };
    match template.render() {
        Ok(html) => Html(html).into_response(),
        Err(err) => (
            StatusCode::INTERNAL_SERVER_ERROR,
            format!("Failed to render template. Error: {err}"),
        ).into_response(),
    }
}

async fn create_password(Form(payload): Form<PasswordItem>) -> impl IntoResponse {
    println!("payload information: {} {} {}", payload.account, payload.password, payload.username);
    let mut passwords = DATA.lock().await;
    let password = payload.password;
    let account = payload.account;
    let username = payload.username;
    let password = PasswordItem {
        account: account.clone(),
        password: password.clone(),
        username: username.clone()
    };
    passwords.push(password.clone());
    print_vec(&passwords);
    let query = format!("INSERT INTO passwords VALUES ('{}', '{}', '{}')", 
                                    password.account.clone(), password.username.clone(), password.password.clone());
    CONN.lock().await.execute(&query, ()).unwrap();
    let added_item = format!("Account: {} Password: {} Username: {}", password.account, password.password, password.username);
    // return Html(&added_item).into_response();
    let template = PasswordItemSnippet { snippet: added_item };
    match template.render() {
        Ok(html) => Html(html).into_response(),
        Err(err) => (
            StatusCode::INTERNAL_SERVER_ERROR,
            format!("Failed to render template. Error: {err}"),
        ).into_response(),
    }
}