local G={}
local box_node = require "DefUtil.DefUI.module.box_node"
local text_node = require "DefUtil.DefUI.module.text_node"
function G.init(self)
	self.box_node={}
	self.text_node={}
	self.box_node["box"]=box_node.create(gui.get_node("box"))
	self.text_node["start"]=text_node.create(gui.get_node("start"))
	self.text_node["howtoplay"]=text_node.create(gui.get_node("howtoplay"))
	self.text_node["credit"]=text_node.create(gui.get_node("credit"))
	self.text_node["exit"]=text_node.create(gui.get_node("exit"))
	self.text_node["title"]=text_node.create(gui.get_node("title"))
	self.box_node["info"]=box_node.create(gui.get_node("info"))
	self.box_node["developers"]=box_node.create(gui.get_node("developers"))
	self.text_node["designer"]=text_node.create(gui.get_node("designer"))
	self.text_node["programmer"]=text_node.create(gui.get_node("programmer"))
	self.text_node["developer_title"]=text_node.create(gui.get_node("developer_title"))
end
return G