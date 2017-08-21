package main

import "fmt"

func main(){
	var i,j,n,pos,temp int
	var total=0
	var  avg_wt,avg_tat float64
	var bt[20]  int
	var p[20]  int
	var wt[20]  int
	var tat[20]  int

	fmt.Printf("Enter number of process:")
	fmt.Scanf("%d",&n)
	fmt.Printf("\nEnter Burst Time:\n")
	for i=0;i<n;i++ {
		fmt.Printf("p%d:",i+1)
		fmt.Scanf("%d",&bt[i])
		p[i]=i+1       //contains process number
	}

	//sorting burst time in ascending order using selection sort
	for i=0;i<n;i++ {
		pos=i
		for j=i+1;j<n;j++{
			if bt[j]<bt[pos] {
				pos=j
			}

		}
		temp=bt[i]
		bt[i]=bt[pos]
		bt[pos]=temp


		temp=p[i]
		p[i]=p[pos]
		p[pos]=temp
	}
	wt[0]=0           //waiting time for first process will be zero

	//calculate waiting time
	for i=1;i<n;i++ {
		wt[i]=0
		for j=0;j<i;j++ {
			wt[i]+=bt[j]
		}
		total+=wt[i]
	}

	avg_wt=float64(total)/float64(n)      //average waiting time
	total=0
	fmt.Printf("\nProcess\t    Burst Time    \tWaiting Time\tTurnaround Time")
	for i=0;i<n ;i++{
		tat[i]=bt[i]+wt[i]    //calculate turnaround time
		total+=tat[i]
		fmt.Printf("\np%d\t\t  %d\t\t    %d\t\t\t%d",p[i],bt[i],wt[i],tat[i])
	}
	avg_tat=float64(total)/float64(n)     //average turnaround time
	fmt.Printf("\n\nAverage Waiting Time=%f",avg_wt)
	fmt.Printf("\nAverage Turnaround Time=%f\n",avg_tat)

}
