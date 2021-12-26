import java.util.Scanner;
//import java.util.HashMap;
public class day22A {
    static Scanner in = new Scanner(System.in);
    public static void main(String[] args) {
        String sw = in.next();
        boolean[][][] arr = new boolean[101][101][101];
        //HashMap<String, Boolean> map = new HashMap<String, Boolean>();
        int count = 0;
        while (!sw.equals("exit")){
            String []temp = in.next().replaceAll(".=", "").split(",");
            int[] x = new int[2], y = new int[2], z = new int[2];
            String[] tmp2 =temp[0].split("\\.\\."); 
            x[0] = Integer.parseInt(tmp2[0]);
            x[1] = Integer.parseInt(tmp2[1]);
            if (x[0] < -50 || x[0] > 50 || x[1] < -50 || x[1] > 50) {
                sw = in.next();
                continue;
            } 
            //if x[0] > x[1] swap
            if (x[0] > x[1]) {
                int temp2 = x[0];
                x[0] = x[1];
                x[1] = temp2;
            }

            tmp2 =temp[1].split("\\.\\.");
            y[0] = Integer.parseInt(tmp2[0]);
            y[1] = Integer.parseInt(tmp2[1]);
            if (y[0] < -50 || y[0] > 50 || y[1] < -50 || y[1] > 50) {
                sw = in.next();
                continue;
            }
            //if y[0] > y[1] swap
            if (y[0] > y[1]) {
                int temp2 = y[0];
                y[0] = y[1];
                y[1] = temp2;
            }

            tmp2 =temp[2].split("\\.\\.");
            z[0] = Integer.parseInt(tmp2[0]);
            z[1] = Integer.parseInt(tmp2[1]);
            if (z[0] < -50 || z[0] > 50 || z[1] < -50 || z[1] > 50) {
                sw = in.next();
                continue;
            }
            //if z[0] > z[1] swap
            if (z[0] > z[1]) {
                int temp2 = z[0];
                z[0] = z[1];
                z[1] = temp2;
            }

            //brute force x,y,z
            for (int i = x[0]+50; i <= x[1]+50; i++) {
                for (int j = y[0]+50; j <= y[1]+50; j++) {
                    for (int k = z[0]+50; k <= z[1]+50; k++) {
                        if(sw.equals("on") && !arr[i][j][k]){
                            arr[i][j][k] = true;
                            count++;
                        }
                        if(sw.equals("off") && arr[i][j][k]){
                            arr[i][j][k] = false;
                            count--;
                        }
                    }
                }
            }
            sw = in.next();
        }
        System.out.println(count);
    }
}