#Input comes from https://adventofcode.com/2018/day/1/input
#copied to text file day1_input.txt for ease of solving problem.

current_frequency = 0
read_in_frequencies = []

File.open("day1_input.txt", "r") do |file|
	file.each_line do |frequency_value|
		read_in_frequencies.push(frequency_value.to_i)
	end
end

for frequency in read_in_frequencies
	current_frequency += frequency
end
#Part 1 solution
puts current_frequency