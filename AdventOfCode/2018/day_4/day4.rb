#Input comes from https://adventofcode.com/2018/day/3/input
#copied to text file day3_input.txt for ease of solving problem.

require_relative "GuardEvent"
require 'pp'

events = []
guardSleepTime = {}
guardSleepTime.default = 0
guardSleepSchedule = {}

File.open("day4_input.txt", "r") do |file|
	file.each_line do |claim_data|
		data = claim_data.split(" ")
		events.push(GuardEvent.new(data[0].tr("[",""), data[1].tr("]",""), data[2..-1]))
	end
end

events = events.sort_by { |event| [event.date, event.time]  }

currentGuardId = ""
startTime = 0

#assuming that they only fallasleep and wakeup at 12-1 am
events.each do |event|
	if /#/.match(event.event[1])
		currentGuardId = event.event[1]
	elsif /falls/.match(event.event[0])
		startTime = event.time.split(":")[1].to_i
	elsif /wakes/.match(event.event[0])
		endTime = event.time.split(":")[1].to_i
		sleepTime = endTime - startTime
		guardsSchedule = guardSleepSchedule[currentGuardId] || Array.new(60,0)
		guardsSchedule[startTime..endTime-1].each.with_index(startTime) do |item, index|
			guardsSchedule[index] = item + 1
		end
		guardSleepSchedule[currentGuardId] = guardsSchedule
		guardSleepTime[currentGuardId] = guardSleepTime[currentGuardId] + sleepTime
	end
end

guard = guardSleepTime.max_by{|key, value| value}
p guard
sleepTime = guardSleepSchedule[guard[0]].rindex(guardSleepSchedule[guard[0]].max)
p sleepTime

#Part 1 solution
p guard[0].tr("#", "").to_i * sleepTime

guardSleepSchedule.each do |guard|
	sleepingTime = guardSleepSchedule[guard[0]].rindex(guardSleepSchedule[guard[0]].max)
	p "GuardId: #{guard[0]}, Time most asleep: #{sleepingTime} for #{guardSleepSchedule[guard[0]][sleepingTime]} instances"
end

#Part 2 solution
p 239*33 #don't wanna store the values right now - took the values from the output and multiplied them here