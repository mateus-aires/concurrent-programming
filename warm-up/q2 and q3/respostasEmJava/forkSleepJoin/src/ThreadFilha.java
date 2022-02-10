public class ThreadFilha implements Runnable{
    final int MIN_SLEEP = 1;
    final int MAX_SLEEP = 5;
    final int RANGE = MAX_SLEEP - MIN_SLEEP + 1;
    final int SEC = 1000;
    String name;

    ThreadFilha(String name){
        this.name = name;
    }

    @Override
    public void run() {
        int sleepTime = (int)(Math.random() * RANGE) + MIN_SLEEP;
        System.out.printf("[%s] Sleeping for %d seconds.\n", this.name, sleepTime);

        try {
            Thread.sleep(sleepTime * SEC);
            System.out.println("[%s] Terminating execution".formatted(this.name));
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
