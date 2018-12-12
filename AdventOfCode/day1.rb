#Input comes from https://adventofcode.com/2018/day/1/input
#copied to text file day1_input.txt for ease of solving problem.

current_frequency = 0
read_in_frequencies = []
reached_frequencies = [0]

File.open("day1_input.txt", "r") do |file|
	file.each_line do |frequency_value|
		read_in_frequencies.push(frequency_value.to_i)
	end
end

for frequency in read_in_frequencies
	current_frequency += frequency
end

terminal_frequency = current_frequency
current_frequency = 0
times = 0
while true
	for frequency in read_in_frequencies
		current_frequency += frequency
		if reached_frequencies.include?(current_frequency)
			repeated_frequency = current_frequency
			break
		else
			reached_frequencies.push(current_frequency)
		end
	end
	if repeated_frequency != nil
		puts "looped #{times}"
		break
	end
	times+=1
end

#Part 1 solution
puts terminal_frequency

#Part 2 solution
#with this data set it has to loop 135 times - so will take some time to complete
puts repeated_frequency