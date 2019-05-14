const URL = 'ws://localhost:8049'

export function setup (game) {
  game.ws = new WebSocket(URL)
  game.ws.onmessage = getmessage(game)
  game.finishTurn = finishTurn(game)
}

function finishTurn (game) {
  return function () {
    game.ws.send(game.orders)
  }
}

function getmessage (game) {
  return function (event) {
    console.log(event, game)
  }
}
