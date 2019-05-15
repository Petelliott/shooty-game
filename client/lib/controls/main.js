export function setup (game) {
  let canvas = document.getElementById('screen')
  game.controls = {
    canvas: {},
    mouse: {}
  }

  canvas.addEventListener('mousemove', (ev) => {
    game.controls.mouse.x = event.clientX * canvas.width / canvas.scrollWidth
    game.controls.mouse.y = event.clientY * canvas.height / canvas.scrollHeight
  })
  canvas.addEventListener('mousedown',
    (ev) => game.controls.mouse.clicked = true)
  canvas.addEventListener('mouseup',
    (ev) => game.controls.mouse.clicked = false)

  window.addEventListener('keydown', (ev) => {
    console.log(game)
  })

  window.addEventListener('resize', setCanvasDimensions(game))
  setCanvasDimensions(game)()
}

function setCanvasDimensions (game) {
  return () => {
    let dim = getDimensions()

    let canvas = document.getElementById('screen')
    canvas.width = dim.x

    game.controls.canvas = dim
  }
}

function getDimensions() {
  let canvas = document.getElementById('screen')
  return { x: canvas.height / canvas.scrollHeight * canvas.scrollWidth, y: 1000 }
}
