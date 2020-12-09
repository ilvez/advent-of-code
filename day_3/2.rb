require 'amazing_print'

@raw_map = IO.readlines('input', chomp: true)

def tree_amount_on_slope(right_step, down_step)
  down_steps = @raw_map
    .each_with_index
    .map { |scene, i| (i % down_step).zero? ? scene : nil }
    .compact[1..]
  x_pos = right_step
  down_steps.select do |step_scene|
    is_tree = step_scene[x_pos % step_scene.size] == '#'
    x_pos += right_step
    is_tree
  end.size
end

ap(
  tree_amount_on_slope(1, 1) *
    tree_amount_on_slope(3, 1) *
    tree_amount_on_slope(5, 1) *
    tree_amount_on_slope(7, 1) *
    tree_amount_on_slope(1, 2)
)
