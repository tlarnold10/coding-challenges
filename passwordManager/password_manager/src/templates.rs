use askama::Template;
use serde::{ Deserialize };

#[derive(Deserialize)]
pub struct PasswordItem {
    pub account: String,
    pub password: String
}

#[derive(Template)]
#[template(path = "index.html")]
pub struct Index {
    pub name: String,
    pub passwords: Vec<PasswordItem>
}
