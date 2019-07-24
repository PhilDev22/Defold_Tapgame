-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.
local M = {}
function get_table_ids(table)
	local index={}
	for k,v in pairs(table) do
		index[ hash(v) ]=k
	end
	return index
end
return M