#Input comes from https://adventofcode.com/2018/day/3/input
#copied to text file day3_input.txt for ease of solving problem.
require_relative "SuitClaim"

suitClaims = []

File.open("day3_input.txt", "r") do |file|
	file.each_line do |claim_data|
		data = claim_data.split(" ")
		startCoordinates = data[2].split(",")
		dimentions = data[3].split("x")
		suitClaims.push(
			SuitClaim.new(
				data[0],
				startCoordinates[0],
				startCoordinates[1].tr(":", ""),
				dimentions[0],
				dimentions[1]))
	end
end

fabricMap = Array.new(1000){Array.new(1000, 0)}

suitClaims.each_with_index do |claim, i|
	x1 = claim.startX
	y1 = claim.startY
	x2 = claim.startX + claim.width
	y2 = claim.startY + claim.height
	puts x1, x2
	puts y1, y2
	for i in x1..x2-1
		for j in y1..y2-1
			fabricMap[j][i] += 1
		end
	end
end

count=0

fabricMap.each do |row|
	row.each do |item|
		if item >= 2
			count += 1
		end
	end
end

fabricMap.each do |row|
	p row
end

#Part 1 solution - 101565
puts count
