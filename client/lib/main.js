export function init() {
  draw()
}

function draw() {
  let canvas = document.getElementById('screen')
  if (canvas.getContext) {
    let ctx = canvas.getContext('2d')
    ctx.fillStyle = '#494949'
    ctx.fillRect(0, 0, 1000, 1000)
  }
}
