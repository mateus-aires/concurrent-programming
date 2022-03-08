package twoPhaseSleep;

import java.util.concurrent.Semaphore;

public class Barrier {

    private static int nThreads;

    private static int count;
    private static Semaphore mutex;
    private static Semaphore firstGate;
    private static Semaphore secondGate;

    public Barrier(int n) {
        nThreads = n;
        count = 0;
        mutex = new Semaphore(1);
        firstGate = new Semaphore(0);
        secondGate = new Semaphore(1);

    }

    public void build() {
        try {
            acquire();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    private void acquire() throws InterruptedException {
        // first phase
        mutex.acquire();
        count += 1;
        if (count == nThreads) {
            secondGate.acquire();
            firstGate.release();
        }
        mutex.release();

        firstGate.acquire();
        firstGate.release();

        // second phase
        mutex.acquire();
        count -= 1;
        if (count == 0) {
            firstGate.acquire();
            secondGate.release();
        }
        mutex.release();

        secondGate.acquire();
        secondGate.release();
    }

}
