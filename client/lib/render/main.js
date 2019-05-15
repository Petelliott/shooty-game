import { drawImage } from './images.js'

export function setup (game) {
  game.images = {
    res: {},
    data: {}
  }
}

export function draw (delta, game) {
  let canvas = document.getElementById('screen')
  if (canvas.getContext) {
    let ctx = canvas.getContext('2d')
    ctx.fillStyle = '#494949'
    ctx.fillRect(0, 0, game.controls.canvas.x, game.controls.canvas.y)
    renderWorld(delta, game, ctx)
    drawImage(ctx, game, 'ground', 0, 0, 100, 100)
    drawImage(ctx, game, 'tilePreview', 100, 100, 900, 900)
  }
}

function renderWorld (delta, game, ctx) {
  ctx.fillStyle = game.controls.mouse.clicked ? '#9900CC' : '#0099CC'
  ctx.beginPath()
  drawImage(ctx, game, 'ground', 0, 0, 100, 100)
  drawImage(ctx, game, 'tilePreview', 100, 100, 900, 900)
  ctx.ellipse(game.controls.mouse.x, game.controls.mouse.y, 10, 10,
    0, 0, 2 * Math.PI)
  ctx.fill()
}
