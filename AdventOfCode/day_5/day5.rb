
#distance between capital and lowercase letters
comparisonAmount = "a".ord - "A".ord

p comparisonAmount

sequence = ""
File.open("day5_input.txt", "r") do |file|
	file.each_line do |data|
		sequence = data
	end
end

p sequence