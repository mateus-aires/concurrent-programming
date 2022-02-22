package twoPhaseSleep;

import java.util.Scanner;

public class Main {

    private static Barrier barrier;
    private static int nThreads;
    private static ThreadsArrayUtil threadsArrayUtil;

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("n >> ");
        nThreads = scanner.nextInt();
        scanner.close();

        barrier = new Barrier(nThreads);
        threadsArrayUtil = new ThreadsArrayUtil(nThreads, barrier);

        threadsArrayUtil.createAndRunThreads();
    }
}
