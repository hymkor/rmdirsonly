usage:
	@echo Usage:
	@echo make test1
	@echo make test2
	@echo make clean

test1:
	mkdir testDir\a\b\c\d
	rmdirsonly testDir

test2:
	mkdir testDir\a\b\c\d
	echo ahaha > testDir\a\b\c\file
	rmdirsonly testDir

clean:
	if exist testDir rmdir /s testDir

