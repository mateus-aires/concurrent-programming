package twoPhaseSleep;

public class ThreadsArrayUtil {

    private static ChildThread[] threadsArray;
    private static Barrier barrier;
    private static int nThreads;
    private static int[] secondSleepTimes;

    public ThreadsArrayUtil(int n, Barrier b) {
        this.nThreads = n;
        threadsArray = new ChildThread[nThreads];
        secondSleepTimes = new int[nThreads];
        barrier = b;
    }

    public void createAndRunThreads() {
        for (int i = 0; i < nThreads; i++) {
            ChildThread childThread = new ChildThread(barrier, i);
            threadsArray[i] = childThread;
            childThread.start();
        }
    }

    public static void setSleepTime(int position, int sleepTime) {
        int setByPosition = position;
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
