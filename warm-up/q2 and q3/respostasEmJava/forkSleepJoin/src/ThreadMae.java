import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class ThreadMae {
    /*
    1. Fork-sleep-join - Crie um programa que recebe um número inteiro n como argumento e
    cria n threads. Cada uma dessas threads deve dormir por um tempo aleatório de no máximo 5 segundos.
    A main-thread deve esperar todas as threads filhas terminarem de executar para em seguida escrever
    na saída padrão o valor de n.
    Faça a thread-mãe esperar as filhas de duas maneiras:
    1) usando o equivalente à função join em c, como mostrado em aula;
    2) usando semáforos.
    */
    public static void main(String[] args) throws InterruptedException {
        Scanner scanner = new Scanner(System.in);
        System.out.print("nThreads >> ");
        int nThreads = scanner.nextInt();
        List<Thread> thrFilhas = new ArrayList<>();

        for (int i = 1; i <= nThreads; i++) {
            ThreadFilha thr_runnable = new ThreadFilha("Thread #0%d".formatted(i));
            Thread thr = new Thread(thr_runnable);
            thr.start();
            thrFilhas.add(thr);
        }
        for (Thread thr : thrFilhas) {
            thr.join();
        }


        System.out.println("Thread mãe executou " + nThreads + " threads");
    }
}
