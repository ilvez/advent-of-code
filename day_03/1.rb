require 'amazing_print'

raw_map = IO.readlines('input', chomp: true)

x_pos = 0
tree_count = raw_map.select do |step_scene|
  is_tree = step_scene[x_pos % step_scene.size] == '#'
  x_pos += 3
  is_tree
end.size

ap tree_count
