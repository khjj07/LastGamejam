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
  data: "tile_set: \"/game/enemy/enemy12.atlas\"\n"
  "default_animation: \"walk\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 17.0
    y: 23.0
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
  data: "tile_set: \"/game/enemy/enemy12.atlas\"\n"
  "default_animation: \"walk\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: -17.0
    y: 23.0
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
    y: 42.0
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
  "default_animation: \"UI_FeverGauge_Background\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: -17.0
    y: 42.0
    z: 0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}