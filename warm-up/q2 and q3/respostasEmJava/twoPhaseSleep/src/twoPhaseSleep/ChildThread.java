package twoPhaseSleep;

public class ChildThread extends Thread {

    private static Barrier barrier;
    final int FIRST_MIN_SLEEP = 1;
    final int FIRST_MAX_SLEEP = 5;
    final int FIRST_RANGE = FIRST_MAX_SLEEP - FIRST_MIN_SLEEP + 1;
    final int SECOND_RANGE = 10;
    private int position;

    public ChildThread(Barrier b, int position) {
        barrier = b;
        this.position = position;
    }

    @Override
    public void run() {
        int firstSleepTime = (int) (Math.random() * FIRST_RANGE) + FIRST_MIN_SLEEP;
        System.out.printf("[%d] Sleeping for %d seconds\n", this.position, firstSleepTime);

        this.takeANap(firstSleepTime * 1000, true);
        ThreadsArrayUtil.setSleepTime(this.position, this.drawNumber());

        barrier.run();

        System.out.print(ThreadsArrayUtil.getSecondSleepTimesString(this.position));

        int secondSleepTime = ThreadsArrayUtil.getSleepTime(this.position);
        this.takeANap(secondSleepTime * 1000, false);
    }

    private int drawNumber() {
        int secondSleepTime = (int) (Math.random() * SECOND_RANGE);
        return secondSleepTime;
    }

    private void takeANap(int sleepTime, boolean isFirst) {
        try {
            sleep(sleepTime);
            if (isFirst) {
                System.out.printf("[%d] Ending first sleep\n", this.position);
            } else {
                System.out.printf("[%d] Ending second sleep\n", this.position);
            }
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

}
