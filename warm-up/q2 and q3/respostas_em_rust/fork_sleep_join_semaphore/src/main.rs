use std::{thread, io, time};
use std::thread::sleep;
use rand::Rng;
use std::sync::{Arc, Mutex};
use std_semaphore::Semaphore;


fn main() {
    let n = get_input();
    let threads_finalizadas = Arc::new(Mutex::new(0));
    let barrier = Arc::new(Semaphore::new(0));

    for n_thr in 1..=n {
        let threads_finalizadas = Arc::clone(&threads_finalizadas);
        let barrier_clone = Arc::clone(&barrier);
        let _handle = thread::spawn(move || {
            random_sleeping(n_thr, &n,threads_finalizadas, barrier_clone);
        });
    }

    barrier.acquire();

    println!("{}", n);
}

fn get_input() -> i32 {
    let mut buffer = String::new();
    let stdin = io::stdin();
    stdin.read_line(&mut buffer).unwrap();
    buffer.trim().parse().unwrap()
}

fn random_sleeping(n_thr: i32, total_threads: &i32, contador: Arc<Mutex<i32>>, barrier: Arc<Semaphore>) {
    let sleep_time = {
        let random = rand::thread_rng().gen_range(1..=5);
        time::Duration::from_secs(random)
    };

    println!("[Thread {}] Sleeping for {} seconds", n_thr, sleep_time.as_secs());
    sleep(sleep_time);
    println!("[Thread {}] Finished", n_thr);

    {
        let mut threads_finalizadas = contador.lock().unwrap();
        *threads_finalizadas += 1;

        if *threads_finalizadas == *total_threads {
            barrier.release();
        }
    };


}