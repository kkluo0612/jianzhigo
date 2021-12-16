pick = []
max_value=[]

def bag(capacity,cur_weight,item,pick_idx):
    if pick_idx >= len(item) or cur_weight==capacity:
        global max_value
        if get_value(item,pick)>get_value(item,max_value):
            max_value=pick.copy()
        else:
            item_weight = item[pick_idx][0]
            if cur_weight + item_weight <=capacity:
                pick[pick_idx]=1
                bag(capacity,cur_weight+item_weight,item.pick_idx+1)
            pick[pick_idx]=0
            bag(capacity,cur_weight,item,pick_idx+1)

def get_value(item,pick_item):
    values=[_[1]for _ in item]
    return sum([a*b for a,b in zip(values,pick_item)])

if __name__=='__main__':
    item=[(3, 5), (2, 2), (1, 4), (1, 2), (4, 10)]
    capacity=8

    print('---item info---')
    print(item)

    print('\n---capcity---')
    print(capacity)

    pick=[0]*len(item)
    bag(capacity,0,item,0)

    print('\n---picks---')
    print(max_value)

    print('\n---value---')
    print(get_value(item,max_value))