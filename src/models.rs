extern crate diesel;

use diesel::pg::PgConnection;
use diesel::prelude::*;
use std::env;

pub fn init_db() -> PgConnection {
  let db_host: String = env::var("DB_HOST").expect("DB_HOST env var is not set");
  let db_username: String = env::var("DB_USERNAME").expect("DB_USERNAME env var is not set");
  let db_name: String = env::var("DB_NAME").expect("DB_NAME env var is not set");
  let db_password: String = env::var("DB_PASSWORD").expect("DB_PASSWORD env var is not set");

  // host=myhost port=myport user=john dbname=gorm password=mypassword sslmode=disable
  let database_url: String = format!("host={DB_HOST} port=5432 user={DB_USERNAME} dbname={DB_NAME} password={DB_PASSWORD} sslmode=disable",
    DB_HOST=db_host,
    DB_USERNAME=db_username,
    DB_NAME=db_name,
    DB_PASSWORD=db_password,
  );

  PgConnection::establish(&database_url).expect(&format!("DB connection failed URL: {}", database_url))
}
