package main

import "fmt"

func main() {
	var count,n,time,remain,flag,wait_time,turnaround_time,time_quantum int
	flag = 0
	wait_time=0
	turnaround_time=0
	var at[10] int
	var bt[10] int
	var rt[10] int


	fmt.Printf("Enter Total Process:\t ")
	fmt.Scanf("%d",&n)
	remain=n
	for count:=0;count<n;count++{
		fmt.Printf("\nFor Process Number %d \n",count+1)
		fmt.Printf("Enter Arrival Time : ")
		fmt.Scanf("%d",&at[count])
		fmt.Printf("Enter Burst Time for Process Process Number : ")
		fmt.Scanf("%d",&bt[count])
		rt[count]=bt[count]
	}
	fmt.Printf("\nEnter Time Quantum:\t")
	fmt.Scanf("%d",&time_quantum)
	fmt.Printf("\n\nProcess\t|Turnaround Time|Waiting Time\n\n")

	time=0
	count=0
	for remain!=0{
		/**/
		if rt[count]<=time_quantum && rt[count]>0 {
			time+=rt[count]
			rt[count]=0
			flag=1

		} else if rt[count]>0{
			rt[count]-=time_quantum
			time+=time_quantum
		}
		if rt[count]==0 && flag==1 {
			remain--
			fmt.Printf("P[%d]\t|\t%d\t|\t%d\n",count+1,time-at[count],time-at[count]-bt[count])
			wait_time+=time-at[count]-bt[count]
			turnaround_time+=time-at[count]
			flag=0
		}
		if count==n-1{
			count=0
		} else if at[count+1]<=time{
			count++
		} else {count=0}
	}
	fmt.Printf("\nAverage Waiting Time = %f \n",float64(wait_time)/float64(n))
	fmt.Printf("Avg Turnaround Time = %f ",float64(turnaround_time)/float64(n))

}