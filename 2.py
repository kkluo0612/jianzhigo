class Heap:
    def __init__(self,capacity):
        self.data=[0]*(capacity+1)
        self.capacity = capacity
        self.count=0

    @classmethod
    def parent(cls,child_index):
        return child_index//2
    
    @classmethod
    def left(cls,parent_index):
        return 2*parent_index

    @classmethod
    def right(cls,parent_index):
        return 2*parent_index+1
    
    def siftup(self):
        i,parent = self.count,Heap.parent(self.count)
        while parent and self.data[i]>self.data[parent]:
            self.data[i],self.data[parent] = self.data[parent],self.data[i]
            i,parent = parent,Heap.parent(parent)

    @classmethod
    def siftdown(cls,a,count,root_index=1):
        i = large_child_index=root_index
        while True:
            left,right=cls.left(i),cls.right(i)
            if left<=count and a[i]<a[left]:
                large_child_index=left
            if right<=count and a[large_child_index]<a[right]:
                large_child_index=right
            if large_child_index==i: break
            a[i],a[large_child_index]=a[large_child_index],a[i]
            i=large_child_index

    def insert(self,value):
        if self.count>=self.capacity: return 
        self.count+=1
        self.data[self.count]=value
        self.siftup()

    def remove_max(self):
        if self.count:
            result=self.data[1]
            self.data[1]=self.data[self.count]
            self.count-=1
            Heap.siftdown(self.data,self.count)
            return result

    @classmethod
    def build_heap(cls,a):
        for i in range((len(a)-1)//2,0,-1):
            cls.siftdown(a,len(a)-1,i)

    @classmethod
    def sort(cls,a):
        cls.build_heap(a)
        k=len(a)-1
        while k>1:
            a[1],a[k]=a[k],a[1]
            k-=1
            cls.siftdown(a,k)

    def __repr__(self):
        return self.data[1:self.count+1].__repr__()

if __name__ == '__main__':
    heap=Heap(10)
    # heap.insert(3)
    # heap.insert(9)
    # heap.insert(1)
    # heap.insert(8)
    # heap.insert(7)
    # heap.insert(3)
    # print(heap)
    # for _ in range(6):
    #     print(heap.remove_max())
    a=[0,6,3,4,0,9,2,7,5,-2,8,1,6,10]
    heap.sort(a)
    print(a[1:])