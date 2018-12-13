#Input comes from https://adventofcode.com/2018/day/2/input
#copied to text file day2_input.txt for ease of solving problem.

two_counts = 0
three_counts = 0
list_of_box_ids = []

File.open("day2_input.txt", "r") do |file|
	file.each_line do |box_id|
		list_of_box_ids.push(box_id.strip)
	end
end

has_twos = false
has_threes = false

for box_id in list_of_box_ids
	characters = box_id.split('')
	for character in characters.uniq
		if box_id.count(character) == 2
			has_twos = true
		elsif box_id.count(character) == 3
			has_threes = true
		end
	end
	if has_twos
		two_counts += 1
		has_twos = false
	end
	if has_threes
		three_counts += 1
		has_threes = false
	end
end

puts "Two counts: #{two_counts}"
puts "Three counts: #{three_counts}"

#Part 1 solution
puts "Checksum: #{two_counts * three_counts}"
