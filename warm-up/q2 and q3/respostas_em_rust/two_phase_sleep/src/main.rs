/*  two-phase sleep
    Crie um programa que recebe um número inteiro n como argumento e cria n threads.
    Cada uma dessas threads deve dormir por um tempo aleatório de no máximo 5 segundos.
    Depois que acordar, cada thread deve sortear um outro número aleatório s
    (entre 0 e 10). Somente depois de todas as n threads terminarem suas escolhas
    (ou seja, ao fim da primeira fase), começamos a segunda fase.
    Nesta segunda fase, a n-ésima thread criada deve dormir pelo tempo s escolhido
    pela thread n - 1 (faça a contagem de maneira modula, ou seja, a primeira thread
    dorme conforme o número sorteado pela última).
*/
use std::{thread, io, time};
use rand;
use std::sync::{Arc, Mutex};
use rand::Rng;
use std_semaphore::Semaphore;


fn main() {
    let n = get_input();

    let second_sleep_times: Arc<Mutex<Vec<i32>>> = {
        let mut temp = vec![];
        for _ in 0..n {
            temp.push(-1);
        }
        Arc::new(Mutex::new(temp))
    };
    let first_sleep_mutex: Arc<Mutex<i32>> = Arc::new(Mutex::new(0));
    let first_sleep_barrier: Arc<Semaphore> = Arc::new(Semaphore::new(0));
    let second_sleep_mutex: Arc<Mutex<i32>> = Arc::new(Mutex::new(0));
    let second_sleep_barrier: Arc<Semaphore> = Arc::new(Semaphore::new(0));

    for i in 0..n {
        let st = Arc::clone(&second_sleep_times);
        let fsm = Arc::clone(&first_sleep_mutex);
        let fsb = Arc::clone(&first_sleep_barrier);
        let ssm = Arc::clone(&second_sleep_mutex);
        let ssb = Arc::clone(&second_sleep_barrier);

        let _handle = thread::spawn(move || {
            two_phase_sleep(
                &i,
                &n,
                st,
                fsm,
                fsb,
                ssm,
                ssb
            );
        });
    }

    second_sleep_barrier.acquire();
    println!("[Thread mãe]: n = {}", n);
}

fn get_input() -> i32{
    let mut buffer = String::new();
    let stdin = io::stdin();
    stdin.read_line(&mut buffer).unwrap();
    let x = buffer.trim().parse::<i32>().unwrap();
    x
}

fn two_phase_sleep(index: &i32, n_threads: &i32, sleep_times: Arc<Mutex<Vec<i32>>>,
                   first_sleep_mutex: Arc<Mutex<i32>>, first_sleep_barrier: Arc<Semaphore>,
                   second_sleep_mutex: Arc<Mutex<i32>>, second_sleep_barrier: Arc<Semaphore>){
    let sleep_time = {
        let seconds = rand::thread_rng().gen_range(1..=5);
        time::Duration::from_secs(seconds)
    };
    println!("[Thread #{}] First sleep: {} seconds", index, sleep_time.as_secs());
    thread::sleep(sleep_time);
    println!("[Thread #{}] Finished first sleep", index);

    {
        let mut times = sleep_times.lock().unwrap();
        let random_sleep = rand::thread_rng().gen_range(0..=10);
        println!("[Thread #{}] Random sleep drawn: {} seconds", index, random_sleep);
        times[*index as usize] = random_sleep;
    }

    {
        let mut first_sleep_counter = first_sleep_mutex.lock().unwrap();
        *first_sleep_counter += 1;

        if *first_sleep_counter == *n_threads {
            println!("Sleep times: {:?}", sleep_times);
            first_sleep_barrier.release();
        }
    }
    first_sleep_barrier.acquire();
    first_sleep_barrier.release();

    let second_sleep_time = {
        let i: usize = ((index + 1) % n_threads) as usize;
        let seconds = sleep_times.lock().unwrap()[i];
        time::Duration::from_secs(seconds.clone() as u64)
    };
    println!("[Thread #{}] Second sleep: {} seconds", index, second_sleep_time.as_secs());
    thread::sleep(second_sleep_time);
    println!("[Thread #{}] Finished second sleep", index);

    {
        let mut second_sleep_counter = second_sleep_mutex.lock().unwrap();
        *second_sleep_counter += 1;

        if *second_sleep_counter == *n_threads {
            println!("[Thread #{}] Releasing mother thread", index);
            second_sleep_barrier.release();
        }
    }
}
