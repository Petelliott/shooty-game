export function init() {
  console.log('Loaded')
}

function draw() {
  let canvas = document.getElementById('screen')
  if (canvas.getContext) {
    let ctx = canvas.getContext('2d')
  }
}
