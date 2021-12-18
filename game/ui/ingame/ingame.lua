local G={}
local box_node = require "DefUtil.DefUI.module.box_node"
local text_node = require "DefUtil.DefUI.module.text_node"
function G.init(self)
	self.box_node={}
	self.text_node={}
	self.box_node["life3"]=box_node.create(gui.get_node("life3"))
	self.box_node["life2"]=box_node.create(gui.get_node("life2"))
	self.box_node["life1"]=box_node.create(gui.get_node("life1"))
	self.text_node["party"]=text_node.create(gui.get_node("party"))
	self.text_node["fever_text"]=text_node.create(gui.get_node("fever_text"))
end
return G