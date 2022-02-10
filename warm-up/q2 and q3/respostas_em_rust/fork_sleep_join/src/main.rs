use std::{thread, io, time};
use std::thread::{JoinHandle, sleep};
use rand::Rng;

fn main() {
    let n = get_input();
    let mut threads_filhas: Vec<JoinHandle<()>> = vec![];

    for n_thr in 1..=n {
        let handle = thread::spawn(move || {
           random_sleeping(n_thr);
        });
        threads_filhas.push(handle);
    }

    for (i, thread) in threads_filhas.into_iter().enumerate() {
        thread.join().expect(&format!("Erro na Thread {}", i));
    }
    println!("{}", n);
}

fn get_input() -> i32 {
    let mut buffer = String::new();
    let stdin = io::stdin();
    stdin.read_line(&mut buffer).unwrap();
    buffer.trim().parse().unwrap()
}

fn random_sleeping(n_thr: i32) {
    let sleep_time = {
        let random = rand::thread_rng().gen_range(1..=5);
        time::Duration::from_secs(random)
    };
    println!("[Thread {}] Sleeping for {} seconds", n_thr, sleep_time.as_secs());
    sleep(sleep_time);
}