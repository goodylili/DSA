# Learning Data Structures and Algorithms with Design Patterns

## Data Structures
A data structure is the organization of data in a computer's memory in a way that reduces storage space and difficulty of performing a task.
Data structures are useful for handling large amounts of data in different cases like database management and internet indexing services.
Data structures can be divided into different types:
1. linear
2. non linear
3. homogenous
4. heterogeneous
5. dynamic

Owing to properties of the classifications above, there are differences in performance of different data structures. Thus, you will 
need to learn how to choose the right data structures for a use case, as well as the structural design pattern.
When there is a problem to be solved with code, the first thing to do is to decide these three features; what data structure(s) are 
best used, what algorithms solves this easily and efficiently, what design pattern will make this easily structured? The next step is to do
a performance analysis of the solution code, in terms of the time and space consumed. Following these procedures will land you an optimal 
solution to the problem.

## Classification of Data Structures
### Heterogeneous Data Structures
If the problem requires more than one data type, you will need to use a heterogeneous data structure. They include:
- linked lists
- ordered lists
- unordered lists

### Linear Data Structures
Linear data structures are the popular ones. They include:
- lists
- sets
- tuples
- queues
- stacks
- heaps

### Non-Linear Data Structures
Non-linear data structures include the following:
- trees
- tables
- containers

### Homogenous Data Structures
These include:
- 2-dimensional and multidimensional arrays

### Dynamic Data Structures
These include:
- dictionaries
- tree sets
- sequences.

## Lists(slices) and Arrays
The array or list(slice) is a sequence of elements. It is the most basic data structure in computing. Elements are added through `PushBack` 
method on a list in Go, and the values of the elements printed out with `Value`.

## Tuples
A tuple is a finitely sorted list of elements. It is a data structure that groups data. They have related fields of different data types and
are typically immutable sequential collections.  To modify a tuple, you just have to change fields using + or *. A database record is 
referred to as a tuple.

## Heaps
A heap data structure is based on the `heap` property. It is used in selection, graph, and k-way merge algorithms. Operations like finding,
merging, insertion, key changes, and deleting are performed on heaps. The value stored at each node is greater than or equal to its children.
If it is in descending order, it is a maximum heap; otherwise, it's a minimum heap.