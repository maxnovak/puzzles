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

p list_of_box_ids