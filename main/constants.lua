-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.
local M = {}
M.SCREEN_WIDTH = 720
M.SCREEN_HEIGHT = 1280
M.MAX_SMILEYS = 20
M.MAX_LIFES = 3
M.SHOW_ADS = true

-- tables with smileys
M.SMILEY_IDS = { "60a", "60b", "60c", "60d", "60e", "60f", "61a", "61b","61c", "61d", "61e","61f", "62a", "62b", "62c", "62d", "62e", "62f", "64a", "92a", "92b", "92d", "435", "600", "601", "602", "603", "604", "605", "606", "607", "609", "610", "611", "612", "613", "614", "615", "616", "617", "618", "619", "620", "622", "625", "626", "628", "629", "630", "632", "635", "636", "641", "642", "644", "648", "649"}
M.SMILEY_SUCCESS_IDS= {"u1f60f", "u1f4aa", "u1f44c", "u1f44d", "u1f44f", "u1f60e", "u1f917", "u1f929", "u1f970", "u1f973"}
M.SMILEY_GAMEOVER_IDS = {"u1f4a9", "u1f61e", "u1f61f", "u1f62b", "u1f62d", "u1f62e", "u1f92f", "u1f610", "u1f613", "u1f629", "u1f635", "u1f648", "u2639"}
M.ITEM_IDS = {"u2764", "u23f2", "u1f0cf", "u2b50"} --  star:"u2b50"
-- item ids
M.EMOJI_HEART_ID = 1
M.EMOJI_TIME_ID = 2
M.EMOJI_JOKER_ID = 3
M.EMOJI_STAR_ID = 4
-- amount of stars needed to restart level
M.CONTINUE_STARS = 500
-- amount of resets for showing interstitial ad
M.MAX_RESETS = 3
-- number of different backgrounds
M.BACKGROUND_COUNT =4
-- initial level time
M.LEVEL_TIME = 60
-- initial level start after restart at gameover
M.LEVEL_TIME_RESTART = 30
-- additional time gaining after comleting a level
M.ADDITIONAL_TIME = 5
-- max time for the current smiley to not being spawned
M.SPAWN_TIME_MAX_CURRENT_SMILEY = 3.5

-- probability in percent for spawning items
M.ITEM_PROB = 5
-- seconds gaining when collecting time-item
M.TIME_ITEM_VALUE = 5
-- time which has to pass for the rewarded ad to be active again
M.REWARDED_AD_INTERVAL = 10 --seconds

-- initial amount of different smileys at one level
M.LEVEL_SMILEYS_BASE = 6
return M