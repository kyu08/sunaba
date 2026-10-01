fn main() {
    match fetch() {
        FetchState::Loading(_) => {
            println!("loading");
        }
        FetchState::Loaded(loaded_state) => {
            println!("uid: {}", loaded_state.user_id);
        }
        FetchState::Failed(failed_state) => {
            println!("err: {}", failed_state.error_code);
        }
    }
}

fn fetch() -> FetchState {
    // FetchState::Loading(LoadingState {})
    // FetchState::Failed(FailedState { error_code: 1 })
    FetchState::Loaded(LoadedState {
        user_id: "3e455ddc-9301-4b61-9e7a-a256471ce278".to_string(),
    })
}

#[derive(Debug)]
enum FetchState {
    Loading(LoadingState),
    Loaded(LoadedState),
    Failed(FailedState),
}

#[derive(Debug)]
struct LoadingState {}
#[derive(Debug)]
struct LoadedState {
    user_id: String,
}
#[derive(Debug)]
struct FailedState {
    error_code: usize,
}
