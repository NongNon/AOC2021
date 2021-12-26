import java.util.Scanner;
/**
 * day21A
 */
public class day21A {
    static Scanner sc = new Scanner(System.in);
    
    public static void main(String[] args) {
        int p1,p2;
        int scoreP1 = 0, scoreP2 = 0;
        String[] temp=sc.nextLine().split(" ");
        p1=Integer.parseInt(temp[temp.length-1]);
        temp=sc.nextLine().split(" ");
        p2=Integer.parseInt(temp[temp.length-1]);
        int i = 6;
        int t=0;
        while (scoreP1<1000 && scoreP2 <1000){
            t++;
            if(i%2==0){
                p1=(p1+i)%10;
                if(p1==0) p1=10;
                scoreP1+=p1;
            }else{
                p2=(p2+i)%10;
                if(p2==0) p2=10;
                scoreP2+=p2;
            }
            i=(i-1==0)?10:i-1;
        }
        int foo=(scoreP1>=1000)?scoreP2:scoreP1;
        System.out.println(3*t*foo);
    }
}