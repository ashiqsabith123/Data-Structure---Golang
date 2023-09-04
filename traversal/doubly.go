package main


type GNode struct{
	next *GNode
	prev *GNode
	data int
}


func () addNode(target int){

	node=newnode


	temp := head

	for temp!=nil{

		if temp.data==target{
			prev = temp.next
			temp.next=node
			node.next=prev

			node.prev=temp
			node.next.prev=node
		}
	}

}


func main(){



}