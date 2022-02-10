import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.util.concurrent.Semaphore;

public class ThreadMae {
    /* two-phase sleep Crie um programa que recebe um número inteiro n como argumento e cria n threads.
    Cada uma dessas threads deve dormir por um tempo aleatório de no máximo 5 segundos.
    Depois que acordar, cada thread deve sortear um outro número aleatório s (entre 0 e 10).
    Somente depois de todas as n threads terminarem suas escolhas (ou seja, ao fim da primeira fase), começamos a segunda fase.
    Nesta segunda fase, a n-ésima thread criada deve dormir pelo tempo s escolhido pela thread n - 1
    (faça a contagem de maneira modula, ou seja, a primeira thread dorme conforme o número sorteado pela última).
    Obs.3 Para a questão 3 use semáforos para coordenar o trabalho entre as threads.
    */
    static List<Integer> sleepTimes = new ArrayList<>();
    static List<Semaphore> semaphores = new ArrayList<>();
    static int nThreads;

    public static void main(String[] args) throws InterruptedException {
        Scanner scanner = new Scanner(System.in);
        System.out.print("n >> ");
        nThreads = scanner.nextInt();
        scanner.close();

        // Inicializando a lista com os tempos da second sleep
        for (int i = 0; i < nThreads ; i++) {
            sleepTimes.add(-1);
        }

        // Criando threads filhas
        for (int i = 0; i < nThreads ; i++) {
            Semaphore semaphore = new Semaphore(0);
            ThreadFilha thr = new ThreadFilha(i, semaphore);
            thr.start();
            semaphores.add(semaphore);
        }

        // Esperar as threads filhas terminarem
        for (Semaphore sem : semaphores) {
            sem.acquire();
        }
        System.out.println(nThreads);
        System.out.println(sleepTimes);
    }

    static class ThreadFilha extends Thread {
        final int MIN_SLEEP = 1;
        final int MAX_SLEEP = 5;
        final int RANGE = MAX_SLEEP - MIN_SLEEP + 1;
        Semaphore semaphore;
        int index;

        ThreadFilha(int index, Semaphore semaphore) {
            super("Thread #%d".formatted(index));
            this.index = index;
            this.semaphore = semaphore;
        }

        @Override
        public void run() {
            int sleepTime = (int) (Math.random() * RANGE) + MIN_SLEEP;
            System.out.printf("[%s] Sleeping for %d seconds\n", getName(), sleepTime);

            // First sleep
            try {
                sleep(sleepTime * 1000L);
                System.out.printf("[%s] Ending first sleep\n", getName());
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            int s = (int) (Math.random() * RANGE) + MIN_SLEEP;
            sleepTimes.set(this.index, s);
            System.out.printf("[%s] Added new sleep time: %d\n", getName(), s);

            // Wait until all threads added their random sleep time
            while (sleepTimes.contains(-1)) {
                ThreadFilha.yield();
            }

            // Pick new sleep time
            if (index == 0) {
                sleepTime = sleepTimes.get(nThreads - 1);
            } else {
                sleepTime = sleepTimes.get(index - 1);
            }
            System.out.printf("[%s] Sleeping AGAIN for %d seconds\n", getName(), sleepTime);

            // Second sleep
            try {
                sleep(sleepTime * 1000L);
                System.out.printf("[%s] Ending second sleep\n", getName());
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            // Release main thread
            semaphore.release();
        }
    }
}
