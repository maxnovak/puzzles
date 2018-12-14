#Input comes from https://adventofcode.com/2018/day/3/input
#copied to text file day3_input.txt for ease of solving problem.
require_relative "SuitClaim"

suitClaims = []

File.open("day3_input.txt", "r") do |file|
	file.each_line do |claim_data|
		data = claim_data.split(" ")
		startCoordinates = data[2].split(",")
		dimentions = data[3].split("x")
		suitClaims.push(SuitClaim.new(data[0], startCoordinates[0], startCoordinates[1].tr(":", ""), dimentions[0], dimentions[1]))
	end
end

p suitClaims