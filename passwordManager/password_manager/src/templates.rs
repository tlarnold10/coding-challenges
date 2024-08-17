use askama::Template;
use serde::{ Deserialize };
use std::fmt;

#[derive(Deserialize)]
#[derive(Clone)]
#[derive(Debug)]
pub struct PasswordItem {
    pub account: String,
    pub username: String,
    pub password: String
}

impl fmt::Display for PasswordItem {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "Account: {} Password: {} Username: {}", self.account, self.password, self.username)
    }
}

#[derive(Template)]
#[template(path = "index.html")]
pub struct Index {
    pub name: String,
    pub passwords: Vec<PasswordItem>
}

#[derive(Template)]
#[template(path = "password.html")]
pub struct PasswordItemSnippet {
    pub snippet: String
}
