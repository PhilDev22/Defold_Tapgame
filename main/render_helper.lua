-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.
local M = {}

function M.get_screen_height()
	return M.screen_height
end

function M.get_screen_width()
	return M.screen_width
end

return M