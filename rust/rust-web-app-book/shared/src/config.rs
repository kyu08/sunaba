use anyhow::Result;

pub struct AppConfig {
    pub database: DatabaseConfig,
}

pub struct RedisConfig {
    pub host: String,
    pub port: u16,
}

impl AppConfig {
    pub fn new() -> Result<Self> {
        let database = DatabaseConfig {
            host: "localhost".into(),
            port: 5432,
            username: "app".into(),
            password: "passwd".into(),
            database: "app".into(),
        };
        Ok(Self { database })
    }
}

pub struct DatabaseConfig {
    pub host: String,
    pub port: u16,
    pub username: String,
    pub password: String,
    pub database: String,
}
