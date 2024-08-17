mod templates;

use axum::{
    routing::{ get, post },
    http::StatusCode,
    Router,
    response::{ IntoResponse, Html },
    extract::{ Form }
};
// use serde::{ Deserialize };
use templates::{ Index, PasswordItem, PasswordItemSnippet };
use askama::Template;
use std::{env, f32::consts::E, sync::MutexGuard};
use rusqlite::{ Connection, Result, Statement };
use lazy_static::lazy_static;
use tokio::sync::Mutex;
use std::io;
use magic_crypt::{ new_magic_crypt, MagicCryptTrait, MagicCrypt64 };

// lazy static is a macro for static variables that are defined at runtime and
// persist throughout the life of the program
lazy_static! {
    static ref DATA: Mutex<Vec<PasswordItem>> = Mutex::new(Vec::new());
    static ref CONN: Mutex<Connection> = Mutex::new(Connection::open("passwords.db").unwrap());
    static ref KEY: MagicCrypt64 = new_magic_crypt!("magickey", 64);
}

fn print_vec(items: &[PasswordItem]) -> () {
    let v_iter = items.iter();
    for item in v_iter {
        println!("password: {}, account: {}, username: {}", item.password, item.account, item.username);
    }
}

#[tokio::main]
async fn main() {
    // connecting to sqlite, obtain master password, and decrypt it
    let query = "SELECT password FROM passwords WHERE account = 'Master'";
    let crypt = KEY.clone();
    let conn = CONN.lock().await;
    let mut stmt = conn.prepare(&query).unwrap();
    let master: String = match stmt.query_row([], |row| row.get(0)) {
        Ok(value) => value,
        Err(e) => String::from("na") 
    };
    if (master.eq("na")) {
        create_master_password(&conn);
    }
    let mut logged_in = false;
    let mut input = String::new();
    let mut attempts: u32 = 0;
    let crypt = crypt.decrypt_base64_to_string(master.clone());
    let decrypted_master = match crypt {
        Ok(result) => result,
        Err(e) => panic!("Error in Magic Crypt {}", e)
    };

    // check for accurate master password input
    while (!logged_in) {
        attempts = attempts + 1;
        input.clear();
        if (attempts == 1) {
            println!("Enter master password: ");
        }
        else {
            println!("Try again: ");
        }
        io::stdin()
            .read_line(&mut input)
            .expect("Failed to read line");
        if (decrypted_master.eq(&input.trim())) {
            logged_in = true;
        }
        else {
            println!("Incorrect master password");
        }

        // if user is not able to enter the correct password 3 times, program errors out
        if (attempts > 2) {
            panic!("Incorrect master password entered 3 times... goodbye");
        }
    }
    
    // need to drop variables in order for database to be connected later
    drop(stmt);
    drop(conn);

    // setting up routes
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

// get endpoint for index
async fn index() -> impl IntoResponse {
    // read in all passwords from database and display them to the screen
    let crypt = KEY.clone();
    let mut passwords = DATA.lock().await;
    passwords.clear();
    let query = "SELECT account, username, password FROM passwords WHERE account != 'Master'";
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

        // decrypt the encrypted password so that the user could actually see it in the UI
        let crypt = crypt.decrypt_base64_to_string(temp_password.password.clone());
        let decrypted_password = match crypt {
            Ok(result) => result,
            Err(e) => panic!("Error in Magic Crypt {}", e)
        };
        
        // adding password to htmx list of passwords
        passwords.push({ PasswordItem {
            account: temp_password.account,
            username: temp_password.username,
            password: decrypted_password.clone()
        }})
    } 

    // index page will display the form and all the data
    let template = Index { name: "trevor".to_string(), passwords: passwords.to_vec() };
    match template.render() {
        Ok(html) => Html(html).into_response(),
        Err(err) => (
            StatusCode::INTERNAL_SERVER_ERROR,
            format!("Failed to render template. Error: {err}"),
        ).into_response(),
    }
}

// post endpoint for creating a password
async fn create_password(Form(payload): Form<PasswordItem>) -> impl IntoResponse {
    let mcrypt = KEY.clone();
    let mut passwords = DATA.lock().await;
    
    // encrypt password that will be written to the database
    let encrypted_password = mcrypt.encrypt_str_to_base64(&payload.password);
    let account = payload.account;
    let password = payload.password;
    let username = payload.username;
    let password = PasswordItem {
        account: account.clone(),
        password: password.clone(),
        username: username.clone()
    };

    // add password to list of passwords for the htmx UI
    passwords.push(password.clone());
    // print_vec(&passwords);
    let query = format!("INSERT INTO passwords VALUES ('{}', '{}', '{}')", 
                                    password.account.clone(), password.username.clone(), encrypted_password.clone());
    CONN.lock().await.execute(&query, ()).unwrap();
    println!("Password added");
    let added_item = format!("Account: {} Password: {} Username: {}", password.account, password.password, password.username);
    
    // need to wrap password element into html to be returned to UI template
    let template = PasswordItemSnippet { snippet: added_item };
    match template.render() {
        Ok(html) => Html(html).into_response(),
        Err(err) => (
            StatusCode::INTERNAL_SERVER_ERROR,
            format!("Failed to render template. Error: {err}"),
        ).into_response(),
    }
}

fn create_master_password(conn: &tokio::sync::MutexGuard<Connection>) {
    let mut input = String::new();
    println!("No master password currently exists, please create: ");
    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read line");

    // encrypt master password
    let mcrypt = KEY.clone();
    let encrypted_password = mcrypt.encrypt_str_to_base64(&input.trim());
    let query: String = format!("INSERT INTO passwords VALUES ('{}', '{}', '{}')", "Master", "", encrypted_password.clone());

    // write encrypted master password to the database
    conn.execute(&query, ()).unwrap();
    println!("Master password created, please restart application");
}