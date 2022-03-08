package twoPhaseSleep;

import java.util.Scanner;

public class Main {

    private static Barrier barrier;
    private static int nThreads;
    private static ChildThread[] threadsArray;
    private static int[] secondSleepTimes;

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("n >> ");
        nThreads = scanner.nextInt();
        scanner.close();

        init();

        for (int i = 0; i < nThreads; i++) {
            ChildThread childThread = new ChildThread(barrier, i);
            threadsArray[i] = childThread;
            childThread.start();
        }
    }

    public static void init() {
        threadsArray = new ChildThread[nThreads];
        secondSleepTimes = new int[nThreads];
        barrier = new Barrier(nThreads);
    }

    public static void setSleepTime(int position, int sleepTime) {
        int setByPosition = position + 1;
        if (position == (nThreads) - 1) {
            setByPosition = 0;
        }
        secondSleepTimes[setByPosition] = sleepTime;
    }

    public static int getSleepTime(int position) {
        return secondSleepTimes[position];
    }

    public static String getSecondSleepTimesString(int position) {
        return String.format("[%d]", position) + String.format(" will sleep for %d seconds\n", secondSleepTimes[position]);
    }
}
