local G={}
local box_node = require "DefUtil.DefUI.module.box_node"
local text_node = require "DefUtil.DefUI.module.text_node"
function G.init(self)
	self.box_node={}
	self.text_node={}
	self.box_node["start_btn"]=box_node.create(gui.get_node("start_btn"))
	self.text_node["text"]=text_node.create(gui.get_node("text"))
end
return G