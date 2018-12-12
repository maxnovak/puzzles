#Input comes from https://adventofcode.com/2018/day/1/input
#copied to text file day1_input.txt for ease of solving problem.

sum = 0
File.open("day1_input.txt", "r") do |f|
	f.each_line do |line|
		sum += line.to_i
	end
end

puts sum