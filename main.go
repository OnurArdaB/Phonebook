package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)
func HasRight(node * TreeNode)bool{
	if(node.right!=nil){
		return true
	}
	return false
}
func TraverseNodes(sb *strings.Builder,padding string,pointer string,node*TreeNode,hasRight bool){
	if (node != nil) {
		sb.WriteString("\n");
		sb.WriteString(padding);
		sb.WriteString(pointer);
		sb.WriteString(node.firstName + " " +node.secondName);

		paddingBuilder :=  strings.Builder{}
		paddingBuilder.WriteString(padding)
		if (hasRight) {
			paddingBuilder.WriteString("│  ");
		} else {
			paddingBuilder.WriteString("   ");
		}

		paddingForBoth := paddingBuilder.String();
		pointerRight := "└──";
		var pointerLeft string
		if(node.right!=nil){
			pointerLeft = "├──"
		}else{
			pointerLeft =  "└──"
		}

		TraverseNodes(sb, paddingForBoth, pointerLeft, node.left, HasRight(node))
		TraverseNodes(sb, paddingForBoth, pointerRight, node.right, false)
	}

}
func TraversePreOrder(ROOT * TreeNode)string{
	if(ROOT == nil){
		return ""
	}
	sb:=strings.Builder{}
	sb.WriteString(ROOT.firstName + " " +ROOT.secondName)
	var pointerRight = "└──"
	var pointerLeft string
	if(ROOT.right!=nil){
		pointerLeft = "├──"
	}else{
		pointerLeft = "└──"
	}
	TraverseNodes(&sb, "", pointerLeft, ROOT.left, HasRight(ROOT));
	TraverseNodes(&sb, "", pointerRight, ROOT.right, false)
	return sb.String()
}

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
} //Works fine
func check(e error) {
	if e != nil {
		panic(e)
	}
} //Works fine
type TreeNode struct{
	firstName,secondName,phoneNumber,city string
	height int
	left * TreeNode
	right *TreeNode
} //Works fine

type BST struct {
	root * TreeNode
	size int
	sizeOfLeft int
	sizeOfRight int
} //Works fine
func (TREE * BST)PrintDiagram(){
	file,err := os.Create("outputfileBST.txt")
	check(err)
	file.WriteString(TraversePreOrder(TREE.root))
	defer file.Close()
}
func (TREE *BST) insert(first,second,phone,city string){//The insertion operation is first come first served for BST.Hold heights in insertion.
	var node = TreeNode{
		firstName:   first,
		secondName:  second,
		phoneNumber: phone,
		city:        city,
		height:       1,
		left:        nil,
		right:       nil,
	}
	var firstIsLeft = false
	if(TREE.root==nil){
		node.height=1
		TREE.root = &node
	}else{
		var tempROOT = TREE.root
		var key=first+second
		if(tempROOT.firstName + tempROOT.secondName>key){
			firstIsLeft=true
		}
		for(tempROOT!=nil){
			node.height++
			if(tempROOT.firstName + tempROOT.secondName>key){
				if(tempROOT.left==nil){
					tempROOT.left=&node
					break
				}
				tempROOT=tempROOT.left
			}else if(tempROOT.firstName + tempROOT.secondName<key){
				if(tempROOT.right==nil){
					tempROOT.right=&node
					break
				}
				tempROOT=tempROOT.right
			}
			if(firstIsLeft){
				TREE.sizeOfLeft = int(Max(int64(TREE.sizeOfLeft),int64(tempROOT.height)))
			}else{
				TREE.sizeOfRight = int(Max(int64(TREE.sizeOfRight),int64(tempROOT.height)))
			}
		}
	}
	TREE.size++
}//Works fine
func (TREE *BST) search(first,second string)(f,s,p,c string){
	var key = first + second
	var tempRoot = TREE.root
	for tempRoot!=nil{
		if(key==tempRoot.firstName+tempRoot.secondName){
			return tempRoot.firstName,tempRoot.secondName,tempRoot.phoneNumber,tempRoot.city
		}else if(key>tempRoot.firstName+tempRoot.secondName){
			tempRoot=tempRoot.right
		}else{
			tempRoot=tempRoot.left
		}
	}
	return
}
func (TREE * BST) searchSubstring(first string)(){
	var tempRoot = TREE.root
	for(tempRoot!=nil){
		if(len(tempRoot.firstName)>len(first)){
			var prefix = tempRoot.firstName[:len(first)]
			if(prefix>first){
				if(tempRoot.left==nil){
					break
				}
				tempRoot = tempRoot.left
			}else if(prefix<first){
				if(tempRoot.right==nil){
					break
				}
				tempRoot = tempRoot.right
			}else{
				break
			}
		}else{
			break
		}
	}
	var found = false
	if(strings.HasPrefix(tempRoot.firstName,first)){
		searchPrefix(tempRoot,&first,&found)
	}else{
		fmt.Println("No contact exists for the entered word")
	}
	//temp root is now the root for the prefix
}
func (TREE *BST) delete(first,second string)(deleted bool){

	var key = first + second
	TREE.root=makeDelete(TREE.root,key)
	return true
}
func (TREE *BST) inorder(){
	file,err := os.Create("phonebookInOrder.txt")
	check(err)
	makeInorder(TREE.root,file)
	defer file.Close()
}
func (TREE *BST) preorder(){
	file,err := os.Create("phonebookPreOrder.txt")
	check(err)
	makePreorder(TREE.root,file)
	defer file.Close()
}

func calculateHeight(node * TreeNode)int{
	if(node==nil){
		return -1
	}
	leftHeight:=calculateHeight(node.left)
	rightHeight:=calculateHeight(node.right)
	return int(Max(int64(leftHeight),int64(rightHeight)))+1
}

func searchPrefix(SUBTREE *TreeNode ,first *string,found *bool){
	if(SUBTREE==nil){
		return
	}
	searchPrefix(SUBTREE.left,first,found)
	if(strings.HasPrefix(SUBTREE.firstName,*first)){
		fmt.Println(strings.ToUpper(SUBTREE.firstName),strings.ToUpper(SUBTREE.secondName),SUBTREE.phoneNumber,SUBTREE.city)
	}
	searchPrefix(SUBTREE.right,first,found)
}

func SearchWord(TREE * BST,word string){
	if(strings.Contains(word," ")){
		var splitted = strings.Split(word," ")
		var firstname,secondname,phone,city = TREE.search(splitted[0],splitted[1])
		if(len(firstname+secondname+phone+city)>0){
			fmt.Println(strings.ToUpper(firstname),strings.ToUpper(secondname),phone,city)
		}else{
			fmt.Println("Contact does not exists...")
		}
	}else{
		TREE.searchSubstring(word)
	}
}

func minValueNode(node * TreeNode)*TreeNode{ //Works fine
	var current = node
	for (current!=nil && current.left!=nil){
		current = current.left
	}
	return current
}

func makeDelete(root * TreeNode,key string)(node *TreeNode){
	if(root==nil){
		return nil
	}
	if(root.firstName + root.secondName<key){
		root.right = makeDelete(root.right,key)
	}else if(root.firstName + root.secondName>key){
		root.left = makeDelete(root.left,key)
	}else{
		if (root.left == nil){
			var temp = root.right
			root=nil
			return temp
		}else if (root.right == nil){
			var temp = root.right
			root=nil
			return temp
		}

		var temp = minValueNode(root.right)

		root.firstName,root.secondName = temp.firstName,temp.secondName

		root.right = makeDelete(root.right, temp.firstName+temp.secondName);
	}

	return root
}

func makeInorder(node*TreeNode,file * os.File){
	if(node==nil){
		return
	}

	makeInorder(node.left,file)
	_,err:=file.WriteString(strings.ToUpper(node.firstName+" "+node.secondName)+" "+node.phoneNumber+" "+node.city)
	check(err)
	makeInorder(node.right,file)

}

func makePreorder(node*TreeNode,file * os.File){
	if(node==nil){
		return
	}
	_,err:=file.WriteString(strings.ToUpper(node.firstName+" "+node.secondName)+" "+node.phoneNumber+" "+node.city)
	check(err)
	makePreorder(node.left,file)
	makePreorder(node.right,file)
}

func rightRotate(node * TreeNode)*TreeNode{
	x:= node.left
	T2:=x.right

	x.right = node
	node.left = T2

	node.height = int(Max(int64(height(node.left)),int64(height(node.right)))) + 1
	x.height = int(Max(int64(height(x.left)),int64(height(x.right)))) + 1


	return x
}

func leftRotate(node * TreeNode)*TreeNode{
	y:=node.right
	T2:=y.left

	y.left = node
	node.right = T2

	node.height = int(Max(int64(height(node.left)),int64(height(node.right)))) + 1
	y.height = int(Max(int64(height(y.left)),int64(height(y.right)))) + 1
	return y
}

func getBalance(node * TreeNode)int{
	if(node == nil){
		return 0
	}
	return height(node.left)-height(node.right)
}

func height(node * TreeNode)int{
	if(node==nil){
		return 0
	}
	return node.height
}

func makeInsert(node *TreeNode,first,second,phone,city string)*TreeNode{
	if(node==nil){
		new:=TreeNode{
			firstName:   first,
			secondName:  second,
			phoneNumber: phone,
			city:        city,
			height:      1,
			left:        nil,
			right:       nil,
		}
		return &new
	}

	if(first+second<node.firstName+node.secondName){
		node.left = makeInsert(node.left,first,second,phone,city)
	}else if(first+second>node.firstName+node.secondName){
		node.right = makeInsert(node.right,first,second,phone,city)
	}else{
		return node
	}

	node.height = 1 + int(Max(int64(height(node.left)),int64(height(node.right))))

	var balance = getBalance(node)

	if(balance>1 && first+second<node.left.firstName+node.left.secondName){
		return rightRotate(node)
	}

	if(-1>balance && first+second>node.right.firstName+node.right.secondName){
		return leftRotate(node)
	}

	if(balance>1 && first+second>node.left.firstName+node.left.secondName){
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if(-1>balance && first+second<node.right.firstName+node.right.secondName){
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func makeDeleteAVL(node *TreeNode,key string)*TreeNode{
	if(node==nil){
		return node
	}

	if(key<node.firstName+node.secondName){
		node.left = makeDeleteAVL(node.left,key)
	}else if(key>node.firstName+node.secondName){
		node.right = makeDeleteAVL(node.right,key)
	}else{
		if(node.left == nil || node.right == nil){
			var temp  *TreeNode
			if(node.left!=nil){
				temp=node.left
			}else {
				temp=node.right
			}

			if(temp==nil){
				temp = node
				node = nil
			}else{
				*node = *temp
			}
		}else{
			temp:=minValueNode(node.right)
			node.firstName,node.secondName = temp.firstName,temp.secondName
			node.right = makeDeleteAVL(node.right,temp.firstName+temp.secondName)
		}
	}
	if(node==nil){
		return node
	}

	node.height = 1 + int(Max(int64(height(node.left)),int64(height(node.right))))

	balance:=getBalance(node)

	if(balance>1&&getBalance(node.left)>=0){
		return rightRotate(node)
	}

	if(balance>1&&getBalance(node.left)<0){
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if(-1>balance&&getBalance(node.right)<=0){
		return leftRotate(node)
	}

	if(-1>balance&&getBalance(node.right)>0){
		node.right = rightRotate(node.left)
		return leftRotate(node)
	}
	return node
}

type AVL struct{
	root * TreeNode
	size int
	sizeOfLeft int
	sizeOfRight int
} //Works
func SearchWordAVL(TREE * AVL,word string){
	if(strings.Contains(word," ")){
		var splitted = strings.Split(word," ")
		var firstname,secondname,phone,city = TREE.search(splitted[0],splitted[1])
		if(len(firstname+secondname+phone+city)>0){
			fmt.Println(strings.ToUpper(firstname),strings.ToUpper(secondname),phone,city)
		}else{
			fmt.Println("Contact does not exists...")
		}

	}else{
		TREE.searchSubstring(word)
	}
}
func (TREE* AVL)PrintDiagram(){
	file,err := os.Create("outputfileAVL.txt")
	check(err)
	file.WriteString(TraversePreOrder(TREE.root))
	defer file.Close()
}
func (TREE *AVL)delete(first,second string){
	TREE.root=makeDeleteAVL(TREE.root,first+second)
}
func (TREE *AVL)insert(first,second,phone,city string){
	TREE.root=makeInsert(TREE.root,first,second,phone,city)
	TREE.size++
} //Works fine
func (TREE *AVL) inorder(){
	file,err := os.Create("phonebookInOrder.txt")
	check(err)
	makeInorder(TREE.root,file)
	defer file.Close()
}
func (TREE *AVL) preorder(){
	file,err := os.Create("phonebookPreOrder.txt")
	check(err)
	makePreorder(TREE.root,file)
	defer file.Close()
} //Works fine
func (TREE *AVL) heights()(x,y int){
	return calculateHeight(TREE.root.left),calculateHeight(TREE.root.right)
}
func (TREE * AVL) searchSubstring(first string)(){
	var tempRoot = TREE.root
	for(tempRoot!=nil){
		if(!strings.HasPrefix(tempRoot.firstName,first)){
			if(len(tempRoot.firstName)>len(first)){
				var prefix = tempRoot.firstName[:len(first)]
				if(prefix>first){
					if(tempRoot.left==nil){
						break
					}
					tempRoot = tempRoot.left
				}else if(prefix<first){
					if(tempRoot.right==nil){
						break
					}
					tempRoot = tempRoot.right
				}
			}else{
				break
			}
		}else{
			break
		}
	}
	var found = false
	if(strings.HasPrefix(tempRoot.firstName,first)){
		searchPrefix(tempRoot,&first,&found)
	}else{
		fmt.Println("No contact exists for the entered word")
	}
}
func (TREE *AVL) search(first,second string)(f,s,p,c string){
	var key = first + second
	var tempRoot = TREE.root
	for tempRoot!=nil{
		if(key==tempRoot.firstName+tempRoot.secondName){
			return tempRoot.firstName,tempRoot.secondName,tempRoot.phoneNumber,tempRoot.city
		}else if(key>tempRoot.firstName+tempRoot.secondName){
			tempRoot=tempRoot.right
		}else{
			tempRoot=tempRoot.left
		}
	}
	return
}

func main(){

	fmt.Print("Please enter the contact file name: ")
	reader := bufio.NewReader(os.Stdin)
	filename, _ := reader.ReadString('\n')
	for(filename[len(filename)-1:]==" "||filename[len(filename)-1:]=="\n"){
		filename=filename[:len(filename)-1]
	}
	fmt.Println(filename)
	fmt.Println("Loading the phonebook into a BST\n")

	var BinarySearchTree = BST{}
	dat, err := ioutil.ReadFile(filename)
	check(err)

	lines:=strings.Split(string(dat),"\n")

	var timeStart = time.Now()

	for _,line:=range lines[:len(lines)]{
		specs := strings.Split(line," ")
		past:=BinarySearchTree.size
		BinarySearchTree.insert(specs[0],specs[1],specs[2],specs[3])
		current:=BinarySearchTree.size
		if(current<=past){
			os.Exit(4)
		}
	}

	var timeStop = time.Now()

	fmt.Println("Phonebook creation in BST took",timeStop.Sub(timeStart),".\n")//turn to milliseconds
	if(BinarySearchTree.sizeOfLeft!=BinarySearchTree.sizeOfRight){
		fmt.Println("The tree is not balanced\n")
	}else{
		fmt.Println("The tree is balanced\n")
	}
	fmt.Println("The heights of the BST are for left:",BinarySearchTree.sizeOfLeft+1,"and right:",BinarySearchTree.sizeOfRight+1)

	//AVL Tree

	fmt.Println("Loading the phonebook into a AVL\n")

	var AVL = AVL{}

	dat, err = ioutil.ReadFile(filename[:len(filename)])
	check(err)

	lines=strings.Split(string(dat),"\n")

	timeStart = time.Now()

	for _,line:=range lines[:len(lines)]{
		specs := strings.Split(line," ")
		past:=AVL.size
		AVL.insert(specs[0],specs[1],specs[2],specs[3])
		current:=AVL.size
		if(current<=past){
			os.Exit(4)
		}
	}

	timeStop = time.Now()

	fmt.Println("Phonebook creation in AVL took",timeStop.Sub(timeStart),".\n")//turn to milliseconds

	var left,right = AVL.heights()

	if(right!=left && right-left>1 || left-right>1){
		fmt.Println("The tree is not balanced\n")
	}else{
		fmt.Println("The tree is balanced\n")
	}

	fmt.Println("The heights of the AVL are for left:",left+1,"and right:",right+1)

	var input int

	var flag bool

	for flag==false {
		fmt.Println("\nChoose which action to perform from 1 to 6:")
		fmt.Println("1 - Search a phonebook contact")
		fmt.Println("2 - Adding a phonebook contact ")
		fmt.Println("3 - Deleting a phonebook contact")
		fmt.Println("4 - Print the phonebook to a file(inorder)\n    Print the phonebook to a file(postorder) ")
		fmt.Println("5 - Draw the Phonebook as a Tree to a file")
		fmt.Println("6 - Press 6 to exit")
		fmt.Scan(&input)
		if (input > 6 && input < 1) {
			fmt.Println("Warning:Wrong Input")
		}
		switch input {
		case 1:
			fmt.Println("\nSearching an item in the phonebook (BST) ...\n====================")
			fmt.Print("Search for a contact: ")
			temp, _ := reader.ReadString('\n')
			temp = strings.Split(temp, "\n")[0]
			fmt.Println("Phonebook : Searching for: (", strings.ToUpper(temp), ")")
			fmt.Println("====================================\n")
			timeStart = time.Now()
			SearchWord(&BinarySearchTree, temp)
			timeStop = time.Now()
			var BSTtime = timeStop.Sub(timeStart)

			fmt.Println("\nSearching an item in the phonebook (AVL) ...\n")
			fmt.Println("Phonebook : Searching for: (", strings.ToUpper(temp), ")")
			fmt.Println("====================================\n")
			timeStart = time.Now()
			SearchWordAVL(&AVL, temp)
			timeStop = time.Now()
			var AVLtime = timeStop.Sub(timeStart)
			fmt.Println("\nThe search in BST took", BSTtime, "\n")
			fmt.Println("The search in AVL took", AVLtime)
		case 2:
			fmt.Print("Enter the information of the contact to be added:\nName:")
			tempName, _ := reader.ReadString('\n')
			listName := strings.Split(tempName, " ")
			if (len(listName) < 2) {
				fmt.Println("You have entered in a wrong format.Please check your input and try again.")
			} else {
				fmt.Print("\nTel:")
				tempPhone, _ := reader.ReadString('\n')
				fmt.Print("\nCity:")
				tempCity, _ := reader.ReadString('\n')
				if (len(listName) > 2) {
					BinarySearchTree.insert(listName[0], listName[1]+" "+listName[2][:len(listName[1])+1], tempPhone[:len(tempPhone)-1], tempCity)
					AVL.insert(listName[0], listName[1]+" "+listName[2][:len(listName[1])+1], tempPhone[:len(tempPhone)-1], tempCity)
				} else if (len(listName) < 4) {
					BinarySearchTree.insert(listName[0], listName[1][:len(listName[1])-1], tempPhone[:len(tempPhone)-1], tempCity)
					AVL.insert(listName[0], listName[1][:len(listName[1])-1], tempPhone[:len(tempPhone)-1], tempCity)
				}
			}
		case 3:
			var nameDel, secondDel string
			fmt.Println("\nDeleting an item from the phonebook ...\n====================")
			fmt.Print("Enter the fullname of the contact to be deleted:")
			fmt.Scan(&nameDel, &secondDel)
			timeStart = time.Now()
			var resultBST = BinarySearchTree.delete(nameDel, secondDel)
			timeStart = time.Now()
			var timeBST = timeStop.Sub(timeStart)
			timeStart = time.Now()
			AVL.delete(nameDel, secondDel)
			timeStart = time.Now()
			var timeAVL = timeStop.Sub(timeStart)
			if (resultBST) {
				fmt.Println("Deleted succesfully...")
				fmt.Println("Deletion from Binary Tree took ", timeBST)
				fmt.Println("Deletion from AVL Tree took ", timeAVL)
			} else {
				os.Exit(3)
			}
		case 4:
			fmt.Println("\nSaving the phonebook to a file (In-Order)...")
			fmt.Println("====================")
			timeStart = time.Now()
			AVL.inorder()
			timeStop = time.Now()
			fmt.Println("Saving the phonebook took ", timeStop.Sub(timeStart), " ...")

			fmt.Println("\nSaving the phonebook to a file (Pre-Order)...")
			fmt.Println("====================")
			timeStart = time.Now()
			AVL.preorder()
			timeStop = time.Now()
			fmt.Println("Saving the phonebook took ", timeStop.Sub(timeStart), " ...")
		case 5:
			timeStart = time.Now()
			AVL.PrintDiagram()
			timeStop = time.Now()
			var DiagramTimeAVL = timeStop.Sub(timeStart)
			timeStart = time.Now()
			BinarySearchTree.PrintDiagram()
			timeStop = time.Now()
			var DiagramTimeBST = timeStop.Sub(timeStart)

			fmt.Println("Saving the phonebook took ",DiagramTimeBST," for BST.")
			fmt.Println("Saving the phonebook took ",DiagramTimeAVL," for AVL.")

		case 6:
			flag = true
			fmt.Println("Exiting...")
		}
	}
}