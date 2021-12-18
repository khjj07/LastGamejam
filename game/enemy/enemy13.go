components {
  id: "enemy"
  component: "/game/enemy/enemy.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "speed"
    value: "75.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "hp"
    value: "20.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "maxhp"
    value: "40.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "tag1"
    value: "human"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "tag2"
    value: "sword"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "fliped_sprite"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "custom_scale"
    value: "1.5, 1.5, 1.0"
    type: PROPERTY_TYPE_VECTOR3
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/game/enemy/enemy13.atlas\"\n"
  "default_animation: \"walk\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 20.0
    y: 3.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "flip_sprite"
  type: "sprite"
  data: "tile_set: \"/game/enemy/enemy13.atlas\"\n"
  "default_animation: \"walk\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: -20.0
    y: 3.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "hp"
  type: "sprite"
  data: "tile_set: \"/game/enemy/hp.atlas\"\n"
  "default_animation: \"UI_Hp_Fill\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: -17.0
    y: 38.0
    z: 0.2
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "hp_back"
  type: "sprite"
  data: "tile_set: \"/game/enemy/hp.atlas\"\n"
  "default_animation: \"UI_Hp_Background\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: -17.0
    y: 38.0
    z: 0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"deadline\"\n"
  "mask: \"startline\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 10.0\n"
  "  data: 10.0\n"
  "  data: 10.0\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  "bullet: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
