# Phonebook
 A phonebook with tree structured printing functionality.
 
 ```md
 Please enter the contact file name: 
Loading the phonebook into a BST 

Phonebook creation in BST took 4 milliseconds

The tree is not balanced

The heights of BST are for left: 5 and right: 7
loading the phonebook into an AVL 

Phonebook creation in AVL took 3 milliseconds

The tree is balanced

The heights of AVL are for left: 4 and right: 3

Choose which action to perform from 3 to 7: 
1 - Search a phonebook contact 
2 - Adding a phonebook contact 
3 - Deleting a phonebook contact 
4 - Print the phonebook to a file(inorder) 
    Print the phonebook to a file(postorder) 
5 - Draw the Phonebook as a Tree to a file
6 - Press 6 to exit


Searching an item in the phonebook (BST) ...
====================
Search for a contact: 
Phonebook : Searching for: (JOHN WATSON)
====================================

JOHN WATSON +905550292913 Izmir

Searching an item in the phonebook (AVL) ...

Phone book : Section (JOHN WATSON)
====================================

JOHN WATSON +905550292913 Izmir

The search in BST took 0 milliseconds...


The search in AVL took 0 milliseconds...


Choose which action to perform from 3 to 7: 
1 - Search a phonebook contact 
2 - Adding a phonebook contact 
3 - Deleting a phonebook contact 
4 - Print the phonebook to a file(inorder) 
    Print the phonebook to a file(postorder) 
5 - Draw the Phonebook as a Tree to a file
6 - Press 6 to exit


Adding an item to the phonebook (BST) ...
====================
Enter the information of the contact to be added: 
Name: 

Tel: 
City: 

Contact has been added successfully to the BST

Adding an item to the phonebook (AVL) ...
====================

Contact has been added successfully to the AVL tree


Adding a contact to the Binary Tree took 13996 nanoseconds...
Adding a contact to the AVL tree took 33591 nanoseconds...

Choose which action to perform from 3 to 7: 
1 - Search a phonebook contact 
2 - Adding a phonebook contact 
3 - Deleting a phonebook contact 
4 - Print the phonebook to a file(inorder) 
    Print the phonebook to a file(postorder) 
5 - Draw the Phonebook as a Tree to a file
6 - Press 6 to exit


Deleting an item from the phonebook ...
====================
Enter the fullname of the contact to be deleted:  
Deleted succesfully...
Deletion from Binary Tree took 236532 nanoseconds...
Deletion from AVL Tree took 32658 nanoseconds...

Choose which action to perform from 3 to 7: 
1 - Search a phonebook contact 
2 - Adding a phonebook contact 
3 - Deleting a phonebook contact 
4 - Print the phonebook to a file(inorder) 
    Print the phonebook to a file(postorder) 
5 - Draw the Phonebook as a Tree to a file
6 - Press 6 to exit


Saving the phonebook to a file (In-Order)...
====================
Saving the phonebook took 350366nanoseconds...

Saving the phonebook to a file (Pre-Order)...
====================
Saving the phonebook took 334970nanoseconds...

Choose which action to perform from 3 to 7: 
1 - Search a phonebook contact 
2 - Adding a phonebook contact 
3 - Deleting a phonebook contact 
4 - Print the phonebook to a file(inorder) 
    Print the phonebook to a file(postorder) 
5 - Draw the Phonebook as a Tree to a file
6 - Press 6 to exit


Choose which action to perform from 3 to 7: 
1 - Search a phonebook contact 
2 - Adding a phonebook contact 
3 - Deleting a phonebook contact 
4 - Print the phonebook to a file(inorder) 
    Print the phonebook to a file(postorder) 
5 - Draw the Phonebook as a Tree to a file
6 - Press 6 to exit


Exiting...
 ```
 
 ## AVL Tree
 ```
 |__ JOHN WATSON
    |-- CHARLES SMITH
    |   |-- ASHLEY TIRADO
    |   |   |-- ASHLEY JEPSEN
    |   |   |__ CHARITY RUGGIERO
    |   |       |-- CAROLYN JOHNSON
    |   |__ FRANK THOMPSON
    |       |-- ELBERT WOMACK
    |       |__ HAROLD STENSETH
    |__ RHONDA CASHION
        |-- PAMELA ALEXANDER
        |   |-- NORMAN HO
        |   |__ PATRICIA GOODE
        |__ SHARON PERSAUD
            |-- ROGER JOCHIM
            |__ SHIRLEY MICKELSON
  ```
  
## BST
```
|__ HAROLD STENSETH
    |-- FRANK THOMPSON
    |   |-- ASHLEY TIRADO
    |   |   |-- ASHLEY JEPSEN
    |   |   |__ CHARLES SMITH
    |   |       |-- CHARITY RUGGIERO
    |   |       |   |-- CAROLYN JOHNSON
    |   |       |__ ELBERT WOMACK
    |__ JOHN WATSON
        |__ SHARON PERSAUD
            |-- ROGER JOCHIM
            |   |-- NORMAN HO
            |   |   |__ PAMELA ALEXANDER
            |   |       |__ RHONDA CASHION
            |   |           |-- PATRICIA GOODE
            |__ SHIRLEY MICKELSON
```
## Traverse Nodes
```go
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
```
