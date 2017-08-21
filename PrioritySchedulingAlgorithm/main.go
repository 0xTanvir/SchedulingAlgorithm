package main

import (
	"fmt"
)

func main() {
	var n,pos,temp int
	var avg_wt,avg_tat float64
	var total = 0
	var burstTime [20] int
	var priority [20] int
	var p [20] int
	var wt [20] int
	var tat [20] int


	fmt.Printf("Enter total number of Process: ")
	fmt.Scanf("%d",&n)
	fmt.Println("Enter Brust time and Priority")

	for i:=0;i<n;i++ {

		fmt.Printf("\nFor Process %d \n",i+1)

		fmt.Printf("Burst Time: ")
		fmt.Scanf("%d",&burstTime[i])
		fmt.Printf("Priority: ")
		fmt.Scanf("%d",&priority[i])
		p[i]=i+1 //contains process number
	}

	//sorting burst time, priority and process number in ascending order using selection sort
	for i:=0;i<n;i++ {
		pos =i
		for j:=i+1;j<n;j++ {
			if priority[j]<priority[pos]{
				pos=j
			}
		}

		temp = priority[i]
		priority[i] = priority[pos]
		priority[pos] = temp

		temp=burstTime[i]
		burstTime[i]=burstTime[pos]
		burstTime[pos]=temp

		temp=p[i]
		p[i]=p[pos]
		p[pos]=temp
	}
	wt[0]=0 //waiting time for first process is zero

	//calculate waiting time
	for i:=1;i<n;i++{
		wt[i]=0
		for j:=0;j<i;j++{
			wt[i]+=burstTime[j]
		}
		total+=wt[i]
	}
	avg_wt = float64(total)/float64(n)  //average waiting time
	total=0

	fmt.Println("\nProcess\t    Burst Time    \tWaiting Time\tTurnaround Time")
	for i:=0;i<n;i++ {
		tat[i]=burstTime[i]+wt[i]     //calculate turnaround time
		total+=tat[i]
		fmt.Printf("\nP[%d]\t\t  %d\t\t    %d\t\t\t%d",p[i],burstTime[i],wt[i],tat[i])
	}

	avg_tat = float64(total)/float64(n)     //average turnaround time
	fmt.Printf("\n\nAverage Waiting Time=%f",avg_wt)
	fmt.Printf("\nAverage Turnaround Time=%f\n",avg_tat)

}
