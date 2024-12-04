from collections import Counter

def parse_input(file_path):
    left_list = []
    right_list = []
    
    with open(file_path, 'r') as file:
        for line in file:
            left, right = map(int, line.split())
            left_list.append(left)
            right_list.append(right)
    
    return left_list, right_list

def calculate_total_distance(left_list, right_list):
    left_list.sort()
    right_list.sort()
    
    total_distance = sum(abs(left - right) for left, right in zip(left_list, right_list))
    return total_distance

def calculate_similarity_score(left_list, right_list):

    right_counts = Counter(right_list)
    
    similarity_score = 0
    for num in left_list:
        similarity_score += num * right_counts.get(num, 0)
    
    return similarity_score

# left_list = [3, 4, 2, 1, 3, 3]
# right_list = [4, 3, 5, 3, 9, 3]

left_list, right_list = parse_input("input1.txt")
total_distance = calculate_total_distance(left_list, right_list)
similarity_score = calculate_similarity_score(left_list, right_list)

print(f"The total distance is: {total_distance}")
print(f"The similarity score is: {similarity_score}")