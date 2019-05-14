export function draw (delta, game) {
  let canvas = document.getElementById('screen')
  if (canvas.getContext) {
    let ctx = canvas.getContext('2d')
    ctx.fillStyle = '#494949'
    ctx.fillRect(0, 0, game.controls.canvas.x, game.controls.canvas.y)
    renderWorld(delta, game, ctx)
  }
}

function renderWorld (delta, game, ctx) {
  ctx.fillStyle = game.controls.mouse.clicked ? '#9900CC' : '#0099CC'
  ctx.beginPath()
  ctx.ellipse(game.controls.mouse.x, game.controls.mouse.y, 10, 10,
    0, 0, 2 * Math.PI)
  ctx.fill()
}
