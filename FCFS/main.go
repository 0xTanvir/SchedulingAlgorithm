package main

import "fmt"

func main() {

	var n,i,j int
	var avwt=0.0
	var avtat=0.0
	var bt[20] int
	var wt[20] int
	var tat[20] int
	fmt.Printf("Enter total number of processes(maximum 20): ")
	fmt.Scanf("%d",&n)

	fmt.Printf("\nEnter Process Burst Time\n")
	for i=0;i<n;i++{
		fmt.Printf("Process %d : ",i+1)
		fmt.Scanf("%d",&bt[i])
	}

	wt[0]=0    //waiting time for first process is 0

	//calculating waiting time
	for i=1;i<n;i++{
		wt[i]=0
		for j=0;j<i;j++{
			wt[i]+=bt[j]
		}
	}

	fmt.Printf("\nProcess\t\tBurst Time\tWaiting Time\tTurnaround Time")

	//calculating turnaround time
	for i=0;i<n;i++{
		tat[i]=bt[i]+wt[i]
		avwt+=float64(wt[i])
		avtat+=float64(tat[i])
		fmt.Printf("\nP[%d]\t\t%d\t\t%d\t\t%d",i+1,bt[i],wt[i],tat[i])
	}

	avwt/=float64(i)
	avtat/=float64(i)
	fmt.Printf("\n\nAverage Waiting Time : %f",avwt)
	fmt.Printf("\nAverage Turnaround Time : %f",avtat)
}