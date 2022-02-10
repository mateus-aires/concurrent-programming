import java.util.Scanner;
import java.util.concurrent.Semaphore;

public class ThreadMae{
    static int totalThreads;
    static int finishedThreads = 0;
    static Semaphore mutex = new Semaphore(1);
    static Semaphore barrier = new Semaphore(0);
    public static void main(String[] args) throws InterruptedException {
        Scanner scanner = new Scanner(System.in);
        System.out.print("n >> ");
        totalThreads = scanner.nextInt();

        for (int i = 1; i <= totalThreads; i++) {
            ThreadFilha thr = new ThreadFilha("Thread #%d".formatted(i));
            thr.start();
        }

        barrier.acquire();
        System.out.println(totalThreads);
    }

    static class ThreadFilha extends Thread {
        final int MIN_SLEEP = 1;
        final int MAX_SLEEP = 5;
        final int RANGE = MAX_SLEEP - MIN_SLEEP + 1;

        ThreadFilha(String name) {
            super(name);
        }

        @Override
        public void run() {
            int sleepTime = (int)(Math.random() * RANGE) + MIN_SLEEP;
            System.out.printf("[%s] Sleeping for %d seconds%n", getName(), sleepTime);

            try {
                sleep(sleepTime * 1000L);

                mutex.acquire();
                finishedThreads++;
                System.out.printf("[%s] Terminating execution%n", getName());
                if (finishedThreads == totalThreads) {
                    barrier.release();
                }
                mutex.release();

            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}

