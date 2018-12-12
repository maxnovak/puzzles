#Input comes from https://adventofcode.com/2018/day/1/input
#copied to text file day1_input.txt for ease of solving problem.

current_frequency = 0
File.open("day1_input.txt", "r") do |file|
	file.each_line do |frequency_value|
		current_frequency += frequency_value.to_i
	end
end

#Part 1 solution
puts current_frequency