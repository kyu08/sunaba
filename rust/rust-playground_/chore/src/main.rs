use std::collections::HashMap;
use std::sync::{Arc, Mutex};
use tokio::task;
use tokio::time::{sleep, Duration};
use update_informer::{registry, Check};

struct Model {
    data: Option<String>,
}

const VERSION_KEY: &str = "data";

#[tokio::main]
async fn main() {
    let model: &mut Model = &mut Model { data: None };
    let data = Arc::new(Mutex::new(HashMap::new()));
    let data_clone = data.clone();

    tokio::spawn(async move {
        let pkg_name = "kyu08/fzf-make";
        let current_version = "0.44.0"; // TODO: get from env vars
        let informer = update_informer::new(registry::GitHub, pkg_name, current_version)
            .interval(Duration::ZERO); // TODO: fix duration
                                       // panic!("{:?} is available", "0000 ");
        let version_result =
            task::spawn_blocking(|| informer.check_version().map_err(|e| e.to_string()))
                .await
                .unwrap();
        if let Ok(Some(new_version)) = version_result {
            let mut data = data_clone.lock().unwrap();
            data.insert(VERSION_KEY.to_string(), new_version.to_string());
            // panic!("{:?} is available", new_version.to_string());
        }
    });

    loop {
        if model.data.is_none() {
            let data = data.lock();
            println!("got lock!");
            if let Ok(d) = data {
                if let Some(fetched_data) = d.get(VERSION_KEY) {
                    model.data = Some(fetched_data.clone());
                }
            }
        }

        println!("{:?}", model.data);
        sleep(Duration::from_secs(1)).await; // ポーリング間の間隔
    }
}
