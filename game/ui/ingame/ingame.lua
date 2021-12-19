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
	self.box_node["fever_background"]=box_node.create(gui.get_node("fever_background"))
	self.box_node["fever_gague"]=box_node.create(gui.get_node("fever_gague"))
	self.box_node["game_over"]=box_node.create(gui.get_node("game_over"))
	self.text_node["retry"]=text_node.create(gui.get_node("retry"))
	self.text_node["rank"]=text_node.create(gui.get_node("rank"))
	self.text_node["nickname"]=text_node.create(gui.get_node("nickname"))
	self.text_node["comment"]=text_node.create(gui.get_node("comment"))
	self.box_node["filter"]=box_node.create(gui.get_node("filter"))
	self.box_node["alert"]=box_node.create(gui.get_node("alert"))
	self.box_node["alert_in/text_parent"]=box_node.create(gui.get_node("alert_in/text_parent"))
	self.text_node["alert_in/prefab"]=text_node.create(gui.get_node("alert_in/prefab"))
	self.box_node["alert_in/prefab_icon"]=box_node.create(gui.get_node("alert_in/prefab_icon"))
end
return G