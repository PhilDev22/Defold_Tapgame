components {
  id: "gui"
  component: "/main/main.gui"
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
components {
  id: "script"
  component: "/main/main.script"
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
components {
  id: "sounds"
  component: "/main/sounds.script"
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
embedded_components {
  id: "sprite_background"
  type: "sprite"
  data: "tile_set: \"/main/backgrounds.atlas\"\n"
  "default_animation: \"background_ingame_1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 320.0
    y: 568.0
    z: -1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sound_pop1"
  type: "sound"
  data: "sound: \"/assets/sounds/pop1.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
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
embedded_components {
  id: "sound_pop2"
  type: "sound"
  data: "sound: \"/assets/sounds/pop2.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
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
embedded_components {
  id: "sound_pop3"
  type: "sound"
  data: "sound: \"/assets/sounds/pop3.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
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
embedded_components {
  id: "sound_pop4"
  type: "sound"
  data: "sound: \"/assets/sounds/pop4.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
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
embedded_components {
  id: "sound_complete"
  type: "sound"
  data: "sound: \"/assets/sounds/complete2.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 0.1\n"
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
embedded_components {
  id: "sound_item1"
  type: "sound"
  data: "sound: \"/assets/sounds/item1.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 0.6\n"
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
embedded_components {
  id: "sound_stars"
  type: "sound"
  data: "sound: \"/assets/sounds/counter2.wav\"\n"
  "looping: 1\n"
  "group: \"master\"\n"
  "gain: 0.4\n"
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
embedded_components {
  id: "sound_item2"
  type: "sound"
  data: "sound: \"/assets/sounds/Item2.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
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
embedded_components {
  id: "sound_click1"
  type: "sound"
  data: "sound: \"/assets/sounds/click.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 0.8\n"
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
embedded_components {
  id: "sound_fail"
  type: "sound"
  data: "sound: \"/assets/sounds/fail.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 0.8\n"
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
embedded_components {
  id: "music"
  type: "sound"
  data: "sound: \"/assets/sounds/music1.ogg\"\n"
  "looping: 1\n"
  "group: \"master\"\n"
  "gain: 0.4\n"
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
