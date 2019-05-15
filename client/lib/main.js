import { draw } from './render/main.js'

import { setup as setupRender } from './render/main.js'
import { setup as setupServer } from './server/main.js'
import { setup as setupControls } from './controls/main.js'

export function init () {
  let game = setup()
  let time1 = window.performance.now()
  let time2 = 0
  const looper = () => {
    time2 = window.performance.now()
    loop(time2 - time1, game)
    time1 = time2

    window.requestAnimationFrame(looper)
  }
  window.requestAnimationFrame(looper)
}

function setup () {
  let game = {}
  game.world = {}

  setupServer(game)
  setupControls(game)
  setupRender(game)

  return game
}

function loop (delta, game) {
  draw(delta, game)
}
