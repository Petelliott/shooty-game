export function drawImage(ctx, game, id, x, y, w, h) {
  const imgData = game.images.data[id]
  if (!imgData) {
    loadImage(game, id)
  } else if (imgData.loaded) {
    ctx.drawImage(game.images.res[imgData.src],
      imgData.sx, imgData.sy, imgData.sw, imgData.sh,
      x, y, w, h)
  }
}

function loadImage(game, id) {
  game.images.data[id] = {
    loaded: false
  }
  let imageData = config.graphics.images[id]
  if (!imageData) { console.warn(`Unknown image ID: ${id}`) }

  let finalImageData = imageData
  if (imageData.srcType === "tileset") {
    finalImageData = config.graphics.tilesets[imageData.src]
  }

  let isTS = finalImageData !== imageData
  let image = new Image()
  image.onload = () => {
    const resId = (Math.random() * 1000000000) | 0
    game.images.res[resId] = image
    game.images.data[id] = {
      loaded: true,
      src: resId,
      sx: isTS ? imageData.x * (finalImageData.tileWidth + finalImageData.tileSpacing) : 0,
      sy: isTS ? imageData.y * (finalImageData.tileHeight + finalImageData.tileSpacing) : 0,
      sh: isTS ? finalImageData.tileHeight : image.naturalHeight,
      sw: isTS ? finalImageData.tileWidth : image.naturalWidth,
    }
  }
  image.src = finalImageData.src

}
